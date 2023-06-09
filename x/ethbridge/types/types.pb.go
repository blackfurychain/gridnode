// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gridnode/ethbridge/v1/types.proto

package types

import (
	fmt "fmt"
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

// Claim type enum
type ClaimType int32

const (
	// Unspecified claim type
	ClaimType_CLAIM_TYPE_UNSPECIFIED ClaimType = 0
	// Burn claim type
	ClaimType_CLAIM_TYPE_BURN ClaimType = 1
	// Lock claim type
	ClaimType_CLAIM_TYPE_LOCK ClaimType = 2
)

var ClaimType_name = map[int32]string{
	0: "CLAIM_TYPE_UNSPECIFIED",
	1: "CLAIM_TYPE_BURN",
	2: "CLAIM_TYPE_LOCK",
}

var ClaimType_value = map[string]int32{
	"CLAIM_TYPE_UNSPECIFIED": 0,
	"CLAIM_TYPE_BURN":        1,
	"CLAIM_TYPE_LOCK":        2,
}

func (x ClaimType) String() string {
	return proto.EnumName(ClaimType_name, int32(x))
}

func (ClaimType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a77e3806a8dc9275, []int{0}
}

// EthBridgeClaim is a structure that contains all the data for a particular
// bridge claim
type EthBridgeClaim struct {
	EthereumChainId int64 `protobuf:"varint,1,opt,name=ethereum_chain_id,json=ethereumChainId,proto3" json:"ethereum_chain_id,omitempty" yaml:"ethereum_chain_id"`
	// bridge_contract_address is an EthereumAddress
	BridgeContractAddress string `protobuf:"bytes,2,opt,name=bridge_contract_address,json=bridgeContractAddress,proto3" json:"bridge_contract_address,omitempty" yaml:"bridge_contract_address"`
	Nonce                 int64  `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty" yaml:"nonce"`
	Symbol                string `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty" yaml:"symbol"`
	// token_contract_address is an EthereumAddress
	TokenContractAddress string `protobuf:"bytes,5,opt,name=token_contract_address,json=tokenContractAddress,proto3" json:"token_contract_address,omitempty" yaml:"token_contract_address"`
	// ethereum_sender is an EthereumAddress
	EthereumSender string `protobuf:"bytes,6,opt,name=ethereum_sender,json=ethereumSender,proto3" json:"ethereum_sender,omitempty" yaml:"ethereum_sender"`
	// cosmos_receiver is an sdk.AccAddress
	CosmosReceiver string `protobuf:"bytes,7,opt,name=cosmos_receiver,json=cosmosReceiver,proto3" json:"cosmos_receiver,omitempty" yaml:"cosmos_receiver"`
	// validator_address is an sdk.ValAddress
	ValidatorAddress string                                 `protobuf:"bytes,8,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
	Amount           github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount" yaml:"amount"`
	ClaimType        ClaimType                              `protobuf:"varint,10,opt,name=claim_type,json=claimType,proto3,enum=gridnode.ethbridge.v1.ClaimType" json:"claim_type,omitempty"`
}

func (m *EthBridgeClaim) Reset()         { *m = EthBridgeClaim{} }
func (m *EthBridgeClaim) String() string { return proto.CompactTextString(m) }
func (*EthBridgeClaim) ProtoMessage()    {}
func (*EthBridgeClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_a77e3806a8dc9275, []int{0}
}
func (m *EthBridgeClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthBridgeClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthBridgeClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthBridgeClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthBridgeClaim.Merge(m, src)
}
func (m *EthBridgeClaim) XXX_Size() int {
	return m.Size()
}
func (m *EthBridgeClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_EthBridgeClaim.DiscardUnknown(m)
}

var xxx_messageInfo_EthBridgeClaim proto.InternalMessageInfo

func (m *EthBridgeClaim) GetEthereumChainId() int64 {
	if m != nil {
		return m.EthereumChainId
	}
	return 0
}

func (m *EthBridgeClaim) GetBridgeContractAddress() string {
	if m != nil {
		return m.BridgeContractAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *EthBridgeClaim) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *EthBridgeClaim) GetTokenContractAddress() string {
	if m != nil {
		return m.TokenContractAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetEthereumSender() string {
	if m != nil {
		return m.EthereumSender
	}
	return ""
}

func (m *EthBridgeClaim) GetCosmosReceiver() string {
	if m != nil {
		return m.CosmosReceiver
	}
	return ""
}

func (m *EthBridgeClaim) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *EthBridgeClaim) GetClaimType() ClaimType {
	if m != nil {
		return m.ClaimType
	}
	return ClaimType_CLAIM_TYPE_UNSPECIFIED
}

type PeggyTokens struct {
	Tokens []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (m *PeggyTokens) Reset()         { *m = PeggyTokens{} }
func (m *PeggyTokens) String() string { return proto.CompactTextString(m) }
func (*PeggyTokens) ProtoMessage()    {}
func (*PeggyTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_a77e3806a8dc9275, []int{1}
}
func (m *PeggyTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeggyTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeggyTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeggyTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeggyTokens.Merge(m, src)
}
func (m *PeggyTokens) XXX_Size() int {
	return m.Size()
}
func (m *PeggyTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_PeggyTokens.DiscardUnknown(m)
}

var xxx_messageInfo_PeggyTokens proto.InternalMessageInfo

func (m *PeggyTokens) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

// GenesisState for ethbridge
type GenesisState struct {
	CethReceiveAccount string   `protobuf:"bytes,1,opt,name=ceth_receive_account,json=cethReceiveAccount,proto3" json:"ceth_receive_account,omitempty"`
	PeggyTokens        []string `protobuf:"bytes,2,rep,name=peggy_tokens,json=peggyTokens,proto3" json:"peggy_tokens,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a77e3806a8dc9275, []int{2}
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

func (m *GenesisState) GetCethReceiveAccount() string {
	if m != nil {
		return m.CethReceiveAccount
	}
	return ""
}

func (m *GenesisState) GetPeggyTokens() []string {
	if m != nil {
		return m.PeggyTokens
	}
	return nil
}

type Pause struct {
	IsPaused bool `protobuf:"varint,1,opt,name=is_paused,json=isPaused,proto3" json:"is_paused,omitempty"`
}

func (m *Pause) Reset()         { *m = Pause{} }
func (m *Pause) String() string { return proto.CompactTextString(m) }
func (*Pause) ProtoMessage()    {}
func (*Pause) Descriptor() ([]byte, []int) {
	return fileDescriptor_a77e3806a8dc9275, []int{3}
}
func (m *Pause) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pause) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pause.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pause) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pause.Merge(m, src)
}
func (m *Pause) XXX_Size() int {
	return m.Size()
}
func (m *Pause) XXX_DiscardUnknown() {
	xxx_messageInfo_Pause.DiscardUnknown(m)
}

var xxx_messageInfo_Pause proto.InternalMessageInfo

func (m *Pause) GetIsPaused() bool {
	if m != nil {
		return m.IsPaused
	}
	return false
}

func init() {
	proto.RegisterEnum("gridnode.ethbridge.v1.ClaimType", ClaimType_name, ClaimType_value)
	proto.RegisterType((*EthBridgeClaim)(nil), "gridnode.ethbridge.v1.EthBridgeClaim")
	proto.RegisterType((*PeggyTokens)(nil), "gridnode.ethbridge.v1.PeggyTokens")
	proto.RegisterType((*GenesisState)(nil), "gridnode.ethbridge.v1.GenesisState")
	proto.RegisterType((*Pause)(nil), "gridnode.ethbridge.v1.Pause")
}

func init() { proto.RegisterFile("gridnode/ethbridge/v1/types.proto", fileDescriptor_a77e3806a8dc9275) }

var fileDescriptor_a77e3806a8dc9275 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdf, 0x4e, 0xdb, 0x3c,
	0x1c, 0x6d, 0xe0, 0x6b, 0x3f, 0x62, 0x58, 0x29, 0x1e, 0x74, 0x11, 0xdb, 0x92, 0x62, 0x6d, 0xa8,
	0x9b, 0xb4, 0x76, 0x8c, 0xbb, 0xdd, 0x20, 0x9a, 0x75, 0xac, 0x1a, 0xb0, 0xce, 0x80, 0xd0, 0xb8,
	0x89, 0xd2, 0xc4, 0x4a, 0x2d, 0x9a, 0xb8, 0x8a, 0xdd, 0x6a, 0x7d, 0x8b, 0x3d, 0x16, 0x97, 0xec,
	0x6e, 0xda, 0x45, 0x34, 0xc1, 0x1b, 0xf4, 0x09, 0xa6, 0xd8, 0x69, 0xa9, 0x0a, 0xbb, 0x8a, 0x73,
	0xce, 0xf1, 0xf9, 0xf9, 0xf7, 0xc7, 0x06, 0x5b, 0x41, 0x4c, 0xfd, 0x88, 0xf9, 0xa4, 0x4e, 0x44,
	0xb7, 0x13, 0x53, 0x3f, 0x20, 0xf5, 0xe1, 0x4e, 0x5d, 0x8c, 0xfa, 0x84, 0xd7, 0xfa, 0x31, 0x13,
	0x0c, 0x6e, 0x4c, 0x24, 0xb5, 0xa9, 0xa4, 0x36, 0xdc, 0xd9, 0x5c, 0x0f, 0x58, 0xc0, 0xa4, 0xa2,
	0x9e, 0xae, 0x94, 0x18, 0xfd, 0xcc, 0x83, 0x62, 0x53, 0x74, 0x1b, 0x52, 0x66, 0xf7, 0x5c, 0x1a,
	0xc2, 0x4f, 0x60, 0x8d, 0x88, 0x2e, 0x89, 0xc9, 0x20, 0x74, 0xbc, 0xae, 0x4b, 0x23, 0x87, 0xfa,
	0x86, 0x56, 0xd1, 0xaa, 0x8b, 0x8d, 0x67, 0xe3, 0xc4, 0x32, 0x46, 0x6e, 0xd8, 0x7b, 0x8f, 0xee,
	0x49, 0x10, 0x5e, 0x9d, 0x60, 0x76, 0x0a, 0xb5, 0x7c, 0x78, 0x01, 0x9e, 0xa8, 0xf8, 0x8e, 0xc7,
	0x22, 0x11, 0xbb, 0x9e, 0x70, 0x5c, 0xdf, 0x8f, 0x09, 0xe7, 0xc6, 0x42, 0x45, 0xab, 0xea, 0x0d,
	0x34, 0x4e, 0x2c, 0x53, 0xf9, 0xfd, 0x43, 0x88, 0xf0, 0x86, 0x62, 0xec, 0x8c, 0xd8, 0x57, 0x38,
	0xdc, 0x06, 0xf9, 0x88, 0x45, 0x1e, 0x31, 0x16, 0xe5, 0xc9, 0x4a, 0xe3, 0xc4, 0x5a, 0x51, 0x4e,
	0x12, 0x46, 0x58, 0xd1, 0xf0, 0x15, 0x28, 0xf0, 0x51, 0xd8, 0x61, 0x3d, 0xe3, 0x3f, 0x19, 0x72,
	0x6d, 0x9c, 0x58, 0x8f, 0x94, 0x50, 0xe1, 0x08, 0x67, 0x02, 0x78, 0x0e, 0xca, 0x82, 0x5d, 0x92,
	0xe8, 0xfe, 0x69, 0xf3, 0x72, 0xeb, 0xd6, 0x38, 0xb1, 0x9e, 0xab, 0xad, 0x0f, 0xeb, 0x10, 0x5e,
	0x97, 0xc4, 0xfc, 0x59, 0x6d, 0x30, 0x2d, 0x8d, 0xc3, 0x49, 0xe4, 0x93, 0xd8, 0x28, 0x48, 0xc7,
	0xcd, 0x71, 0x62, 0x95, 0xe7, 0xea, 0xa9, 0x04, 0x08, 0x17, 0x27, 0xc8, 0x89, 0x04, 0x52, 0x13,
	0x8f, 0xf1, 0x90, 0x71, 0x27, 0x26, 0x1e, 0xa1, 0x43, 0x12, 0x1b, 0xff, 0xcf, 0x9b, 0xcc, 0x09,
	0x10, 0x2e, 0x2a, 0x04, 0x67, 0x00, 0x6c, 0x81, 0xb5, 0xa1, 0xdb, 0xa3, 0xbe, 0x2b, 0x58, 0x3c,
	0xcd, 0x6e, 0x49, 0xda, 0xcc, 0xf4, 0xf6, 0x9e, 0x04, 0xe1, 0xd2, 0x14, 0x9b, 0x24, 0x75, 0x0e,
	0x0a, 0x6e, 0xc8, 0x06, 0x91, 0x30, 0x74, 0xb9, 0x7f, 0xef, 0x2a, 0xb1, 0x72, 0xbf, 0x13, 0x6b,
	0x3b, 0xa0, 0xa2, 0x3b, 0xe8, 0xd4, 0x3c, 0x16, 0xd6, 0x55, 0xf4, 0xec, 0xf3, 0x86, 0xfb, 0x97,
	0xd9, 0xa0, 0xb6, 0x22, 0x71, 0xd7, 0x06, 0xe5, 0x82, 0x70, 0x66, 0x07, 0xf7, 0x00, 0xf0, 0xd2,
	0x41, 0x74, 0x52, 0xad, 0x01, 0x2a, 0x5a, 0xb5, 0xf8, 0xae, 0x52, 0x7b, 0x70, 0xa8, 0x6b, 0x72,
	0x62, 0x4f, 0x47, 0x7d, 0x82, 0x75, 0x6f, 0xb2, 0x44, 0x2f, 0xc1, 0x72, 0x9b, 0x04, 0xc1, 0xe8,
	0x34, 0xed, 0x05, 0x87, 0x65, 0x50, 0x90, 0x5d, 0xe1, 0x86, 0x56, 0x59, 0xac, 0xea, 0x38, 0xfb,
	0x43, 0x1e, 0x58, 0x39, 0x20, 0x11, 0xe1, 0x94, 0x9f, 0x08, 0x57, 0x10, 0xf8, 0x16, 0xac, 0x7b,
	0x44, 0x74, 0x27, 0xd5, 0x73, 0x5c, 0xcf, 0x93, 0xe9, 0xa5, 0xa3, 0xaf, 0x63, 0x98, 0x72, 0x59,
	0x1d, 0xf7, 0x15, 0x03, 0xb7, 0xc0, 0x4a, 0x3f, 0x0d, 0xe4, 0x64, 0xfe, 0x0b, 0xd2, 0x7f, 0xb9,
	0x7f, 0x17, 0x1c, 0xbd, 0x00, 0xf9, 0xb6, 0x3b, 0xe0, 0x04, 0x3e, 0x05, 0x3a, 0xe5, 0x4e, 0x3f,
	0x5d, 0xab, 0xdb, 0xb4, 0x84, 0x97, 0x28, 0x97, 0x9c, 0xff, 0xfa, 0x2b, 0xd0, 0xa7, 0x99, 0xc0,
	0x4d, 0x50, 0xb6, 0x0f, 0xf7, 0x5b, 0x47, 0xce, 0xe9, 0xb7, 0x76, 0xd3, 0x39, 0x3b, 0x3e, 0x69,
	0x37, 0xed, 0xd6, 0xc7, 0x56, 0xf3, 0x43, 0x29, 0x07, 0x1f, 0x83, 0xd5, 0x19, 0xae, 0x71, 0x86,
	0x8f, 0x4b, 0xda, 0x1c, 0x78, 0xf8, 0xc5, 0xfe, 0x5c, 0x5a, 0x68, 0x1c, 0x5d, 0xdd, 0x98, 0xda,
	0xf5, 0x8d, 0xa9, 0xfd, 0xb9, 0x31, 0xb5, 0x1f, 0xb7, 0x66, 0xee, 0xfa, 0xd6, 0xcc, 0xfd, 0xba,
	0x35, 0x73, 0x17, 0xbb, 0x33, 0x0d, 0x3a, 0x88, 0xa9, 0x4f, 0x63, 0x16, 0xc9, 0x4b, 0x5c, 0x9f,
	0xbe, 0x2d, 0xdf, 0x67, 0x5e, 0x17, 0xd9, 0xb1, 0x4e, 0x41, 0x3e, 0x17, 0xbb, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x93, 0x2c, 0x06, 0xbd, 0x80, 0x04, 0x00, 0x00,
}

func (m *EthBridgeClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthBridgeClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthBridgeClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClaimType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ClaimType))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.CosmosReceiver) > 0 {
		i -= len(m.CosmosReceiver)
		copy(dAtA[i:], m.CosmosReceiver)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CosmosReceiver)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.EthereumSender) > 0 {
		i -= len(m.EthereumSender)
		copy(dAtA[i:], m.EthereumSender)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EthereumSender)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TokenContractAddress) > 0 {
		i -= len(m.TokenContractAddress)
		copy(dAtA[i:], m.TokenContractAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TokenContractAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x22
	}
	if m.Nonce != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BridgeContractAddress) > 0 {
		i -= len(m.BridgeContractAddress)
		copy(dAtA[i:], m.BridgeContractAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BridgeContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.EthereumChainId != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.EthereumChainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PeggyTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeggyTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeggyTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tokens[iNdEx])
			copy(dAtA[i:], m.Tokens[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Tokens[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
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
	if len(m.PeggyTokens) > 0 {
		for iNdEx := len(m.PeggyTokens) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PeggyTokens[iNdEx])
			copy(dAtA[i:], m.PeggyTokens[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.PeggyTokens[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CethReceiveAccount) > 0 {
		i -= len(m.CethReceiveAccount)
		copy(dAtA[i:], m.CethReceiveAccount)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CethReceiveAccount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pause) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pause) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pause) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsPaused {
		i--
		if m.IsPaused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthBridgeClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EthereumChainId != 0 {
		n += 1 + sovTypes(uint64(m.EthereumChainId))
	}
	l = len(m.BridgeContractAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovTypes(uint64(m.Nonce))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.TokenContractAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EthereumSender)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.CosmosReceiver)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.ClaimType != 0 {
		n += 1 + sovTypes(uint64(m.ClaimType))
	}
	return n
}

func (m *PeggyTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for _, s := range m.Tokens {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CethReceiveAccount)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.PeggyTokens) > 0 {
		for _, s := range m.PeggyTokens {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Pause) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsPaused {
		n += 2
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthBridgeClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: EthBridgeClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthBridgeClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumChainId", wireType)
			}
			m.EthereumChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthereumChainId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthereumSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosReceiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CosmosReceiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			m.ClaimType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimType |= ClaimType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *PeggyTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PeggyTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeggyTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
				return fmt.Errorf("proto: wrong wireType = %d for field CethReceiveAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CethReceiveAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeggyTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeggyTokens = append(m.PeggyTokens, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Pause) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Pause: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pause: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPaused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.IsPaused = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
