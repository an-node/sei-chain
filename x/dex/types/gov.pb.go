// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/gov.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type UpdateTickSizeProposal struct {
	Title        string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title" yaml:"title"`
	Description  string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description" yaml:"description"`
	TickSizeList []TickSize `protobuf:"bytes,3,rep,name=tickSizeList,proto3" json:"tick_size_list" yaml:"tick_size_list"`
}

func (m *UpdateTickSizeProposal) Reset()      { *m = UpdateTickSizeProposal{} }
func (*UpdateTickSizeProposal) ProtoMessage() {}
func (*UpdateTickSizeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab07ca1a96062d0, []int{0}
}
func (m *UpdateTickSizeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateTickSizeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateTickSizeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateTickSizeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTickSizeProposal.Merge(m, src)
}
func (m *UpdateTickSizeProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateTickSizeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTickSizeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTickSizeProposal proto.InternalMessageInfo

// AddAssetMetadataProposal is a gov Content type for adding a new asset
// to the dex module's asset list.
type AddAssetMetadataProposal struct {
	Title       string          `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	AssetList   []AssetMetadata `protobuf:"bytes,3,rep,name=assetList,proto3" json:"assetList" yaml:"asset_list"`
}

func (m *AddAssetMetadataProposal) Reset()      { *m = AddAssetMetadataProposal{} }
func (*AddAssetMetadataProposal) ProtoMessage() {}
func (*AddAssetMetadataProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab07ca1a96062d0, []int{1}
}
func (m *AddAssetMetadataProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddAssetMetadataProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddAssetMetadataProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddAssetMetadataProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAssetMetadataProposal.Merge(m, src)
}
func (m *AddAssetMetadataProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddAssetMetadataProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAssetMetadataProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddAssetMetadataProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateTickSizeProposal)(nil), "seiprotocol.seichain.dex.UpdateTickSizeProposal")
	proto.RegisterType((*AddAssetMetadataProposal)(nil), "seiprotocol.seichain.dex.AddAssetMetadataProposal")
}

func init() { proto.RegisterFile("dex/gov.proto", fileDescriptor_dab07ca1a96062d0) }

var fileDescriptor_dab07ca1a96062d0 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x8f, 0xda, 0x30,
	0x18, 0x87, 0x13, 0x50, 0x2b, 0x11, 0x68, 0xd5, 0xa6, 0x14, 0x05, 0x86, 0x18, 0x59, 0x6a, 0xcb,
	0x42, 0x22, 0xb5, 0x4b, 0xc5, 0x46, 0x16, 0x96, 0x56, 0xaa, 0xd2, 0x76, 0xe9, 0x42, 0x4d, 0x62,
	0x05, 0x8b, 0x80, 0x23, 0xec, 0xf6, 0x80, 0x4f, 0x70, 0xe3, 0x8d, 0x77, 0x1b, 0x1f, 0x87, 0x91,
	0xf1, 0x26, 0xeb, 0x04, 0xcb, 0x09, 0xdd, 0x94, 0x4f, 0x70, 0x8a, 0xc3, 0x9f, 0x44, 0x3a, 0xa4,
	0xdb, 0xec, 0xc7, 0x7e, 0xe3, 0xdf, 0xfb, 0xe4, 0xd5, 0x5e, 0xf9, 0x78, 0x66, 0x07, 0xf4, 0xbf,
	0x15, 0x4d, 0x29, 0xa7, 0xba, 0xc1, 0x30, 0x91, 0x2b, 0x8f, 0x86, 0x16, 0xc3, 0xc4, 0x1b, 0x22,
	0x32, 0xb1, 0x7c, 0x3c, 0x6b, 0x54, 0x03, 0x1a, 0x50, 0x79, 0x64, 0x27, 0xab, 0xf4, 0x7e, 0xa3,
	0x9a, 0x94, 0x23, 0xc6, 0x30, 0xef, 0x87, 0x84, 0xf1, 0x3d, 0x7d, 0x97, 0x50, 0x4e, 0xbc, 0x51,
	0x9f, 0x91, 0x05, 0x4e, 0x21, 0xbc, 0x29, 0x68, 0xb5, 0xdf, 0x91, 0x8f, 0x38, 0xfe, 0x45, 0xbc,
	0xd1, 0x4f, 0xb2, 0xc0, 0x3f, 0xa6, 0x34, 0xa2, 0x0c, 0x85, 0xba, 0xad, 0xbd, 0xe0, 0x84, 0x87,
	0xd8, 0x50, 0x9b, 0x6a, 0xab, 0xe4, 0xd4, 0x77, 0x02, 0xa4, 0x20, 0x16, 0xa0, 0x32, 0x47, 0xe3,
	0xb0, 0x03, 0xe5, 0x16, 0xba, 0x29, 0xd6, 0x7b, 0x5a, 0xd9, 0xc7, 0xcc, 0x9b, 0x92, 0x88, 0x13,
	0x3a, 0x31, 0x0a, 0xb2, 0xec, 0xc3, 0x4e, 0x80, 0x2c, 0x8e, 0x05, 0xd0, 0xd3, 0xe2, 0x0c, 0x84,
	0x6e, 0xf6, 0x8a, 0x7e, 0xa1, 0x55, 0xf8, 0x3e, 0xcd, 0x37, 0xc2, 0xb8, 0x51, 0x6c, 0x16, 0x5b,
	0xe5, 0xcf, 0xd0, 0x3a, 0xa7, 0xc1, 0x3a, 0x64, 0x77, 0xec, 0x95, 0x00, 0xca, 0x4e, 0x80, 0xd7,
	0xc7, 0x3e, 0xa5, 0x81, 0x58, 0x80, 0xf7, 0x87, 0xc4, 0x59, 0x0e, 0xdd, 0xdc, 0x43, 0x9d, 0xca,
	0xe5, 0x12, 0x28, 0xd7, 0x4b, 0xa0, 0xdc, 0x2f, 0x81, 0x02, 0x1f, 0x54, 0xcd, 0xe8, 0xfa, 0x7e,
	0x37, 0x11, 0xf9, 0x1d, 0x73, 0xe4, 0x23, 0x8e, 0x8e, 0x76, 0x3e, 0xe6, 0xed, 0xbc, 0x39, 0x27,
	0xe5, 0xeb, 0x53, 0x52, 0x6a, 0xcf, 0xb1, 0xf0, 0x57, 0x2b, 0xc9, 0x7f, 0x98, 0x51, 0xf0, 0xe9,
	0xbc, 0x82, 0x5c, 0x4a, 0xa7, 0x9e, 0x78, 0x88, 0x05, 0x78, 0x9b, 0x3e, 0x72, 0x9a, 0x05, 0xe8,
	0x9e, 0x3e, 0x9a, 0x6f, 0xd7, 0xe9, 0xad, 0x36, 0xa6, 0xba, 0xde, 0x98, 0xea, 0xdd, 0xc6, 0x54,
	0xaf, 0xb6, 0xa6, 0xb2, 0xde, 0x9a, 0xca, 0xed, 0xd6, 0x54, 0xfe, 0xb4, 0x03, 0xc2, 0x87, 0xff,
	0x06, 0x96, 0x47, 0xc7, 0x36, 0xc3, 0xa4, 0x7d, 0x48, 0x20, 0x37, 0x32, 0x82, 0x3d, 0xb3, 0xe5,
	0x74, 0xcd, 0x23, 0xcc, 0x06, 0x2f, 0xe5, 0xf9, 0x97, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62,
	0xf2, 0x29, 0x85, 0xc6, 0x02, 0x00, 0x00,
}

func (m *UpdateTickSizeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateTickSizeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateTickSizeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TickSizeList) > 0 {
		for iNdEx := len(m.TickSizeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TickSizeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AddAssetMetadataProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddAssetMetadataProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddAssetMetadataProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetList) > 0 {
		for iNdEx := len(m.AssetList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateTickSizeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.TickSizeList) > 0 {
		for _, e := range m.TickSizeList {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *AddAssetMetadataProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.AssetList) > 0 {
		for _, e := range m.AssetList {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateTickSizeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdateTickSizeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateTickSizeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickSizeList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TickSizeList = append(m.TickSizeList, TickSize{})
			if err := m.TickSizeList[len(m.TickSizeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AddAssetMetadataProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddAssetMetadataProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddAssetMetadataProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetList = append(m.AssetList, AssetMetadata{})
			if err := m.AssetList[len(m.AssetList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGov
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGov
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
