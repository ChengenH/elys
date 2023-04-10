// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/oracle/proposal.proto

package types

import (
	fmt "fmt"
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

type ProposalAddAssetInfo struct {
	Title         string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Denom         string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	Display       string `protobuf:"bytes,4,opt,name=display,proto3" json:"display,omitempty"`
	BandTicker    string `protobuf:"bytes,5,opt,name=bandTicker,proto3" json:"bandTicker,omitempty"`
	BinanceTicker string `protobuf:"bytes,6,opt,name=binanceTicker,proto3" json:"binanceTicker,omitempty"`
	OsmosisTicker string `protobuf:"bytes,7,opt,name=osmosisTicker,proto3" json:"osmosisTicker,omitempty"`
}

func (m *ProposalAddAssetInfo) Reset()         { *m = ProposalAddAssetInfo{} }
func (m *ProposalAddAssetInfo) String() string { return proto.CompactTextString(m) }
func (*ProposalAddAssetInfo) ProtoMessage()    {}
func (*ProposalAddAssetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f712a27b3e5fed, []int{0}
}
func (m *ProposalAddAssetInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalAddAssetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalAddAssetInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalAddAssetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalAddAssetInfo.Merge(m, src)
}
func (m *ProposalAddAssetInfo) XXX_Size() int {
	return m.Size()
}
func (m *ProposalAddAssetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalAddAssetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalAddAssetInfo proto.InternalMessageInfo

func (m *ProposalAddAssetInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ProposalAddAssetInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProposalAddAssetInfo) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *ProposalAddAssetInfo) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *ProposalAddAssetInfo) GetBandTicker() string {
	if m != nil {
		return m.BandTicker
	}
	return ""
}

func (m *ProposalAddAssetInfo) GetBinanceTicker() string {
	if m != nil {
		return m.BinanceTicker
	}
	return ""
}

func (m *ProposalAddAssetInfo) GetOsmosisTicker() string {
	if m != nil {
		return m.OsmosisTicker
	}
	return ""
}

type ProposalRemoveAssetInfo struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Denom       string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *ProposalRemoveAssetInfo) Reset()         { *m = ProposalRemoveAssetInfo{} }
func (m *ProposalRemoveAssetInfo) String() string { return proto.CompactTextString(m) }
func (*ProposalRemoveAssetInfo) ProtoMessage()    {}
func (*ProposalRemoveAssetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f712a27b3e5fed, []int{1}
}
func (m *ProposalRemoveAssetInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalRemoveAssetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalRemoveAssetInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalRemoveAssetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalRemoveAssetInfo.Merge(m, src)
}
func (m *ProposalRemoveAssetInfo) XXX_Size() int {
	return m.Size()
}
func (m *ProposalRemoveAssetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalRemoveAssetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalRemoveAssetInfo proto.InternalMessageInfo

func (m *ProposalRemoveAssetInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ProposalRemoveAssetInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProposalRemoveAssetInfo) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*ProposalAddAssetInfo)(nil), "elysnetwork.elys.oracle.ProposalAddAssetInfo")
	proto.RegisterType((*ProposalRemoveAssetInfo)(nil), "elysnetwork.elys.oracle.ProposalRemoveAssetInfo")
}

func init() { proto.RegisterFile("elys/oracle/proposal.proto", fileDescriptor_02f712a27b3e5fed) }

var fileDescriptor_02f712a27b3e5fed = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x63, 0xa0, 0xad, 0x30, 0x62, 0xb1, 0x2a, 0xd5, 0x62, 0xb0, 0xaa, 0x8a, 0x01, 0x09,
	0x91, 0x0c, 0x3c, 0x41, 0x91, 0x18, 0xd8, 0x50, 0xc5, 0xc4, 0x96, 0xc4, 0x97, 0x62, 0x35, 0xf1,
	0xb5, 0x6c, 0xf3, 0x93, 0xb7, 0xe0, 0xb1, 0x18, 0x3b, 0x32, 0xa2, 0x64, 0xe3, 0x29, 0x50, 0x9c,
	0x44, 0x6a, 0x1f, 0xa0, 0x9b, 0xcf, 0x77, 0x3e, 0x0f, 0x57, 0x87, 0x5e, 0x40, 0x51, 0xb9, 0x04,
	0x6d, 0x9a, 0x17, 0x90, 0x18, 0x8b, 0x06, 0x5d, 0x5a, 0xc4, 0xc6, 0xa2, 0x47, 0x36, 0x6b, 0x3b,
	0x0d, 0xfe, 0x03, 0xed, 0x26, 0x6e, 0xdf, 0x71, 0xe7, 0x2d, 0xfe, 0x08, 0x9d, 0x3e, 0xf6, 0xee,
	0x52, 0xca, 0xa5, 0x73, 0xe0, 0x1f, 0xf4, 0x0b, 0xb2, 0x29, 0x1d, 0x79, 0xe5, 0x0b, 0xe0, 0x64,
	0x4e, 0xae, 0x4e, 0x57, 0x5d, 0x60, 0x73, 0x7a, 0x26, 0xc1, 0xe5, 0x56, 0x19, 0xaf, 0x50, 0xf3,
	0xa3, 0xd0, 0xed, 0xa2, 0xf6, 0x9f, 0x04, 0x8d, 0x25, 0x3f, 0xee, 0xfe, 0x85, 0xc0, 0x38, 0x9d,
	0x48, 0xe5, 0x4c, 0x91, 0x56, 0xfc, 0x24, 0xf0, 0x21, 0x32, 0x41, 0x69, 0x96, 0x6a, 0xf9, 0xa4,
	0xf2, 0x0d, 0x58, 0x3e, 0x0a, 0xe5, 0x0e, 0x61, 0x97, 0xf4, 0x3c, 0x53, 0x3a, 0xd5, 0x39, 0xf4,
	0xca, 0x38, 0x28, 0xfb, 0xb0, 0xb5, 0xd0, 0x95, 0xe8, 0x94, 0xeb, 0xad, 0x49, 0x67, 0xed, 0xc1,
	0xc5, 0x9a, 0xce, 0x86, 0x5b, 0x57, 0x50, 0xe2, 0x3b, 0x1c, 0xe8, 0xdc, 0xbb, 0xfb, 0xef, 0x5a,
	0x90, 0x6d, 0x2d, 0xc8, 0x6f, 0x2d, 0xc8, 0x57, 0x23, 0xa2, 0x6d, 0x23, 0xa2, 0x9f, 0x46, 0x44,
	0xcf, 0xd7, 0x6b, 0xe5, 0x5f, 0xdf, 0xb2, 0x38, 0xc7, 0x32, 0x69, 0x77, 0xb8, 0xe9, 0x47, 0x09,
	0x21, 0xf9, 0x1c, 0xe6, 0xf3, 0x95, 0x01, 0x97, 0x8d, 0xc3, 0x78, 0xb7, 0xff, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xed, 0xd8, 0xaa, 0x81, 0xda, 0x01, 0x00, 0x00,
}

func (m *ProposalAddAssetInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalAddAssetInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalAddAssetInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OsmosisTicker) > 0 {
		i -= len(m.OsmosisTicker)
		copy(dAtA[i:], m.OsmosisTicker)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.OsmosisTicker)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.BinanceTicker) > 0 {
		i -= len(m.BinanceTicker)
		copy(dAtA[i:], m.BinanceTicker)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.BinanceTicker)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.BandTicker) > 0 {
		i -= len(m.BandTicker)
		copy(dAtA[i:], m.BandTicker)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.BandTicker)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Display) > 0 {
		i -= len(m.Display)
		copy(dAtA[i:], m.Display)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Display)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProposalRemoveAssetInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalRemoveAssetInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalRemoveAssetInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposalAddAssetInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Display)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.BandTicker)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.BinanceTicker)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.OsmosisTicker)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *ProposalRemoveAssetInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposalAddAssetInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: ProposalAddAssetInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalAddAssetInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
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
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Display", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Display = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandTicker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BandTicker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BinanceTicker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BinanceTicker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OsmosisTicker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OsmosisTicker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *ProposalRemoveAssetInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: ProposalRemoveAssetInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalRemoveAssetInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
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
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)