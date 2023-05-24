// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/parameter/ante_handler_param.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type AnteHandlerParam struct {
	MinCommissionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=minCommissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"minCommissionRate" yaml:"min_commission_rate"`
	MaxVotingPower    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=maxVotingPower,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"maxVotingPower" yaml:"max_voting_power"`
	MinSelfDelegation github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=minSelfDelegation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"minSelfDelegation" yaml:"min_self_delegation"`
}

func (m *AnteHandlerParam) Reset()         { *m = AnteHandlerParam{} }
func (m *AnteHandlerParam) String() string { return proto.CompactTextString(m) }
func (*AnteHandlerParam) ProtoMessage()    {}
func (*AnteHandlerParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ae17167f40c52b, []int{0}
}
func (m *AnteHandlerParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnteHandlerParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnteHandlerParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnteHandlerParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnteHandlerParam.Merge(m, src)
}
func (m *AnteHandlerParam) XXX_Size() int {
	return m.Size()
}
func (m *AnteHandlerParam) XXX_DiscardUnknown() {
	xxx_messageInfo_AnteHandlerParam.DiscardUnknown(m)
}

var xxx_messageInfo_AnteHandlerParam proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AnteHandlerParam)(nil), "elysnetwork.elys.parameter.AnteHandlerParam")
}

func init() {
	proto.RegisterFile("elys/parameter/ante_handler_param.proto", fileDescriptor_61ae17167f40c52b)
}

var fileDescriptor_61ae17167f40c52b = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x6a, 0xeb, 0x30,
	0x14, 0xc6, 0xed, 0x5c, 0xb8, 0x70, 0x3d, 0x5c, 0xda, 0x50, 0x68, 0xf0, 0xe0, 0x14, 0x0f, 0x6d,
	0x97, 0x58, 0x84, 0x6e, 0xdd, 0x9a, 0x66, 0x48, 0xa0, 0x43, 0x48, 0xa1, 0x43, 0x17, 0x23, 0x3b,
	0x27, 0x8e, 0x88, 0xa5, 0xe3, 0x4a, 0x6a, 0xfe, 0xf4, 0x29, 0xfa, 0x58, 0x19, 0x33, 0x96, 0x0e,
	0xa1, 0x24, 0x6f, 0xd0, 0xbd, 0x50, 0x2c, 0x9b, 0x90, 0xb4, 0x53, 0x26, 0x1d, 0x49, 0xe7, 0x7c,
	0x3f, 0x3e, 0xbe, 0xe3, 0x5c, 0x40, 0x3a, 0x57, 0x24, 0xa3, 0x92, 0x72, 0xd0, 0x20, 0x09, 0x15,
	0x1a, 0xc2, 0x11, 0x15, 0x83, 0x14, 0x64, 0x68, 0x9e, 0x83, 0x4c, 0xa2, 0xc6, 0xaa, 0x9b, 0x37,
	0x0a, 0xd0, 0x53, 0x94, 0xe3, 0x20, 0xaf, 0x83, 0xed, 0x90, 0x7b, 0x92, 0x60, 0x82, 0xa6, 0x8d,
	0xe4, 0x55, 0x31, 0xe1, 0x7a, 0x31, 0x2a, 0x8e, 0x8a, 0x44, 0x54, 0x01, 0x99, 0x34, 0x23, 0xd0,
	0xb4, 0x49, 0x62, 0x64, 0xa2, 0xf8, 0xf7, 0xbf, 0x2a, 0xce, 0xd1, 0x8d, 0xd0, 0xd0, 0x29, 0x68,
	0xbd, 0x5c, 0xae, 0xfa, 0xe2, 0x1c, 0x73, 0x26, 0x6e, 0x91, 0x73, 0xa6, 0x14, 0x43, 0xd1, 0xa7,
	0x1a, 0x6a, 0xf6, 0x99, 0x7d, 0xf9, 0xaf, 0x75, 0xb7, 0x58, 0xd5, 0xad, 0xf7, 0x55, 0xfd, 0x3c,
	0x61, 0x7a, 0xf4, 0x1c, 0x05, 0x31, 0x72, 0x52, 0x22, 0x8a, 0xa3, 0xa1, 0x06, 0x63, 0xa2, 0xe7,
	0x19, 0xa8, 0xa0, 0x0d, 0xf1, 0xe7, 0xaa, 0xee, 0xce, 0x29, 0x4f, 0xaf, 0x7d, 0xce, 0x44, 0x18,
	0x6f, 0x15, 0x43, 0x49, 0x35, 0xf8, 0xfd, 0xdf, 0x98, 0xea, 0x93, 0xf3, 0x9f, 0xd3, 0xd9, 0x03,
	0x6a, 0x26, 0x92, 0x1e, 0x4e, 0x41, 0xd6, 0x2a, 0x06, 0xdc, 0x3d, 0x18, 0x7c, 0x5a, 0x82, 0xe9,
	0x2c, 0x9c, 0x18, 0xb9, 0x30, 0xcb, 0xf5, 0xfc, 0xfe, 0x0f, 0x40, 0x69, 0xf7, 0x1e, 0xd2, 0x61,
	0x1b, 0x52, 0x48, 0xa8, 0x66, 0x28, 0x6a, 0x7f, 0x0e, 0xb6, 0xdb, 0x15, 0x7a, 0xdf, 0xae, 0x82,
	0x74, 0x18, 0x0e, 0xb6, 0x92, 0x85, 0xdd, 0x7d, 0x4c, 0xab, 0xb3, 0x58, 0x7b, 0xf6, 0x72, 0xed,
	0xd9, 0x1f, 0x6b, 0xcf, 0x7e, 0xdd, 0x78, 0xd6, 0x72, 0xe3, 0x59, 0x6f, 0x1b, 0xcf, 0x7a, 0x0c,
	0x76, 0x90, 0x79, 0xd4, 0x8d, 0x32, 0x77, 0x73, 0x21, 0xb3, 0x9d, 0x75, 0x31, 0xf8, 0xe8, 0xaf,
	0x09, 0xf4, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x83, 0x82, 0x55, 0xec, 0x4d, 0x02, 0x00, 0x00,
}

func (m *AnteHandlerParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnteHandlerParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnteHandlerParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinSelfDelegation.Size()
		i -= size
		if _, err := m.MinSelfDelegation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnteHandlerParam(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MaxVotingPower.Size()
		i -= size
		if _, err := m.MaxVotingPower.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnteHandlerParam(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinCommissionRate.Size()
		i -= size
		if _, err := m.MinCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnteHandlerParam(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAnteHandlerParam(dAtA []byte, offset int, v uint64) int {
	offset -= sovAnteHandlerParam(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AnteHandlerParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinCommissionRate.Size()
	n += 1 + l + sovAnteHandlerParam(uint64(l))
	l = m.MaxVotingPower.Size()
	n += 1 + l + sovAnteHandlerParam(uint64(l))
	l = m.MinSelfDelegation.Size()
	n += 1 + l + sovAnteHandlerParam(uint64(l))
	return n
}

func sovAnteHandlerParam(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnteHandlerParam(x uint64) (n int) {
	return sovAnteHandlerParam(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnteHandlerParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnteHandlerParam
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
			return fmt.Errorf("proto: AnteHandlerParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnteHandlerParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnteHandlerParam
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
				return ErrInvalidLengthAnteHandlerParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnteHandlerParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxVotingPower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnteHandlerParam
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
				return ErrInvalidLengthAnteHandlerParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnteHandlerParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxVotingPower.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSelfDelegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnteHandlerParam
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
				return ErrInvalidLengthAnteHandlerParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnteHandlerParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSelfDelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnteHandlerParam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnteHandlerParam
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
func skipAnteHandlerParam(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnteHandlerParam
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
					return 0, ErrIntOverflowAnteHandlerParam
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
					return 0, ErrIntOverflowAnteHandlerParam
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
				return 0, ErrInvalidLengthAnteHandlerParam
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAnteHandlerParam
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAnteHandlerParam
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAnteHandlerParam        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnteHandlerParam          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAnteHandlerParam = fmt.Errorf("proto: unexpected end of group")
)