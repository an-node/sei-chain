package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tokenfactorytypes "github.com/sei-protocol/sei-chain/x/tokenfactory/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/std"

	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/app"
	dextypes "github.com/sei-protocol/sei-chain/x/dex/types"
	oracletypes "github.com/sei-protocol/sei-chain/x/oracle/types"
)

var TestConfig EncodingConfig

const (
	VortexData = "{\"position_effect\":\"Open\",\"leverage\":\"1\"}"
)

var FromMili = sdk.NewDec(1000000)

func init() {
	cdc := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	TestConfig = EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             cdc,
	}
	std.RegisterLegacyAminoCodec(TestConfig.Amino)
	std.RegisterInterfaces(TestConfig.InterfaceRegistry)
	app.ModuleBasics.RegisterLegacyAminoCodec(TestConfig.Amino)
	app.ModuleBasics.RegisterInterfaces(TestConfig.InterfaceRegistry)
}

func run(config Config) {
	// Start metrics collector in another thread
	metricsServer := MetricsServer{}
	go metricsServer.StartMetricsClient(config)
	sleepDuration := time.Duration(config.LoadInterval) * time.Second

	if config.Constant {
		fmt.Printf("Running in constant mode with interval=%d\n", config.LoadInterval)
		for {
			runOnce(config)
			fmt.Printf("Sleeping for %f seconds before next run...\n", sleepDuration.Seconds())
			time.Sleep(sleepDuration)
		}
	} else {
		runOnce(config)
		fmt.Print("Sleeping for 60 seconds for metrics to be scraped...\n")
		time.Sleep(time.Duration(60))
	}
}

func runOnce(config Config) {
	client := NewLoadTestClient(config)
	client.SetValidators()

	if config.TxsPerBlock < config.MsgsPerTx {
		panic("Must have more TxsPerBlock than MsgsPerTx")
	}

	configString, _ := json.Marshal(config)
	fmt.Printf("Running with \n %s \n", string(configString))

	fmt.Printf("%s - Starting block prepare\n", time.Now().Format("2006-01-02T15:04:05"))
	workgroups, sendersList := client.BuildTxs()

	go client.SendTxs(workgroups, sendersList)

	// Waits until SendTx is done processing before proceeding to write and validate TXs
	client.GatherTxHashes()

	// Records the resulting TxHash to file
	client.WriteTxHashToFile()
	fmt.Printf("%s - Finished\n", time.Now().Format("2006-01-02T15:04:05"))

	// Validate Tx will close the connection when it's done
	go client.ValidateTxs()
}

func (c *LoadTestClient) generateMessage(config Config, key cryptotypes.PrivKey, msgPerTx uint64) (sdk.Msg, bool) {
	var msg sdk.Msg
	messageTypes := strings.Split(config.MessageType, ",")
	rand.Seed(time.Now().UnixNano())
	messageType := messageTypes[rand.Intn(len(messageTypes))]
	fmt.Printf("Message type: %s\n", messageType)
	switch messageType {
	case Bank:
		msg = &banktypes.MsgSend{
			FromAddress: sdk.AccAddress(key.PubKey().Address()).String(),
			ToAddress:   sdk.AccAddress(key.PubKey().Address()).String(),
			Amount: sdk.NewCoins(sdk.Coin{
				Denom:  "usei",
				Amount: sdk.NewInt(1),
			}),
		}
	case Dex:
		price := config.PriceDistr.Sample()
		quantity := config.QuantityDistr.Sample()
		contract := config.ContractDistr.Sample()
		orderPlacements := generateDexOrderPlacements(config, key, msgPerTx, price, quantity)
		amount, err := sdk.ParseCoinsNormalized(fmt.Sprintf("%d%s", price.Mul(quantity).Ceil().RoundInt64(), "usei"))
		if err != nil {
			panic(err)
		}
		msg = &dextypes.MsgPlaceOrders{
			Creator:      sdk.AccAddress(key.PubKey().Address()).String(),
			Orders:       orderPlacements,
			ContractAddr: contract,
			Funds:        amount,
		}
	case Staking:

		delegatorAddr := sdk.AccAddress(key.PubKey().Address()).String()
		chosenValidator := c.Validators[rand.Intn(len(c.Validators))].OperatorAddress
		// Randomly pick someone to redelegate / unbond from
		srcAddr := ""
		for k := range c.DelegationMap[delegatorAddr] {
			if k == chosenValidator {
				continue
			}
			srcAddr = k
			break
		}
		msg = c.generateStakingMsg(delegatorAddr, chosenValidator, srcAddr)
	case Tokenfactory:
		denomCreatorAddr := sdk.AccAddress(key.PubKey().Address()).String()
		// No denoms, let's mint
		randNum := rand.Float64()
		denom, ok := c.TokenFactoryDenomOwner[denomCreatorAddr]
		switch {
		case !ok || randNum <= 0.33:
			subDenom := fmt.Sprintf("tokenfactory-created-denom-%d", time.Now().UnixMilli())
			denom = fmt.Sprintf("factory/%s/%s", denomCreatorAddr, subDenom)
			msg = &tokenfactorytypes.MsgCreateDenom{
				Sender:   denomCreatorAddr,
				Subdenom: subDenom,
			}
			c.TokenFactoryDenomOwner[denomCreatorAddr] = denom
		case randNum <= 0.66:
			msg = &tokenfactorytypes.MsgMint{
				Sender: denomCreatorAddr,
				Amount: sdk.Coin{Denom: denom, Amount: sdk.NewInt(10)},
			}
		default:
			msg = &tokenfactorytypes.MsgBurn{
				Sender: denomCreatorAddr,
				Amount: sdk.Coin{Denom: denom, Amount: sdk.NewInt(1)},
			}
		}
	case FailureBankMalformed:
		var denom string
		if rand.Float64() < 0.5 {
			denom = "unknown"
		} else {
			denom = "other"
		}
		msg = &banktypes.MsgSend{
			FromAddress: sdk.AccAddress(key.PubKey().Address()).String(),
			ToAddress:   sdk.AccAddress(key.PubKey().Address()).String(),
			Amount: sdk.NewCoins(sdk.Coin{
				Denom:  denom,
				Amount: sdk.NewInt(1),
			}),
		}
	case FailureBankInvalid:
		var amountUsei int64
		amountUsei = 1000000000000000000
		msg = &banktypes.MsgSend{
			FromAddress: sdk.AccAddress(key.PubKey().Address()).String(),
			ToAddress:   sdk.AccAddress(key.PubKey().Address()).String(),
			Amount: sdk.NewCoins(sdk.Coin{
				Denom:  "usei",
				Amount: sdk.NewInt(amountUsei),
			}),
		}
	case FailureDexMalformed:
		price := config.PriceDistr.InvalidSample()
		quantity := config.QuantityDistr.InvalidSample()
		contract := config.ContractDistr.Sample()
		orderPlacements := generateDexOrderPlacements(config, key, msgPerTx, price, quantity)
		amount, err := sdk.ParseCoinsNormalized(fmt.Sprintf("%d%s", price.Mul(quantity).Ceil().RoundInt64(), "usei"))
		if err != nil {
			panic(err)
		}
		msg = &dextypes.MsgPlaceOrders{
			Creator:      sdk.AccAddress(key.PubKey().Address()).String(),
			Orders:       orderPlacements,
			ContractAddr: contract,
			Funds:        amount,
		}
	case FailureDexInvalid:
		price := config.PriceDistr.Sample()
		quantity := config.QuantityDistr.Sample()
		contract := config.ContractDistr.Sample()
		orderPlacements := generateDexOrderPlacements(config, key, msgPerTx, price, quantity)
		var amountUsei int64
		if rand.Float64() < 0.5 {
			amountUsei = 10000 * price.Mul(quantity).Ceil().RoundInt64()
		} else {
			amountUsei = 0
		}
		amount, err := sdk.ParseCoinsNormalized(fmt.Sprintf("%d%s", amountUsei, "usei"))
		if err != nil {
			panic(err)
		}
		msg = &dextypes.MsgPlaceOrders{
			Creator:      sdk.AccAddress(key.PubKey().Address()).String(),
			Orders:       orderPlacements,
			ContractAddr: contract,
			Funds:        amount,
		}

	default:
		fmt.Printf("Unrecognized message type %s", config.MessageType)
	}

	if strings.Contains(config.MessageType, "failure") {
		return msg, true
	}
	return msg, false
}

func sampleDexOrderType(config Config) (orderType dextypes.OrderType) {
	if config.MessageType == "failure_bank_malformed" {
		orderType = -1
	} else {
		msgType := config.MsgTypeDistr.SampleDexMsgs()
		switch msgType {
		case Limit:
			orderType = dextypes.OrderType_LIMIT
		case Market:
			orderType = dextypes.OrderType_MARKET
		default:
			panic(fmt.Sprintf("Unknown message type %s\n", msgType))
		}
	}
	return orderType
}

func generateDexOrderPlacements(config Config, key cryptotypes.PrivKey, msgPerTx uint64, price sdk.Dec, quantity sdk.Dec) (orderPlacements []*dextypes.Order) {
	orderType := sampleDexOrderType(config)

	var direction dextypes.PositionDirection
	if rand.Float64() < 0.5 {
		direction = dextypes.PositionDirection_LONG
	} else {
		direction = dextypes.PositionDirection_SHORT
	}

	contract := config.ContractDistr.Sample()
	for j := 0; j < int(msgPerTx); j++ {
		orderPlacements = append(orderPlacements, &dextypes.Order{
			Account:           sdk.AccAddress(key.PubKey().Address()).String(),
			ContractAddr:      contract,
			PositionDirection: direction,
			Price:             price.Quo(FromMili),
			Quantity:          quantity.Quo(FromMili),
			PriceDenom:        "SEI",
			AssetDenom:        "ATOM",
			OrderType:         orderType,
			Data:              VortexData,
		})
	}
	return orderPlacements
}

func generateOracleMessage(key cryptotypes.PrivKey) sdk.Msg {
	valAddr := sdk.ValAddress(key.PubKey().Address()).String()
	addr := sdk.AccAddress(key.PubKey().Address()).String()
	msg := &oracletypes.MsgAggregateExchangeRateVote{
		ExchangeRates: "1usei,2uatom",
		Feeder:        addr,
		Validator:     valAddr,
	}
	return msg
}

func (c *LoadTestClient) generateStakingMsg(delegatorAddr string, chosenValidator string, srcAddr string) sdk.Msg {
	// Randomly unbond, redelegate or delegate
	// However, if there are no delegations, do so first
	var msg sdk.Msg
	msgType := c.LoadTestConfig.MsgTypeDistr.SampleStakingMsgs()
	if _, ok := c.DelegationMap[delegatorAddr]; !ok || msgType == "delegate" || srcAddr == "" {
		msg = &stakingtypes.MsgDelegate{
			DelegatorAddress: delegatorAddr,
			ValidatorAddress: chosenValidator,
			Amount:           sdk.Coin{Denom: "usei", Amount: sdk.NewInt(1)},
		}
		c.DelegationMap[delegatorAddr] = map[string]int{}
		c.DelegationMap[delegatorAddr][chosenValidator] = 1
	} else {
		if msgType == "redelegate" {
			msg = &stakingtypes.MsgBeginRedelegate{
				DelegatorAddress:    delegatorAddr,
				ValidatorSrcAddress: srcAddr,
				ValidatorDstAddress: chosenValidator,
				Amount:              sdk.Coin{Denom: "usei", Amount: sdk.NewInt(1)},
			}
			c.DelegationMap[delegatorAddr][chosenValidator]++
		} else {
			msg = &stakingtypes.MsgUndelegate{
				DelegatorAddress: delegatorAddr,
				ValidatorAddress: srcAddr,
				Amount:           sdk.Coin{Denom: "usei", Amount: sdk.NewInt(1)},
			}
		}
		// Update delegation map
		c.DelegationMap[delegatorAddr][srcAddr]--
		if c.DelegationMap[delegatorAddr][srcAddr] == 0 {
			delete(c.DelegationMap, delegatorAddr)
		}
	}
	return msg
}

func getLastHeight() int {
	out, err := exec.Command("curl", "http://localhost:26657/blockchain").Output()
	if err != nil {
		panic(err)
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(out, &dat); err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(dat["last_height"].(string))
	if err != nil {
		panic(err)
	}
	return height
}

func GetDefaultConfigFilePath() string {
	pwd, _ := os.Getwd()
	return pwd + "/loadtest/config.json"
}

func ReadConfig(path string) Config {
	config := Config{}
	file, _ := os.ReadFile(path)
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}
	return config
}

func main() {
	configFilePath := flag.String("config-file", GetDefaultConfigFilePath(), "Path to the config.json file to use for this run")
	flag.Parse()

	config := ReadConfig(*configFilePath)
	fmt.Printf("Using config file: %s\n", *configFilePath)
	run(config)
}
