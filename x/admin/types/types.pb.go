// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gridnode/admin/v1/types.proto

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

type AdminType int32

const (
	AdminType_CLPDEX        AdminType = 0
	AdminType_PMTPREWARDS   AdminType = 1
	AdminType_TOKENREGISTRY AdminType = 2
	AdminType_ETHBRIDGE     AdminType = 3
	AdminType_ADMIN         AdminType = 4
	AdminType_MARGIN        AdminType = 5
)

var AdminType_name = map[int32]string{
	0: "CLPDEX",
	1: "PMTPREWARDS",
	2: "TOKENREGISTRY",
	3: "ETHBRIDGE",
	4: "ADMIN",
	5: "MARGIN",
}

var AdminType_value = map[string]int32{
	"CLPDEX":        0,
	"PMTPREWARDS":   1,
	"TOKENREGISTRY": 2,
	"ETHBRIDGE":     3,
	"ADMIN":         4,
	"MARGIN":        5,
}

func (x AdminType) String() string {
	return proto.EnumName(AdminType_name, int32(x))
}

func (AdminType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_af48241f5e09954e, []int{0}
}

type GenesisState struct {
	AdminAccounts []*AdminAccount `protobuf:"bytes,1,rep,name=admin_accounts,json=adminAccounts,proto3" json:"admin_accounts,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_af48241f5e09954e, []int{0}
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

func (m *GenesisState) GetAdminAccounts() []*AdminAccount {
	if m != nil {
		return m.AdminAccounts
	}
	return nil
}

type AdminAccount struct {
	AdminType    AdminType `protobuf:"varint,1,opt,name=admin_type,json=adminType,proto3,enum=gridnode.admin.v1.AdminType" json:"admin_type,omitempty"`
	AdminAddress string    `protobuf:"bytes,2,opt,name=admin_address,json=adminAddress,proto3" json:"admin_address,omitempty"`
}

func (m *AdminAccount) Reset()         { *m = AdminAccount{} }
func (m *AdminAccount) String() string { return proto.CompactTextString(m) }
func (*AdminAccount) ProtoMessage()    {}
func (*AdminAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_af48241f5e09954e, []int{1}
}
func (m *AdminAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdminAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdminAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AdminAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminAccount.Merge(m, src)
}
func (m *AdminAccount) XXX_Size() int {
	return m.Size()
}
func (m *AdminAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminAccount.DiscardUnknown(m)
}

var xxx_messageInfo_AdminAccount proto.InternalMessageInfo

func (m *AdminAccount) GetAdminType() AdminType {
	if m != nil {
		return m.AdminType
	}
	return AdminType_CLPDEX
}

func (m *AdminAccount) GetAdminAddress() string {
	if m != nil {
		return m.AdminAddress
	}
	return ""
}

type Params struct {
	SubmitProposalFee github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,1,opt,name=submit_proposal_fee,json=submitProposalFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"submit_proposal_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_af48241f5e09954e, []int{2}
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

func init() {
	proto.RegisterEnum("gridnode.admin.v1.AdminType", AdminType_name, AdminType_value)
	proto.RegisterType((*GenesisState)(nil), "gridnode.admin.v1.GenesisState")
	proto.RegisterType((*AdminAccount)(nil), "gridnode.admin.v1.AdminAccount")
	proto.RegisterType((*Params)(nil), "gridnode.admin.v1.Params")
}

func init() { proto.RegisterFile("gridnode/admin/v1/types.proto", fileDescriptor_af48241f5e09954e) }

var fileDescriptor_af48241f5e09954e = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0x8d, 0x55, 0xca, 0xbb, 0x76, 0xa4, 0x86, 0x43, 0x85, 0x20, 0xad, 0xca, 0x81,
	0x0a, 0x89, 0x58, 0x1b, 0x47, 0x4e, 0x29, 0xcd, 0x42, 0x04, 0x2d, 0x91, 0x1b, 0xfe, 0x5e, 0x2a,
	0x37, 0x31, 0x9d, 0x05, 0x89, 0xa3, 0xd8, 0x9d, 0xd8, 0xb7, 0xe0, 0x63, 0xed, 0xb8, 0x23, 0xe2,
	0x30, 0xa1, 0xf6, 0x8b, 0xa0, 0xfc, 0xd9, 0x54, 0x09, 0xed, 0x94, 0x27, 0x7e, 0x5e, 0xff, 0xde,
	0xc7, 0xf6, 0x0b, 0x4f, 0x56, 0x85, 0x48, 0x32, 0x99, 0x70, 0xc2, 0x92, 0x54, 0x64, 0xe4, 0xfc,
	0x98, 0xe8, 0x8b, 0x9c, 0x2b, 0x27, 0x2f, 0xa4, 0x96, 0xb8, 0x7b, 0x63, 0x3b, 0x95, 0xed, 0x9c,
	0x1f, 0x3f, 0x7a, 0xb8, 0x92, 0x2b, 0x59, 0xb9, 0xa4, 0x54, 0x75, 0xe1, 0xf0, 0x23, 0xb4, 0x7d,
	0x9e, 0x71, 0x25, 0xd4, 0x5c, 0x33, 0xcd, 0xf1, 0x29, 0x1c, 0x55, 0x3b, 0x16, 0x2c, 0x8e, 0xe5,
	0x3a, 0xd3, 0xaa, 0x87, 0x06, 0xfb, 0xa3, 0xc3, 0x93, 0xbe, 0xf3, 0x1f, 0xd1, 0x71, 0x4b, 0xe1,
	0xd6, 0x75, 0xb4, 0xc3, 0x76, 0xfe, 0xd4, 0x30, 0x87, 0xf6, 0xae, 0x8d, 0x5f, 0x01, 0xd4, 0xdc,
	0x32, 0x65, 0x0f, 0x0d, 0xd0, 0xe8, 0xe8, 0xe4, 0xf1, 0x5d, 0xcc, 0xe8, 0x22, 0xe7, 0xd4, 0x64,
	0x37, 0x12, 0x3f, 0x85, 0x4e, 0x13, 0x2a, 0x49, 0x0a, 0xae, 0x54, 0x6f, 0x6f, 0x80, 0x46, 0x26,
	0x6d, 0xd7, 0x2d, 0xeb, 0xb5, 0xa1, 0x80, 0x56, 0xc8, 0x0a, 0x96, 0x2a, 0xbc, 0x80, 0x07, 0x6a,
	0xbd, 0x4c, 0x85, 0x5e, 0xe4, 0x85, 0xcc, 0xa5, 0x62, 0x3f, 0x16, 0xdf, 0x78, 0xdd, 0xd4, 0x1c,
	0x93, 0xcb, 0xeb, 0xbe, 0xf1, 0xe7, 0xba, 0xff, 0x6c, 0x25, 0xf4, 0xd9, 0x7a, 0xe9, 0xc4, 0x32,
	0x25, 0xb1, 0x54, 0xa9, 0x54, 0xcd, 0xe7, 0x85, 0x4a, 0xbe, 0x37, 0x77, 0xf9, 0x41, 0x64, 0x9a,
	0x76, 0x6b, 0x56, 0xd8, 0xa0, 0x4e, 0x39, 0x7f, 0xce, 0xc0, 0xbc, 0xcd, 0x89, 0x01, 0x5a, 0xaf,
	0xdf, 0x85, 0x13, 0xef, 0xb3, 0x65, 0xe0, 0xfb, 0x70, 0x18, 0x4e, 0xa3, 0x90, 0x7a, 0x9f, 0x5c,
	0x3a, 0x99, 0x5b, 0x08, 0x77, 0xa1, 0x13, 0xbd, 0x7f, 0xeb, 0xcd, 0xa8, 0xe7, 0x07, 0xf3, 0x88,
	0x7e, 0xb1, 0xf6, 0x70, 0x07, 0x4c, 0x2f, 0x7a, 0x33, 0xa6, 0xc1, 0xc4, 0xf7, 0xac, 0x7d, 0x6c,
	0xc2, 0x81, 0x3b, 0x99, 0x06, 0x33, 0xeb, 0x5e, 0x49, 0x9a, 0xba, 0xd4, 0x0f, 0x66, 0xd6, 0xc1,
	0x38, 0xb8, 0xdc, 0xd8, 0xe8, 0x6a, 0x63, 0xa3, 0xbf, 0x1b, 0x1b, 0xfd, 0xda, 0xda, 0xc6, 0xd5,
	0xd6, 0x36, 0x7e, 0x6f, 0x6d, 0xe3, 0x2b, 0xd9, 0x09, 0xee, 0x17, 0x22, 0x11, 0x85, 0xcc, 0xe2,
	0x33, 0x26, 0x32, 0x72, 0x3b, 0x12, 0x3f, 0x9b, 0xa1, 0xa8, 0x4e, 0xb1, 0x6c, 0x55, 0x2f, 0xfd,
	0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x61, 0x29, 0xc8, 0x33, 0x02, 0x00, 0x00,
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
	if len(m.AdminAccounts) > 0 {
		for iNdEx := len(m.AdminAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AdminAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AdminAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdminAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AdminAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdminAddress) > 0 {
		i -= len(m.AdminAddress)
		copy(dAtA[i:], m.AdminAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AdminAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.AdminType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.AdminType))
		i--
		dAtA[i] = 0x8
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
		size := m.SubmitProposalFee.Size()
		i -= size
		if _, err := m.SubmitProposalFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AdminAccounts) > 0 {
		for _, e := range m.AdminAccounts {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *AdminAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AdminType != 0 {
		n += 1 + sovTypes(uint64(m.AdminType))
	}
	l = len(m.AdminAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SubmitProposalFee.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminAccounts = append(m.AdminAccounts, &AdminAccount{})
			if err := m.AdminAccounts[len(m.AdminAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *AdminAccount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AdminAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdminAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminType", wireType)
			}
			m.AdminType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdminType |= AdminType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAddress", wireType)
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
			m.AdminAddress = string(dAtA[iNdEx:postIndex])
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitProposalFee", wireType)
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
			if err := m.SubmitProposalFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
