// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lfb/slashing/v1beta1/genesis.proto

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

// GenesisState defines the slashing module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of related to deposit.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// signing_infos represents a map between validator addresses and their
	// signing infos.
	SigningInfos []SigningInfo `protobuf:"bytes,2,rep,name=signing_infos,json=signingInfos,proto3" json:"signing_infos" yaml:"signing_infos"`
	// signing_infos represents a map between validator addresses and their
	// missed blocks.
	MissedBlocks []ValidatorMissedBlocks `protobuf:"bytes,3,rep,name=missed_blocks,json=missedBlocks,proto3" json:"missed_blocks" yaml:"missed_blocks"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b579d05c9c8e39d7, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSigningInfos() []SigningInfo {
	if m != nil {
		return m.SigningInfos
	}
	return nil
}

func (m *GenesisState) GetMissedBlocks() []ValidatorMissedBlocks {
	if m != nil {
		return m.MissedBlocks
	}
	return nil
}

// SigningInfo stores validator signing info of corresponding address.
type SigningInfo struct {
	// address is the validator address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// validator_signing_info represents the signing info of this validator.
	ValidatorSigningInfo ValidatorSigningInfo `protobuf:"bytes,2,opt,name=validator_signing_info,json=validatorSigningInfo,proto3" json:"validator_signing_info" yaml:"validator_signing_info"`
}

func (m *SigningInfo) Reset()         { *m = SigningInfo{} }
func (m *SigningInfo) String() string { return proto.CompactTextString(m) }
func (*SigningInfo) ProtoMessage()    {}
func (*SigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b579d05c9c8e39d7, []int{1}
}
func (m *SigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningInfo.Merge(m, src)
}
func (m *SigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *SigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SigningInfo proto.InternalMessageInfo

func (m *SigningInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SigningInfo) GetValidatorSigningInfo() ValidatorSigningInfo {
	if m != nil {
		return m.ValidatorSigningInfo
	}
	return ValidatorSigningInfo{}
}

// ValidatorMissedBlocks contains array of missed blocks of corresponding
// address.
type ValidatorMissedBlocks struct {
	// address is the validator address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// missed_blocks is an array of missed blocks by the validator.
	MissedBlocks []MissedBlock `protobuf:"bytes,2,rep,name=missed_blocks,json=missedBlocks,proto3" json:"missed_blocks" yaml:"missed_blocks"`
}

func (m *ValidatorMissedBlocks) Reset()         { *m = ValidatorMissedBlocks{} }
func (m *ValidatorMissedBlocks) String() string { return proto.CompactTextString(m) }
func (*ValidatorMissedBlocks) ProtoMessage()    {}
func (*ValidatorMissedBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_b579d05c9c8e39d7, []int{2}
}
func (m *ValidatorMissedBlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorMissedBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorMissedBlocks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorMissedBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorMissedBlocks.Merge(m, src)
}
func (m *ValidatorMissedBlocks) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorMissedBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorMissedBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorMissedBlocks proto.InternalMessageInfo

func (m *ValidatorMissedBlocks) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorMissedBlocks) GetMissedBlocks() []MissedBlock {
	if m != nil {
		return m.MissedBlocks
	}
	return nil
}

// MissedBlock contains height and missed status as boolean.
type MissedBlock struct {
	// index is the height at which the block was missed.
	Index int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// missed is the missed status.
	Missed bool `protobuf:"varint,2,opt,name=missed,proto3" json:"missed,omitempty"`
}

func (m *MissedBlock) Reset()         { *m = MissedBlock{} }
func (m *MissedBlock) String() string { return proto.CompactTextString(m) }
func (*MissedBlock) ProtoMessage()    {}
func (*MissedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_b579d05c9c8e39d7, []int{3}
}
func (m *MissedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MissedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MissedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MissedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissedBlock.Merge(m, src)
}
func (m *MissedBlock) XXX_Size() int {
	return m.Size()
}
func (m *MissedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_MissedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_MissedBlock proto.InternalMessageInfo

func (m *MissedBlock) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *MissedBlock) GetMissed() bool {
	if m != nil {
		return m.Missed
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lfb.slashing.v1beta1.GenesisState")
	proto.RegisterType((*SigningInfo)(nil), "lfb.slashing.v1beta1.SigningInfo")
	proto.RegisterType((*ValidatorMissedBlocks)(nil), "lfb.slashing.v1beta1.ValidatorMissedBlocks")
	proto.RegisterType((*MissedBlock)(nil), "lfb.slashing.v1beta1.MissedBlock")
}

func init() {
	proto.RegisterFile("lfb/slashing/v1beta1/genesis.proto", fileDescriptor_b579d05c9c8e39d7)
}

var fileDescriptor_b579d05c9c8e39d7 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbf, 0x8e, 0xd3, 0x30,
	0x00, 0xc6, 0xe3, 0x16, 0x0a, 0x38, 0xbd, 0xc5, 0x0a, 0xa7, 0xe8, 0x74, 0xe4, 0x0e, 0x23, 0xa4,
	0x0a, 0x44, 0xa2, 0x3b, 0xb6, 0x63, 0xcb, 0x82, 0x18, 0x90, 0x50, 0x2a, 0x31, 0xb0, 0x54, 0x4e,
	0xed, 0xa4, 0x56, 0x1d, 0xbb, 0x8a, 0x43, 0xd5, 0x3e, 0x01, 0x2b, 0x1b, 0x2b, 0xaf, 0xc0, 0x5b,
	0x74, 0xec, 0xc8, 0x54, 0xa1, 0xf6, 0x0d, 0x78, 0x02, 0x54, 0x27, 0xa5, 0x69, 0x15, 0x15, 0xb1,
	0xc5, 0xd2, 0xef, 0xfb, 0xe3, 0x2f, 0x09, 0xc4, 0x22, 0x89, 0x03, 0x2d, 0x88, 0x1e, 0x71, 0x99,
	0x06, 0xd3, 0x9b, 0x98, 0x15, 0xe4, 0x26, 0x48, 0x99, 0x64, 0x9a, 0x6b, 0x7f, 0x92, 0xab, 0x42,
	0x21, 0x47, 0x24, 0xb1, 0xbf, 0x63, 0xfc, 0x8a, 0xb9, 0x70, 0x52, 0x95, 0x2a, 0x03, 0x04, 0xdb,
	0xa7, 0x92, 0xbd, 0x78, 0xd6, 0xe8, 0xf7, 0x57, 0x6c, 0x20, 0xfc, 0xbd, 0x05, 0xbb, 0x6f, 0xcb,
	0x88, 0x7e, 0x41, 0x0a, 0x86, 0xee, 0x60, 0x67, 0x42, 0x72, 0x92, 0x69, 0x17, 0x5c, 0x83, 0x9e,
	0x7d, 0x7b, 0xe9, 0x37, 0x45, 0xfa, 0x1f, 0x0c, 0x13, 0xde, 0x5b, 0xac, 0xae, 0xac, 0xa8, 0x52,
	0x20, 0x0a, 0xcf, 0x34, 0x4f, 0x25, 0x97, 0xe9, 0x80, 0xcb, 0x44, 0x69, 0xb7, 0x75, 0xdd, 0xee,
	0xd9, 0xb7, 0x4f, 0x9b, 0x2d, 0xfa, 0x25, 0xfa, 0x4e, 0x26, 0x2a, 0xbc, 0xdc, 0xfa, 0xfc, 0x5e,
	0x5d, 0x39, 0x73, 0x92, 0x89, 0x3b, 0x7c, 0xe0, 0x82, 0xa3, 0xae, 0xde, 0xa3, 0x1a, 0x49, 0x78,
	0x96, 0x71, 0xad, 0x19, 0x1d, 0xc4, 0x42, 0x0d, 0xc7, 0xda, 0x6d, 0x9b, 0x94, 0x97, 0xcd, 0x29,
	0x1f, 0x89, 0xe0, 0x94, 0x14, 0x2a, 0x7f, 0x6f, 0x34, 0xa1, 0x91, 0x1c, 0xe7, 0x1d, 0xf8, 0xe1,
	0xa8, 0x9b, 0xd5, 0x58, 0xfc, 0x03, 0x40, 0xbb, 0xd6, 0x15, 0xb9, 0xf0, 0x01, 0xa1, 0x34, 0x67,
	0xba, 0x9c, 0xe8, 0x51, 0xb4, 0x3b, 0xa2, 0x2f, 0x00, 0x9e, 0x4f, 0x77, 0x79, 0x83, 0xfa, 0x25,
	0xdc, 0x96, 0x19, 0xf3, 0xc5, 0x3f, 0x3a, 0xd6, 0x27, 0x79, 0x5e, 0x55, 0x7c, 0x52, 0x56, 0x6c,
	0xf6, 0xc5, 0x91, 0x33, 0x6d, 0x10, 0xe3, 0x6f, 0x00, 0x3e, 0x6e, 0xbc, 0xf9, 0x89, 0xf6, 0xf4,
	0x78, 0xd7, 0x93, 0x6f, 0xaf, 0x66, 0xfa, 0x5f, 0x6b, 0xbe, 0x81, 0x76, 0x4d, 0x8a, 0x1c, 0x78,
	0x9f, 0x4b, 0xca, 0x66, 0xa6, 0x4c, 0x3b, 0x2a, 0x0f, 0xe8, 0x1c, 0x76, 0x4a, 0x91, 0xd9, 0xed,
	0x61, 0x54, 0x9d, 0xc2, 0x70, 0xb1, 0xf6, 0xc0, 0x72, 0xed, 0x81, 0x5f, 0x6b, 0x0f, 0x7c, 0xdd,
	0x78, 0xd6, 0x72, 0xe3, 0x59, 0x3f, 0x37, 0x9e, 0xf5, 0xa9, 0x97, 0xf2, 0x62, 0xf4, 0x39, 0xf6,
	0x87, 0x2a, 0x0b, 0x04, 0x97, 0x2c, 0x10, 0x49, 0xfc, 0x4a, 0xd3, 0x71, 0x30, 0xdb, 0xff, 0x02,
	0xc5, 0x7c, 0xc2, 0x74, 0xdc, 0x31, 0x1f, 0xfe, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x96,
	0xdc, 0xc7, 0xe2, 0x6f, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MissedBlocks) > 0 {
		for iNdEx := len(m.MissedBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MissedBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SigningInfos) > 0 {
		for iNdEx := len(m.SigningInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SigningInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *SigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ValidatorSigningInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorMissedBlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorMissedBlocks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorMissedBlocks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MissedBlocks) > 0 {
		for iNdEx := len(m.MissedBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MissedBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MissedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MissedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MissedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Missed {
		i--
		if m.Missed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Index != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.SigningInfos) > 0 {
		for _, e := range m.SigningInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MissedBlocks) > 0 {
		for _, e := range m.MissedBlocks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *SigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.ValidatorSigningInfo.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *ValidatorMissedBlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MissedBlocks) > 0 {
		for _, e := range m.MissedBlocks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *MissedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovGenesis(uint64(m.Index))
	}
	if m.Missed {
		n += 2
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SigningInfos = append(m.SigningInfos, SigningInfo{})
			if err := m.SigningInfos[len(m.SigningInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MissedBlocks = append(m.MissedBlocks, ValidatorMissedBlocks{})
			if err := m.MissedBlocks[len(m.MissedBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *SigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: SigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSigningInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ValidatorSigningInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ValidatorMissedBlocks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ValidatorMissedBlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorMissedBlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MissedBlocks = append(m.MissedBlocks, MissedBlock{})
			if err := m.MissedBlocks[len(m.MissedBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *MissedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: MissedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MissedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Missed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Missed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
