// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchainqueries/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type RegisteredQuery struct {
	// The unique id of the registered query.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The address that registered the query.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// The query type identifier: `kv` or `tx` now
	QueryType string `protobuf:"bytes,3,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	// The KV-storage keys for which we want to get values from remote chain
	Keys []*KVKey `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty"`
	// The filter for transaction search ICQ
	TransactionsFilter string `protobuf:"bytes,5,opt,name=transactions_filter,json=transactionsFilter,proto3" json:"transactions_filter,omitempty"`
	// The IBC connection ID for getting ConsensusState to verify proofs
	ConnectionId string `protobuf:"bytes,6,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// Parameter that defines how often the query must be updated.
	UpdatePeriod uint64 `protobuf:"varint,7,opt,name=update_period,json=updatePeriod,proto3" json:"update_period,omitempty"`
	// The local chain last block height when the query result was updated.
	LastSubmittedResultLocalHeight uint64 `protobuf:"varint,8,opt,name=last_submitted_result_local_height,json=lastSubmittedResultLocalHeight,proto3" json:"last_submitted_result_local_height,omitempty"`
	// The remote chain last block height when the query result was updated.
	LastSubmittedResultRemoteHeight uint64 `protobuf:"varint,9,opt,name=last_submitted_result_remote_height,json=lastSubmittedResultRemoteHeight,proto3" json:"last_submitted_result_remote_height,omitempty"`
	// Amount of coins deposited for the query.
	Deposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,10,rep,name=deposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"deposit"`
	// Timeout before query becomes available for everybody to remove.
	SubmitTimeout uint64 `protobuf:"varint,12,opt,name=submit_timeout,json=submitTimeout,proto3" json:"submit_timeout,omitempty"`
}

func (m *RegisteredQuery) Reset()         { *m = RegisteredQuery{} }
func (m *RegisteredQuery) String() string { return proto.CompactTextString(m) }
func (*RegisteredQuery) ProtoMessage()    {}
func (*RegisteredQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e6c14f58b92f58, []int{0}
}
func (m *RegisteredQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteredQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteredQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteredQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredQuery.Merge(m, src)
}
func (m *RegisteredQuery) XXX_Size() int {
	return m.Size()
}
func (m *RegisteredQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredQuery proto.InternalMessageInfo

func (m *RegisteredQuery) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegisteredQuery) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RegisteredQuery) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *RegisteredQuery) GetKeys() []*KVKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *RegisteredQuery) GetTransactionsFilter() string {
	if m != nil {
		return m.TransactionsFilter
	}
	return ""
}

func (m *RegisteredQuery) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *RegisteredQuery) GetUpdatePeriod() uint64 {
	if m != nil {
		return m.UpdatePeriod
	}
	return 0
}

func (m *RegisteredQuery) GetLastSubmittedResultLocalHeight() uint64 {
	if m != nil {
		return m.LastSubmittedResultLocalHeight
	}
	return 0
}

func (m *RegisteredQuery) GetLastSubmittedResultRemoteHeight() uint64 {
	if m != nil {
		return m.LastSubmittedResultRemoteHeight
	}
	return 0
}

func (m *RegisteredQuery) GetDeposit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Deposit
	}
	return nil
}

func (m *RegisteredQuery) GetSubmitTimeout() uint64 {
	if m != nil {
		return m.SubmitTimeout
	}
	return 0
}

type KVKey struct {
	// Path (storage prefix) to the storage where you want to read value by key
	// (usually name of cosmos-sdk module: 'staking', 'bank', etc.)
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Key you want to read from the storage
	Key []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *KVKey) Reset()         { *m = KVKey{} }
func (m *KVKey) String() string { return proto.CompactTextString(m) }
func (*KVKey) ProtoMessage()    {}
func (*KVKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e6c14f58b92f58, []int{1}
}
func (m *KVKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KVKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KVKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KVKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVKey.Merge(m, src)
}
func (m *KVKey) XXX_Size() int {
	return m.Size()
}
func (m *KVKey) XXX_DiscardUnknown() {
	xxx_messageInfo_KVKey.DiscardUnknown(m)
}

var xxx_messageInfo_KVKey proto.InternalMessageInfo

func (m *KVKey) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *KVKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

// GenesisState defines the interchainadapter module's genesis state.
type GenesisState struct {
	Params            Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	RegisteredQueries []*RegisteredQuery `protobuf:"bytes,2,rep,name=registered_queries,json=registeredQueries,proto3" json:"registered_queries,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e6c14f58b92f58, []int{2}
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

func (m *GenesisState) GetRegisteredQueries() []*RegisteredQuery {
	if m != nil {
		return m.RegisteredQueries
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisteredQuery)(nil), "neutron.interchainadapter.interchainqueries.RegisteredQuery")
	proto.RegisterType((*KVKey)(nil), "neutron.interchainadapter.interchainqueries.KVKey")
	proto.RegisterType((*GenesisState)(nil), "neutron.interchainadapter.interchainqueries.GenesisState")
}

func init() { proto.RegisterFile("interchainqueries/genesis.proto", fileDescriptor_68e6c14f58b92f58) }

var fileDescriptor_68e6c14f58b92f58 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x4e, 0x14, 0x41,
	0x10, 0xde, 0x59, 0x16, 0x70, 0x9b, 0x05, 0xb5, 0xe5, 0x30, 0x92, 0x38, 0x4b, 0x96, 0x98, 0x6c,
	0x62, 0x98, 0x11, 0xb8, 0x78, 0xf0, 0x84, 0x09, 0xfe, 0xc0, 0x01, 0x1a, 0xe2, 0xc1, 0xcb, 0xa4,
	0x77, 0xa6, 0x9c, 0xed, 0xec, 0x6e, 0xf7, 0xd8, 0x5d, 0xa3, 0xce, 0x5b, 0xf8, 0x1c, 0xbe, 0x85,
	0x37, 0x4e, 0x86, 0xa3, 0x27, 0x35, 0xf0, 0x22, 0x66, 0x6a, 0x76, 0x05, 0x85, 0x0b, 0xa7, 0xa9,
	0xf9, 0xea, 0xfb, 0xbe, 0xea, 0xee, 0xaa, 0x62, 0x5d, 0xa5, 0x11, 0x6c, 0x32, 0x94, 0x4a, 0x7f,
	0x28, 0xc0, 0x2a, 0x70, 0x51, 0x06, 0x1a, 0x9c, 0x72, 0x61, 0x6e, 0x0d, 0x1a, 0xfe, 0x44, 0x43,
	0x81, 0xd6, 0xe8, 0xf0, 0x92, 0x28, 0x53, 0x99, 0x23, 0xd8, 0xf0, 0x9a, 0x74, 0x6d, 0x35, 0x33,
	0x99, 0x21, 0x5d, 0x54, 0x45, 0xb5, 0xc5, 0x5a, 0x70, 0xbd, 0x46, 0x2e, 0xad, 0x9c, 0xb8, 0x59,
	0x3e, 0x31, 0x6e, 0x62, 0x5c, 0x34, 0x90, 0x0e, 0xa2, 0x8f, 0x5b, 0x03, 0x40, 0xb9, 0x15, 0x25,
	0x46, 0xe9, 0x3a, 0xdf, 0xfb, 0xd6, 0x62, 0x77, 0x05, 0x64, 0xca, 0x21, 0x58, 0x48, 0x8f, 0x0a,
	0xb0, 0x25, 0x5f, 0x61, 0x4d, 0x95, 0xfa, 0xde, 0xba, 0xd7, 0x6f, 0x89, 0xa6, 0x4a, 0xf9, 0x2a,
	0x9b, 0x37, 0x9f, 0x34, 0x58, 0xbf, 0xb9, 0xee, 0xf5, 0xdb, 0xa2, 0xfe, 0xe1, 0x8f, 0x18, 0xab,
	0x2a, 0x96, 0x31, 0x96, 0x39, 0xf8, 0x73, 0x94, 0x6a, 0x13, 0x72, 0x52, 0xe6, 0xc0, 0xf7, 0x58,
	0x6b, 0x04, 0xa5, 0xf3, 0x5b, 0xeb, 0x73, 0xfd, 0xa5, 0xed, 0xed, 0xf0, 0x16, 0x57, 0x0d, 0xf7,
	0xdf, 0xee, 0x43, 0x29, 0x48, 0xcf, 0x23, 0xf6, 0x00, 0xad, 0xd4, 0x4e, 0x26, 0xa8, 0x8c, 0x76,
	0xf1, 0x7b, 0x35, 0x46, 0xb0, 0xfe, 0x3c, 0xd5, 0xe3, 0x57, 0x53, 0x7b, 0x94, 0xe1, 0x1b, 0x6c,
	0x39, 0x31, 0x5a, 0x03, 0x81, 0xb1, 0x4a, 0xfd, 0x05, 0xa2, 0x76, 0x2e, 0xc1, 0xd7, 0x69, 0x45,
	0x2a, 0xf2, 0x54, 0x22, 0xc4, 0x39, 0x58, 0x65, 0x52, 0x7f, 0x91, 0x6e, 0xdb, 0xa9, 0xc1, 0x43,
	0xc2, 0xf8, 0x1b, 0xd6, 0x1b, 0x4b, 0x87, 0xb1, 0x2b, 0x06, 0x13, 0x85, 0x08, 0x69, 0x6c, 0xc1,
	0x15, 0x63, 0x8c, 0xc7, 0x26, 0x91, 0xe3, 0x78, 0x08, 0x2a, 0x1b, 0xa2, 0x7f, 0x87, 0x94, 0x41,
	0xc5, 0x3c, 0x9e, 0x11, 0x05, 0xf1, 0x0e, 0x2a, 0xda, 0x2b, 0x62, 0xf1, 0x03, 0xb6, 0x71, 0xb3,
	0x97, 0x85, 0x89, 0x41, 0x98, 0x99, 0xb5, 0xc9, 0xac, 0x7b, 0x83, 0x99, 0x20, 0xde, 0xd4, 0x0d,
	0xd8, 0x62, 0x0a, 0xb9, 0x71, 0x0a, 0x7d, 0x46, 0xef, 0xfb, 0x30, 0xac, 0xfb, 0x1c, 0x56, 0x7d,
	0x0e, 0xa7, 0x7d, 0x0e, 0x5f, 0x18, 0xa5, 0x77, 0x9f, 0x9e, 0xfe, 0xec, 0x36, 0xbe, 0xfe, 0xea,
	0xf6, 0x33, 0x85, 0xc3, 0x62, 0x10, 0x26, 0x66, 0x12, 0x4d, 0x87, 0xa2, 0xfe, 0x6c, 0xba, 0x74,
	0x14, 0x55, 0x4d, 0x74, 0x24, 0x70, 0x62, 0xe6, 0xcd, 0x1f, 0xb3, 0x95, 0xfa, 0xbc, 0x31, 0xaa,
	0x09, 0x98, 0x02, 0xfd, 0x0e, 0x9d, 0x6f, 0xb9, 0x46, 0x4f, 0x6a, 0xb0, 0xb7, 0xc9, 0xe6, 0xa9,
	0x63, 0x9c, 0xb3, 0x56, 0x2e, 0x71, 0x48, 0xa3, 0xd3, 0x16, 0x14, 0xf3, 0x7b, 0x6c, 0x6e, 0x04,
	0x25, 0x8d, 0x4e, 0x47, 0x54, 0x61, 0xef, 0xbb, 0xc7, 0x3a, 0x2f, 0xeb, 0x3d, 0x38, 0x46, 0x89,
	0xc0, 0x8f, 0xd8, 0x42, 0x3d, 0xb3, 0x24, 0x5c, 0xda, 0xde, 0xb9, 0xd5, 0xb0, 0x1c, 0x92, 0x74,
	0xb7, 0x55, 0x5d, 0x53, 0x4c, 0x8d, 0xf8, 0x88, 0x71, 0xfb, 0x77, 0xaa, 0xe3, 0x29, 0xd5, 0x6f,
	0xd2, 0x5b, 0x3d, 0xbf, 0x95, 0xfd, 0x7f, 0xcb, 0x21, 0xee, 0xdb, 0x7f, 0x00, 0x05, 0x6e, 0x57,
	0x9c, 0x9e, 0x07, 0xde, 0xd9, 0x79, 0xe0, 0xfd, 0x3e, 0x0f, 0xbc, 0x2f, 0x17, 0x41, 0xe3, 0xec,
	0x22, 0x68, 0xfc, 0xb8, 0x08, 0x1a, 0xef, 0x9e, 0x5d, 0x79, 0xf3, 0x69, 0xd1, 0x4d, 0x63, 0xb3,
	0x59, 0x1c, 0x7d, 0x8e, 0xae, 0xaf, 0x2f, 0x75, 0x62, 0xb0, 0x40, 0xeb, 0xb9, 0xf3, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x34, 0xee, 0xf2, 0xd7, 0x44, 0x04, 0x00, 0x00,
}

func (m *RegisteredQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteredQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisteredQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SubmitTimeout != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SubmitTimeout))
		i--
		dAtA[i] = 0x60
	}
	if len(m.Deposit) > 0 {
		for iNdEx := len(m.Deposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.LastSubmittedResultRemoteHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastSubmittedResultRemoteHeight))
		i--
		dAtA[i] = 0x48
	}
	if m.LastSubmittedResultLocalHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastSubmittedResultLocalHeight))
		i--
		dAtA[i] = 0x40
	}
	if m.UpdatePeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.UpdatePeriod))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TransactionsFilter) > 0 {
		i -= len(m.TransactionsFilter)
		copy(dAtA[i:], m.TransactionsFilter)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TransactionsFilter)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Keys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *KVKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KVKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KVKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
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
	if len(m.RegisteredQueries) > 0 {
		for iNdEx := len(m.RegisteredQueries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegisteredQueries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *RegisteredQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Keys) > 0 {
		for _, e := range m.Keys {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.TransactionsFilter)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.UpdatePeriod != 0 {
		n += 1 + sovGenesis(uint64(m.UpdatePeriod))
	}
	if m.LastSubmittedResultLocalHeight != 0 {
		n += 1 + sovGenesis(uint64(m.LastSubmittedResultLocalHeight))
	}
	if m.LastSubmittedResultRemoteHeight != 0 {
		n += 1 + sovGenesis(uint64(m.LastSubmittedResultRemoteHeight))
	}
	if len(m.Deposit) > 0 {
		for _, e := range m.Deposit {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.SubmitTimeout != 0 {
		n += 1 + sovGenesis(uint64(m.SubmitTimeout))
	}
	return n
}

func (m *KVKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.RegisteredQueries) > 0 {
		for _, e := range m.RegisteredQueries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisteredQuery) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RegisteredQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteredQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
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
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
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
			m.Keys = append(m.Keys, &KVKey{})
			if err := m.Keys[len(m.Keys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionsFilter", wireType)
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
			m.TransactionsFilter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
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
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatePeriod", wireType)
			}
			m.UpdatePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSubmittedResultLocalHeight", wireType)
			}
			m.LastSubmittedResultLocalHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSubmittedResultLocalHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSubmittedResultRemoteHeight", wireType)
			}
			m.LastSubmittedResultRemoteHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSubmittedResultRemoteHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
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
			m.Deposit = append(m.Deposit, types.Coin{})
			if err := m.Deposit[len(m.Deposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitTimeout", wireType)
			}
			m.SubmitTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubmitTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *KVKey) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: KVKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KVKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
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
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
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
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredQueries", wireType)
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
			m.RegisteredQueries = append(m.RegisteredQueries, &RegisteredQuery{})
			if err := m.RegisteredQueries[len(m.RegisteredQueries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
