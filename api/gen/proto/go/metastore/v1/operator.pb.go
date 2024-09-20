// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: metastore/v1/operator.proto

package metastorev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// State values are chosen to match the Hashicorp Raft library states. See:
// https://github.com/hashicorp/raft/blob/42d34464b2d203e389e11ed6d43db698792c0604/state.go#L15-L27.
type State int32

const (
	State_Follower  State = 0
	State_Candidate State = 1
	State_Leader    State = 2
	State_Shutdown  State = 3
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "Follower",
		1: "Candidate",
		2: "Leader",
		3: "Shutdown",
	}
	State_value = map[string]int32{
		"Follower":  0,
		"Candidate": 1,
		"Leader":    2,
		"Shutdown":  3,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_metastore_v1_operator_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_metastore_v1_operator_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_metastore_v1_operator_proto_rawDescGZIP(), []int{0}
}

// Suffrage values are chosen to match the Hashicorp Raft library suffrage
// values. See:
// https://github.com/hashicorp/raft/blob/42d34464b2d203e389e11ed6d43db698792c0604/configuration.go#L12-L24.
type Suffrage int32

const (
	Suffrage_Voter    Suffrage = 0
	Suffrage_NonVoter Suffrage = 1
	Suffrage_Staging  Suffrage = 2
)

// Enum value maps for Suffrage.
var (
	Suffrage_name = map[int32]string{
		0: "Voter",
		1: "NonVoter",
		2: "Staging",
	}
	Suffrage_value = map[string]int32{
		"Voter":    0,
		"NonVoter": 1,
		"Staging":  2,
	}
)

func (x Suffrage) Enum() *Suffrage {
	p := new(Suffrage)
	*p = x
	return p
}

func (x Suffrage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Suffrage) Descriptor() protoreflect.EnumDescriptor {
	return file_metastore_v1_operator_proto_enumTypes[1].Descriptor()
}

func (Suffrage) Type() protoreflect.EnumType {
	return &file_metastore_v1_operator_proto_enumTypes[1]
}

func (x Suffrage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Suffrage.Descriptor instead.
func (Suffrage) EnumDescriptor() ([]byte, []int) {
	return file_metastore_v1_operator_proto_rawDescGZIP(), []int{1}
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitIndex      uint64 `protobuf:"varint,1,opt,name=commit_index,json=commitIndex,proto3" json:"commit_index,omitempty"`
	AppliedIndex     uint64 `protobuf:"varint,2,opt,name=applied_index,json=appliedIndex,proto3" json:"applied_index,omitempty"`
	LastIndex        uint64 `protobuf:"varint,3,opt,name=last_index,json=lastIndex,proto3" json:"last_index,omitempty"`
	FsmPendingLength uint64 `protobuf:"varint,4,opt,name=fsm_pending_length,json=fsmPendingLength,proto3" json:"fsm_pending_length,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_operator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_operator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_metastore_v1_operator_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetCommitIndex() uint64 {
	if x != nil {
		return x.CommitIndex
	}
	return 0
}

func (x *Log) GetAppliedIndex() uint64 {
	if x != nil {
		return x.AppliedIndex
	}
	return 0
}

func (x *Log) GetLastIndex() uint64 {
	if x != nil {
		return x.LastIndex
	}
	return 0
}

func (x *Log) GetFsmPendingLength() uint64 {
	if x != nil {
		return x.FsmPendingLength
	}
	return 0
}

type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastIndex uint64 `protobuf:"varint,1,opt,name=last_index,json=lastIndex,proto3" json:"last_index,omitempty"`
	LastTerm  uint64 `protobuf:"varint,2,opt,name=last_term,json=lastTerm,proto3" json:"last_term,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_operator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_operator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_metastore_v1_operator_proto_rawDescGZIP(), []int{1}
}

func (x *Snapshot) GetLastIndex() uint64 {
	if x != nil {
		return x.LastIndex
	}
	return 0
}

func (x *Snapshot) GetLastTerm() uint64 {
	if x != nil {
		return x.LastTerm
	}
	return 0
}

type Protocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version            uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	MinVersion         uint64 `protobuf:"varint,2,opt,name=min_version,json=minVersion,proto3" json:"min_version,omitempty"`
	MaxVersion         uint64 `protobuf:"varint,3,opt,name=max_version,json=maxVersion,proto3" json:"max_version,omitempty"`
	MinSnapshotVersion uint64 `protobuf:"varint,4,opt,name=min_snapshot_version,json=minSnapshotVersion,proto3" json:"min_snapshot_version,omitempty"`
	MaxSnapshotVersion uint64 `protobuf:"varint,5,opt,name=max_snapshot_version,json=maxSnapshotVersion,proto3" json:"max_snapshot_version,omitempty"`
}

func (x *Protocol) Reset() {
	*x = Protocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_operator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Protocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Protocol) ProtoMessage() {}

func (x *Protocol) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_operator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Protocol.ProtoReflect.Descriptor instead.
func (*Protocol) Descriptor() ([]byte, []int) {
	return file_metastore_v1_operator_proto_rawDescGZIP(), []int{2}
}

func (x *Protocol) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Protocol) GetMinVersion() uint64 {
	if x != nil {
		return x.MinVersion
	}
	return 0
}

func (x *Protocol) GetMaxVersion() uint64 {
	if x != nil {
		return x.MaxVersion
	}
	return 0
}

func (x *Protocol) GetMinSnapshotVersion() uint64 {
	if x != nil {
		return x.MinSnapshotVersion
	}
	return 0
}

func (x *Protocol) GetMaxSnapshotVersion() uint64 {
	if x != nil {
		return x.MaxSnapshotVersion
	}
	return 0
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address  string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Suffrage Suffrage `protobuf:"varint,3,opt,name=suffrage,proto3,enum=metastore.v1.Suffrage" json:"suffrage,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_operator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_operator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_metastore_v1_operator_proto_rawDescGZIP(), []int{3}
}

func (x *Peer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Peer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Peer) GetSuffrage() Suffrage {
	if x != nil {
		return x.Suffrage
	}
	return Suffrage_Voter
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_operator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_operator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_metastore_v1_operator_proto_rawDescGZIP(), []int{4}
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State    State  `protobuf:"varint,2,opt,name=state,proto3,enum=metastore.v1.State" json:"state,omitempty"`
	LeaderId string `protobuf:"bytes,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	// True if the state reported by this node matches the state the cluster
	// reports. For example, if this node claims to be a leader, but the rest of
	// the cluster disagrees, this value will be false.
	IsStateVerified bool `protobuf:"varint,4,opt,name=is_state_verified,json=isStateVerified,proto3" json:"is_state_verified,omitempty"`
	// Unix timestamp in milliseconds of when the leader last contacted this node.
	LastLeaderContact int64     `protobuf:"varint,5,opt,name=last_leader_contact,json=lastLeaderContact,proto3" json:"last_leader_contact,omitempty"`
	Term              uint64    `protobuf:"varint,6,opt,name=term,proto3" json:"term,omitempty"`
	Suffrage          Suffrage  `protobuf:"varint,7,opt,name=suffrage,proto3,enum=metastore.v1.Suffrage" json:"suffrage,omitempty"`
	Log               *Log      `protobuf:"bytes,8,opt,name=log,proto3" json:"log,omitempty"`
	Snapshot          *Snapshot `protobuf:"bytes,9,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	Protocol          *Protocol `protobuf:"bytes,10,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Peers             []*Peer   `protobuf:"bytes,11,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_operator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_operator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_metastore_v1_operator_proto_rawDescGZIP(), []int{5}
}

func (x *InfoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InfoResponse) GetState() State {
	if x != nil {
		return x.State
	}
	return State_Follower
}

func (x *InfoResponse) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *InfoResponse) GetIsStateVerified() bool {
	if x != nil {
		return x.IsStateVerified
	}
	return false
}

func (x *InfoResponse) GetLastLeaderContact() int64 {
	if x != nil {
		return x.LastLeaderContact
	}
	return 0
}

func (x *InfoResponse) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *InfoResponse) GetSuffrage() Suffrage {
	if x != nil {
		return x.Suffrage
	}
	return Suffrage_Voter
}

func (x *InfoResponse) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *InfoResponse) GetSnapshot() *Snapshot {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

func (x *InfoResponse) GetProtocol() *Protocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

func (x *InfoResponse) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

var File_metastore_v1_operator_proto protoreflect.FileDescriptor

var file_metastore_v1_operator_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d,
	0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x9a, 0x01, 0x0a, 0x03,
	0x4c, 0x6f, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x73,
	0x6d, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x66, 0x73, 0x6d, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x46, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x65, 0x72, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x65, 0x72, 0x6d,
	0x22, 0xca, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d, 0x69,
	0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d,
	0x61, 0x78, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x69, 0x6e,
	0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x6d, 0x69, 0x6e, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x6d,
	0x61, 0x78, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x64, 0x0a,
	0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x32, 0x0a, 0x08, 0x73, 0x75, 0x66, 0x66, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x66, 0x66, 0x72, 0x61, 0x67, 0x65, 0x52, 0x08, 0x73, 0x75, 0x66, 0x66, 0x72,
	0x61, 0x67, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xc1, 0x03, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x69,
	0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x32, 0x0a, 0x08, 0x73,
	0x75, 0x66, 0x66, 0x72, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x66,
	0x66, 0x72, 0x61, 0x67, 0x65, 0x52, 0x08, 0x73, 0x75, 0x66, 0x66, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x23, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x03, 0x6c, 0x6f, 0x67, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x08,
	0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x28, 0x0a, 0x05,
	0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2a, 0x3e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0c, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74,
	0x64, 0x6f, 0x77, 0x6e, 0x10, 0x03, 0x2a, 0x30, 0x0a, 0x08, 0x53, 0x75, 0x66, 0x66, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x4e, 0x6f, 0x6e, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x32, 0x52, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xba, 0x01, 0x0a,
	0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d,
	0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58,
	0xaa, 0x02, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x18, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_metastore_v1_operator_proto_rawDescOnce sync.Once
	file_metastore_v1_operator_proto_rawDescData = file_metastore_v1_operator_proto_rawDesc
)

func file_metastore_v1_operator_proto_rawDescGZIP() []byte {
	file_metastore_v1_operator_proto_rawDescOnce.Do(func() {
		file_metastore_v1_operator_proto_rawDescData = protoimpl.X.CompressGZIP(file_metastore_v1_operator_proto_rawDescData)
	})
	return file_metastore_v1_operator_proto_rawDescData
}

var file_metastore_v1_operator_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_metastore_v1_operator_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metastore_v1_operator_proto_goTypes = []any{
	(State)(0),           // 0: metastore.v1.State
	(Suffrage)(0),        // 1: metastore.v1.Suffrage
	(*Log)(nil),          // 2: metastore.v1.Log
	(*Snapshot)(nil),     // 3: metastore.v1.Snapshot
	(*Protocol)(nil),     // 4: metastore.v1.Protocol
	(*Peer)(nil),         // 5: metastore.v1.Peer
	(*InfoRequest)(nil),  // 6: metastore.v1.InfoRequest
	(*InfoResponse)(nil), // 7: metastore.v1.InfoResponse
}
var file_metastore_v1_operator_proto_depIdxs = []int32{
	1, // 0: metastore.v1.Peer.suffrage:type_name -> metastore.v1.Suffrage
	0, // 1: metastore.v1.InfoResponse.state:type_name -> metastore.v1.State
	1, // 2: metastore.v1.InfoResponse.suffrage:type_name -> metastore.v1.Suffrage
	2, // 3: metastore.v1.InfoResponse.log:type_name -> metastore.v1.Log
	3, // 4: metastore.v1.InfoResponse.snapshot:type_name -> metastore.v1.Snapshot
	4, // 5: metastore.v1.InfoResponse.protocol:type_name -> metastore.v1.Protocol
	5, // 6: metastore.v1.InfoResponse.peers:type_name -> metastore.v1.Peer
	6, // 7: metastore.v1.OperatorService.Info:input_type -> metastore.v1.InfoRequest
	7, // 8: metastore.v1.OperatorService.Info:output_type -> metastore.v1.InfoResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_metastore_v1_operator_proto_init() }
func file_metastore_v1_operator_proto_init() {
	if File_metastore_v1_operator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metastore_v1_operator_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Log); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metastore_v1_operator_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Snapshot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metastore_v1_operator_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Protocol); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metastore_v1_operator_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Peer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metastore_v1_operator_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*InfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_metastore_v1_operator_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*InfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metastore_v1_operator_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metastore_v1_operator_proto_goTypes,
		DependencyIndexes: file_metastore_v1_operator_proto_depIdxs,
		EnumInfos:         file_metastore_v1_operator_proto_enumTypes,
		MessageInfos:      file_metastore_v1_operator_proto_msgTypes,
	}.Build()
	File_metastore_v1_operator_proto = out.File
	file_metastore_v1_operator_proto_rawDesc = nil
	file_metastore_v1_operator_proto_goTypes = nil
	file_metastore_v1_operator_proto_depIdxs = nil
}
