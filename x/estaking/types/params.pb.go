// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/estaking/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type LegacyParams struct {
	StakeIncentives *IncentiveInfo `protobuf:"bytes,1,opt,name=stake_incentives,json=stakeIncentives,proto3" json:"stake_incentives,omitempty"`
	EdenCommitVal   string         `protobuf:"bytes,2,opt,name=eden_commit_val,json=edenCommitVal,proto3" json:"eden_commit_val,omitempty"`
	EdenbCommitVal  string         `protobuf:"bytes,3,opt,name=edenb_commit_val,json=edenbCommitVal,proto3" json:"edenb_commit_val,omitempty"`
	// Maximum eden reward apr for stakers - [0 - 0.3]
	MaxEdenRewardAprStakers cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=max_eden_reward_apr_stakers,json=maxEdenRewardAprStakers,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"max_eden_reward_apr_stakers"`
	EdenBoostApr            cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=eden_boost_apr,json=edenBoostApr,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"eden_boost_apr"`
	// Tracking dex rewards given to stakers
	DexRewardsStakers DexRewardsTracker `protobuf:"bytes,6,opt,name=dex_rewards_stakers,json=dexRewardsStakers,proto3" json:"dex_rewards_stakers"`
}

func (m *LegacyParams) Reset()         { *m = LegacyParams{} }
func (m *LegacyParams) String() string { return proto.CompactTextString(m) }
func (*LegacyParams) ProtoMessage()    {}
func (*LegacyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_66041162e1ecb63b, []int{0}
}
func (m *LegacyParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LegacyParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LegacyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyParams.Merge(m, src)
}
func (m *LegacyParams) XXX_Size() int {
	return m.Size()
}
func (m *LegacyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyParams.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyParams proto.InternalMessageInfo

func (m *LegacyParams) GetStakeIncentives() *IncentiveInfo {
	if m != nil {
		return m.StakeIncentives
	}
	return nil
}

func (m *LegacyParams) GetEdenCommitVal() string {
	if m != nil {
		return m.EdenCommitVal
	}
	return ""
}

func (m *LegacyParams) GetEdenbCommitVal() string {
	if m != nil {
		return m.EdenbCommitVal
	}
	return ""
}

func (m *LegacyParams) GetDexRewardsStakers() DexRewardsTracker {
	if m != nil {
		return m.DexRewardsStakers
	}
	return DexRewardsTracker{}
}

// Params defines the parameters for the module.
type Params struct {
	StakeIncentives *IncentiveInfo `protobuf:"bytes,1,opt,name=stake_incentives,json=stakeIncentives,proto3" json:"stake_incentives,omitempty"`
	EdenCommitVal   string         `protobuf:"bytes,2,opt,name=eden_commit_val,json=edenCommitVal,proto3" json:"eden_commit_val,omitempty"`
	EdenbCommitVal  string         `protobuf:"bytes,3,opt,name=edenb_commit_val,json=edenbCommitVal,proto3" json:"edenb_commit_val,omitempty"`
	// Maximum eden reward apr for stakers - [0 - 0.3]
	MaxEdenRewardAprStakers cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=max_eden_reward_apr_stakers,json=maxEdenRewardAprStakers,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"max_eden_reward_apr_stakers"`
	EdenBoostApr            cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=eden_boost_apr,json=edenBoostApr,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"eden_boost_apr"`
	// Tracking dex rewards given to stakers
	DexRewardsStakers              DexRewardsTracker           `protobuf:"bytes,6,opt,name=dex_rewards_stakers,json=dexRewardsStakers,proto3" json:"dex_rewards_stakers"`
	ProviderVestingEpochIdentifier string                      `protobuf:"bytes,7,opt,name=provider_vesting_epoch_identifier,json=providerVestingEpochIdentifier,proto3" json:"provider_vesting_epoch_identifier,omitempty"`
	ProviderStakingRewardsPortion  cosmossdk_io_math.LegacyDec `protobuf:"bytes,8,opt,name=provider_staking_rewards_portion,json=providerStakingRewardsPortion,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"provider_staking_rewards_portion"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_66041162e1ecb63b, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetStakeIncentives() *IncentiveInfo {
	if m != nil {
		return m.StakeIncentives
	}
	return nil
}

func (m *Params) GetEdenCommitVal() string {
	if m != nil {
		return m.EdenCommitVal
	}
	return ""
}

func (m *Params) GetEdenbCommitVal() string {
	if m != nil {
		return m.EdenbCommitVal
	}
	return ""
}

func (m *Params) GetDexRewardsStakers() DexRewardsTracker {
	if m != nil {
		return m.DexRewardsStakers
	}
	return DexRewardsTracker{}
}

func (m *Params) GetProviderVestingEpochIdentifier() string {
	if m != nil {
		return m.ProviderVestingEpochIdentifier
	}
	return ""
}

func init() {
	proto.RegisterType((*LegacyParams)(nil), "elys.estaking.LegacyParams")
	proto.RegisterType((*Params)(nil), "elys.estaking.Params")
}

func init() { proto.RegisterFile("elys/estaking/params.proto", fileDescriptor_66041162e1ecb63b) }

var fileDescriptor_66041162e1ecb63b = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x8d, 0x5f, 0xdb, 0x3c, 0x18, 0x5a, 0x5a, 0x0c, 0x12, 0x26, 0xa5, 0x6e, 0xe8, 0x02, 0xb2,
	0x89, 0x2d, 0xe0, 0x0b, 0x1a, 0x52, 0x55, 0x91, 0x58, 0x54, 0x2e, 0x0a, 0x12, 0x9b, 0xd1, 0xc4,
	0xbe, 0x75, 0x46, 0x89, 0x3d, 0xa3, 0x99, 0x21, 0x75, 0xf8, 0x06, 0x16, 0x7c, 0x0c, 0x1f, 0x51,
	0x89, 0x4d, 0xc5, 0x0a, 0xb1, 0xa8, 0x50, 0xf2, 0x23, 0x68, 0x66, 0x6c, 0xb7, 0x61, 0xd9, 0x1d,
	0x12, 0xbb, 0x99, 0x7b, 0xce, 0x9c, 0x7b, 0xee, 0xb9, 0xd2, 0xa0, 0x16, 0x4c, 0xe7, 0x32, 0x04,
	0xa9, 0xc8, 0x84, 0xe6, 0x69, 0xc8, 0x89, 0x20, 0x99, 0x0c, 0xb8, 0x60, 0x8a, 0xb9, 0x5b, 0x1a,
	0x0b, 0x2a, 0xac, 0xf5, 0x28, 0x65, 0x29, 0x33, 0x48, 0xa8, 0x4f, 0x96, 0xd4, 0xda, 0x5b, 0x15,
	0xa0, 0x79, 0x0c, 0xb9, 0xa2, 0x33, 0x28, 0xe1, 0x17, 0xab, 0x70, 0x02, 0x05, 0x16, 0x70, 0x4e,
	0x44, 0x22, 0xb1, 0x12, 0x24, 0x9e, 0x80, 0x28, 0x89, 0x4f, 0x62, 0x26, 0x33, 0x26, 0xb1, 0x6d,
	0x60, 0x2f, 0x16, 0x3a, 0xf8, 0xb6, 0x86, 0x36, 0xdf, 0x42, 0x4a, 0xe2, 0xf9, 0x89, 0xb1, 0xe7,
	0x1e, 0xa3, 0x1d, 0x2d, 0x08, 0xb8, 0xee, 0x26, 0x3d, 0xa7, 0xed, 0x74, 0xee, 0xbd, 0x7a, 0x1a,
	0xac, 0x78, 0x0e, 0x06, 0x15, 0x61, 0x90, 0x9f, 0xb1, 0x68, 0xdb, 0xbc, 0xaa, 0x6b, 0xd2, 0x7d,
	0x8e, 0xb6, 0x21, 0x81, 0x1c, 0xc7, 0x2c, 0xcb, 0xa8, 0xc2, 0x33, 0x32, 0xf5, 0xfe, 0x6b, 0x3b,
	0x9d, 0xbb, 0xd1, 0x96, 0x2e, 0xbf, 0x31, 0xd5, 0x21, 0x99, 0xba, 0x1d, 0xb4, 0xa3, 0x0b, 0xa3,
	0x9b, 0xc4, 0x35, 0x43, 0xbc, 0x6f, 0xea, 0xd7, 0x4c, 0x86, 0x76, 0x33, 0x52, 0x60, 0xa3, 0x6a,
	0x07, 0xc5, 0x84, 0x0b, 0x6c, 0x1a, 0x0b, 0xe9, 0xad, 0xeb, 0x47, 0xbd, 0x97, 0x17, 0x57, 0xfb,
	0x8d, 0x9f, 0x57, 0xfb, 0xbb, 0x76, 0x4c, 0x99, 0x4c, 0x02, 0xca, 0xc2, 0x8c, 0xa8, 0x71, 0x60,
	0xe7, 0xec, 0x43, 0xfc, 0xfd, 0x6b, 0x17, 0x95, 0x29, 0xf4, 0x21, 0x8e, 0x1e, 0x67, 0xa4, 0x38,
	0x4a, 0x20, 0x8f, 0x8c, 0xe6, 0x21, 0x17, 0xa7, 0x56, 0xd1, 0x7d, 0x8f, 0x8c, 0x05, 0x3c, 0x62,
	0x4c, 0x2a, 0xdd, 0xcb, 0xdb, 0xb8, 0x6d, 0x8f, 0x4d, 0x2d, 0xd4, 0xd3, 0x3a, 0x87, 0x5c, 0xb8,
	0x43, 0xf4, 0xf0, 0xe6, 0xb6, 0xaa, 0x09, 0x9a, 0x26, 0xe7, 0xf6, 0x1f, 0x39, 0xf7, 0xa1, 0xb0,
	0xce, 0xe4, 0x3b, 0xbb, 0xd5, 0xde, 0xba, 0xee, 0x1f, 0x3d, 0x48, 0x6a, 0xa0, 0x34, 0x7c, 0xf0,
	0x79, 0x03, 0x35, 0xff, 0xed, 0xf1, 0xaf, 0xdf, 0xa3, 0x3b, 0x40, 0xcf, 0xb8, 0x60, 0x33, 0x9a,
	0x80, 0xc0, 0x33, 0x90, 0x8a, 0xe6, 0x29, 0x06, 0xce, 0xe2, 0x31, 0xa6, 0x89, 0xde, 0xcb, 0x19,
	0x05, 0xe1, 0xfd, 0x6f, 0xc2, 0xf5, 0x2b, 0xe2, 0xd0, 0xf2, 0x8e, 0x34, 0x6d, 0x50, 0xb3, 0xdc,
	0x4f, 0xa8, 0x5d, 0x4b, 0x95, 0x4e, 0x6a, 0xbf, 0x9c, 0x09, 0x45, 0x59, 0xee, 0xdd, 0xb9, 0x6d,
	0x1a, 0x7b, 0x95, 0xf4, 0xa9, 0x55, 0x2e, 0xc7, 0x38, 0xb1, 0xba, 0xbd, 0xe3, 0x8b, 0x85, 0xef,
	0x5c, 0x2e, 0x7c, 0xe7, 0xd7, 0xc2, 0x77, 0xbe, 0x2c, 0xfd, 0xc6, 0xe5, 0xd2, 0x6f, 0xfc, 0x58,
	0xfa, 0x8d, 0x0f, 0xdd, 0x94, 0xaa, 0xf1, 0xc7, 0x51, 0x10, 0xb3, 0x2c, 0xd4, 0x29, 0x75, 0x73,
	0x50, 0xe7, 0x4c, 0x4c, 0xcc, 0x25, 0x2c, 0xae, 0x3f, 0x35, 0x35, 0xe7, 0x20, 0x47, 0x4d, 0xf3,
	0x59, 0xbd, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xf5, 0x2c, 0x37, 0x52, 0x05, 0x00, 0x00,
}

func (m *LegacyParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LegacyParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DexRewardsStakers.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.EdenBoostApr.Size()
		i -= size
		if _, err := m.EdenBoostApr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MaxEdenRewardAprStakers.Size()
		i -= size
		if _, err := m.MaxEdenRewardAprStakers.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.EdenbCommitVal) > 0 {
		i -= len(m.EdenbCommitVal)
		copy(dAtA[i:], m.EdenbCommitVal)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EdenbCommitVal)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EdenCommitVal) > 0 {
		i -= len(m.EdenCommitVal)
		copy(dAtA[i:], m.EdenCommitVal)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EdenCommitVal)))
		i--
		dAtA[i] = 0x12
	}
	if m.StakeIncentives != nil {
		{
			size, err := m.StakeIncentives.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ProviderStakingRewardsPortion.Size()
		i -= size
		if _, err := m.ProviderStakingRewardsPortion.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.ProviderVestingEpochIdentifier) > 0 {
		i -= len(m.ProviderVestingEpochIdentifier)
		copy(dAtA[i:], m.ProviderVestingEpochIdentifier)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ProviderVestingEpochIdentifier)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size, err := m.DexRewardsStakers.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.EdenBoostApr.Size()
		i -= size
		if _, err := m.EdenBoostApr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MaxEdenRewardAprStakers.Size()
		i -= size
		if _, err := m.MaxEdenRewardAprStakers.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.EdenbCommitVal) > 0 {
		i -= len(m.EdenbCommitVal)
		copy(dAtA[i:], m.EdenbCommitVal)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EdenbCommitVal)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EdenCommitVal) > 0 {
		i -= len(m.EdenCommitVal)
		copy(dAtA[i:], m.EdenCommitVal)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EdenCommitVal)))
		i--
		dAtA[i] = 0x12
	}
	if m.StakeIncentives != nil {
		{
			size, err := m.StakeIncentives.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LegacyParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StakeIncentives != nil {
		l = m.StakeIncentives.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.EdenCommitVal)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.EdenbCommitVal)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MaxEdenRewardAprStakers.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.EdenBoostApr.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DexRewardsStakers.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StakeIncentives != nil {
		l = m.StakeIncentives.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.EdenCommitVal)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.EdenbCommitVal)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MaxEdenRewardAprStakers.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.EdenBoostApr.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DexRewardsStakers.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.ProviderVestingEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.ProviderStakingRewardsPortion.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LegacyParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: LegacyParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LegacyParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeIncentives", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StakeIncentives == nil {
				m.StakeIncentives = &IncentiveInfo{}
			}
			if err := m.StakeIncentives.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenCommitVal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EdenCommitVal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenbCommitVal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EdenbCommitVal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEdenRewardAprStakers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxEdenRewardAprStakers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenBoostApr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EdenBoostApr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DexRewardsStakers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DexRewardsStakers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeIncentives", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StakeIncentives == nil {
				m.StakeIncentives = &IncentiveInfo{}
			}
			if err := m.StakeIncentives.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenCommitVal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EdenCommitVal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenbCommitVal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EdenbCommitVal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEdenRewardAprStakers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxEdenRewardAprStakers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenBoostApr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EdenBoostApr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DexRewardsStakers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DexRewardsStakers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderVestingEpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderVestingEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderStakingRewardsPortion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProviderStakingRewardsPortion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
