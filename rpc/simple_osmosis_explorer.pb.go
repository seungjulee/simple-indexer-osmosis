// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc/simple_osmosis_explorer.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetBlocksByProposerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetBlocksByProposerRequest) Reset() {
	*x = GetBlocksByProposerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksByProposerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksByProposerRequest) ProtoMessage() {}

func (x *GetBlocksByProposerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocksByProposerRequest.ProtoReflect.Descriptor instead.
func (*GetBlocksByProposerRequest) Descriptor() ([]byte, []int) {
	return file_rpc_simple_osmosis_explorer_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlocksByProposerRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId         string                 `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Hash            string                 `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Height          int64                  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	ProposerAddress string                 `protobuf:"bytes,4,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty"`
	Time            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	Txs             string                 `protobuf:"bytes,6,opt,name=txs,proto3" json:"txs,omitempty"`
	NumTxs          int64                  `protobuf:"varint,7,opt,name=num_txs,json=numTxs,proto3" json:"num_txs,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_rpc_simple_osmosis_explorer_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Block) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Block) GetProposerAddress() string {
	if x != nil {
		return x.ProposerAddress
	}
	return ""
}

func (x *Block) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Block) GetTxs() string {
	if x != nil {
		return x.Txs
	}
	return ""
}

func (x *Block) GetNumTxs() int64 {
	if x != nil {
		return x.NumTxs
	}
	return 0
}

type GetBlocksByProposerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *GetBlocksByProposerResponse) Reset() {
	*x = GetBlocksByProposerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksByProposerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksByProposerResponse) ProtoMessage() {}

func (x *GetBlocksByProposerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocksByProposerResponse.ProtoReflect.Descriptor instead.
func (*GetBlocksByProposerResponse) Descriptor() ([]byte, []int) {
	return file_rpc_simple_osmosis_explorer_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlocksByProposerResponse) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type GetNumberOfTXsInLastNBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int64 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *GetNumberOfTXsInLastNBlocksRequest) Reset() {
	*x = GetNumberOfTXsInLastNBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumberOfTXsInLastNBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumberOfTXsInLastNBlocksRequest) ProtoMessage() {}

func (x *GetNumberOfTXsInLastNBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumberOfTXsInLastNBlocksRequest.ProtoReflect.Descriptor instead.
func (*GetNumberOfTXsInLastNBlocksRequest) Descriptor() ([]byte, []int) {
	return file_rpc_simple_osmosis_explorer_proto_rawDescGZIP(), []int{3}
}

func (x *GetNumberOfTXsInLastNBlocksRequest) GetN() int64 {
	if x != nil {
		return x.N
	}
	return 0
}

type GetNumberOfTXsInLastNBlocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalNumTxs int64    `protobuf:"varint,1,opt,name=total_num_txs,json=totalNumTxs,proto3" json:"total_num_txs,omitempty"`
	Blocks      []*Block `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *GetNumberOfTXsInLastNBlocksResponse) Reset() {
	*x = GetNumberOfTXsInLastNBlocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumberOfTXsInLastNBlocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumberOfTXsInLastNBlocksResponse) ProtoMessage() {}

func (x *GetNumberOfTXsInLastNBlocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumberOfTXsInLastNBlocksResponse.ProtoReflect.Descriptor instead.
func (*GetNumberOfTXsInLastNBlocksResponse) Descriptor() ([]byte, []int) {
	return file_rpc_simple_osmosis_explorer_proto_rawDescGZIP(), []int{4}
}

func (x *GetNumberOfTXsInLastNBlocksResponse) GetTotalNumTxs() int64 {
	if x != nil {
		return x.TotalNumTxs
	}
	return 0
}

func (x *GetNumberOfTXsInLastNBlocksResponse) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type GetTopNPeersByScoreInLastNBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NPeer  int64 `protobuf:"varint,1,opt,name=n_peer,json=nPeer,proto3" json:"n_peer,omitempty"`
	NBlock int64 `protobuf:"varint,2,opt,name=n_block,json=nBlock,proto3" json:"n_block,omitempty"`
}

func (x *GetTopNPeersByScoreInLastNBlocksRequest) Reset() {
	*x = GetTopNPeersByScoreInLastNBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopNPeersByScoreInLastNBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopNPeersByScoreInLastNBlocksRequest) ProtoMessage() {}

func (x *GetTopNPeersByScoreInLastNBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopNPeersByScoreInLastNBlocksRequest.ProtoReflect.Descriptor instead.
func (*GetTopNPeersByScoreInLastNBlocksRequest) Descriptor() ([]byte, []int) {
	return file_rpc_simple_osmosis_explorer_proto_rawDescGZIP(), []int{5}
}

func (x *GetTopNPeersByScoreInLastNBlocksRequest) GetNPeer() int64 {
	if x != nil {
		return x.NPeer
	}
	return 0
}

func (x *GetTopNPeersByScoreInLastNBlocksRequest) GetNBlock() int64 {
	if x != nil {
		return x.NBlock
	}
	return 0
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHeight   int64  `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	RemoteIp      string `protobuf:"bytes,2,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
	DefaultNodeId string `protobuf:"bytes,3,opt,name=default_node_id,json=defaultNodeId,proto3" json:"default_node_id,omitempty"`
	Score         int64  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[6]
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
	return file_rpc_simple_osmosis_explorer_proto_rawDescGZIP(), []int{6}
}

func (x *Peer) GetBlockHeight() int64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *Peer) GetRemoteIp() string {
	if x != nil {
		return x.RemoteIp
	}
	return ""
}

func (x *Peer) GetDefaultNodeId() string {
	if x != nil {
		return x.DefaultNodeId
	}
	return ""
}

func (x *Peer) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type Peers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*Peer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *Peers) Reset() {
	*x = Peers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peers) ProtoMessage() {}

func (x *Peers) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peers.ProtoReflect.Descriptor instead.
func (*Peers) Descriptor() ([]byte, []int) {
	return file_rpc_simple_osmosis_explorer_proto_rawDescGZIP(), []int{7}
}

func (x *Peers) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type GetTopNPeersByScoreInLastNBlocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopNPeersByBlockHeight map[int64]*Peers `protobuf:"bytes,2,rep,name=top_n_peers_by_block_height,json=topNPeersByBlockHeight,proto3" json:"top_n_peers_by_block_height,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetTopNPeersByScoreInLastNBlocksResponse) Reset() {
	*x = GetTopNPeersByScoreInLastNBlocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopNPeersByScoreInLastNBlocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopNPeersByScoreInLastNBlocksResponse) ProtoMessage() {}

func (x *GetTopNPeersByScoreInLastNBlocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_simple_osmosis_explorer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopNPeersByScoreInLastNBlocksResponse.ProtoReflect.Descriptor instead.
func (*GetTopNPeersByScoreInLastNBlocksResponse) Descriptor() ([]byte, []int) {
	return file_rpc_simple_osmosis_explorer_proto_rawDescGZIP(), []int{8}
}

func (x *GetTopNPeersByScoreInLastNBlocksResponse) GetTopNPeersByBlockHeight() map[int64]*Peers {
	if x != nil {
		return x.TopNPeersByBlockHeight
	}
	return nil
}

var File_rpc_simple_osmosis_explorer_proto protoreflect.FileDescriptor

var file_rpc_simple_osmosis_explorer_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0xd4, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x78, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x78, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x78, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x54, 0x78, 0x73, 0x22, 0x41, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x32, 0x0a, 0x22, 0x47,
	0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x58, 0x73, 0x49, 0x6e, 0x4c,
	0x61, 0x73, 0x74, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x6e, 0x22,
	0x6d, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x58,
	0x73, 0x49, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x78, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x54, 0x78, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x59,
	0x0a, 0x27, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x42, 0x79,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x5f, 0x70,
	0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x50, 0x65, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x84, 0x01, 0x0a, 0x04, 0x50, 0x65,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x49, 0x70, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x22, 0x28, 0x0a, 0x05, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x65, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x8a, 0x02, 0x0a, 0x28, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x49, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x1b, 0x74, 0x6f, 0x70, 0x5f,
	0x6e, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x42, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x6f, 0x70,
	0x4e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x42, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x74, 0x6f, 0x70, 0x4e, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x42, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x1a, 0x55, 0x0a, 0x1b, 0x54, 0x6f, 0x70, 0x4e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x42, 0x79, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xe4, 0x02, 0x0a, 0x15, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x4f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65,
	0x72, 0x12, 0x58, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x79,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x58, 0x73, 0x49, 0x6e, 0x4c,
	0x61, 0x73, 0x74, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x27, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x58, 0x73, 0x49,
	0x6e, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x58, 0x73, 0x49, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7f, 0x0a,
	0x20, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x12, 0x2c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x42, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x4c, 0x61, 0x73,
	0x74, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x50, 0x65, 0x65,
	0x72, 0x73, 0x42, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x4e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_simple_osmosis_explorer_proto_rawDescOnce sync.Once
	file_rpc_simple_osmosis_explorer_proto_rawDescData = file_rpc_simple_osmosis_explorer_proto_rawDesc
)

func file_rpc_simple_osmosis_explorer_proto_rawDescGZIP() []byte {
	file_rpc_simple_osmosis_explorer_proto_rawDescOnce.Do(func() {
		file_rpc_simple_osmosis_explorer_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_simple_osmosis_explorer_proto_rawDescData)
	})
	return file_rpc_simple_osmosis_explorer_proto_rawDescData
}

var file_rpc_simple_osmosis_explorer_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_rpc_simple_osmosis_explorer_proto_goTypes = []interface{}{
	(*GetBlocksByProposerRequest)(nil),              // 0: rpc.GetBlocksByProposerRequest
	(*Block)(nil),                                   // 1: rpc.Block
	(*GetBlocksByProposerResponse)(nil),             // 2: rpc.GetBlocksByProposerResponse
	(*GetNumberOfTXsInLastNBlocksRequest)(nil),      // 3: rpc.GetNumberOfTXsInLastNBlocksRequest
	(*GetNumberOfTXsInLastNBlocksResponse)(nil),     // 4: rpc.GetNumberOfTXsInLastNBlocksResponse
	(*GetTopNPeersByScoreInLastNBlocksRequest)(nil), // 5: rpc.GetTopNPeersByScoreInLastNBlocksRequest
	(*Peer)(nil),  // 6: rpc.Peer
	(*Peers)(nil), // 7: rpc.Peers
	(*GetTopNPeersByScoreInLastNBlocksResponse)(nil), // 8: rpc.GetTopNPeersByScoreInLastNBlocksResponse
	nil,                           // 9: rpc.GetTopNPeersByScoreInLastNBlocksResponse.TopNPeersByBlockHeightEntry
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_rpc_simple_osmosis_explorer_proto_depIdxs = []int32{
	10, // 0: rpc.Block.time:type_name -> google.protobuf.Timestamp
	1,  // 1: rpc.GetBlocksByProposerResponse.blocks:type_name -> rpc.Block
	1,  // 2: rpc.GetNumberOfTXsInLastNBlocksResponse.blocks:type_name -> rpc.Block
	6,  // 3: rpc.Peers.peers:type_name -> rpc.Peer
	9,  // 4: rpc.GetTopNPeersByScoreInLastNBlocksResponse.top_n_peers_by_block_height:type_name -> rpc.GetTopNPeersByScoreInLastNBlocksResponse.TopNPeersByBlockHeightEntry
	7,  // 5: rpc.GetTopNPeersByScoreInLastNBlocksResponse.TopNPeersByBlockHeightEntry.value:type_name -> rpc.Peers
	0,  // 6: rpc.SimpleOsmosisExplorer.GetBlocksByProposer:input_type -> rpc.GetBlocksByProposerRequest
	3,  // 7: rpc.SimpleOsmosisExplorer.GetNumberOfTXsInLastNBlocks:input_type -> rpc.GetNumberOfTXsInLastNBlocksRequest
	5,  // 8: rpc.SimpleOsmosisExplorer.GetTopNPeersByScoreInLastNBlocks:input_type -> rpc.GetTopNPeersByScoreInLastNBlocksRequest
	2,  // 9: rpc.SimpleOsmosisExplorer.GetBlocksByProposer:output_type -> rpc.GetBlocksByProposerResponse
	4,  // 10: rpc.SimpleOsmosisExplorer.GetNumberOfTXsInLastNBlocks:output_type -> rpc.GetNumberOfTXsInLastNBlocksResponse
	8,  // 11: rpc.SimpleOsmosisExplorer.GetTopNPeersByScoreInLastNBlocks:output_type -> rpc.GetTopNPeersByScoreInLastNBlocksResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_rpc_simple_osmosis_explorer_proto_init() }
func file_rpc_simple_osmosis_explorer_proto_init() {
	if File_rpc_simple_osmosis_explorer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_simple_osmosis_explorer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocksByProposerRequest); i {
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
		file_rpc_simple_osmosis_explorer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_rpc_simple_osmosis_explorer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocksByProposerResponse); i {
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
		file_rpc_simple_osmosis_explorer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumberOfTXsInLastNBlocksRequest); i {
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
		file_rpc_simple_osmosis_explorer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumberOfTXsInLastNBlocksResponse); i {
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
		file_rpc_simple_osmosis_explorer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopNPeersByScoreInLastNBlocksRequest); i {
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
		file_rpc_simple_osmosis_explorer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_simple_osmosis_explorer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peers); i {
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
		file_rpc_simple_osmosis_explorer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopNPeersByScoreInLastNBlocksResponse); i {
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
			RawDescriptor: file_rpc_simple_osmosis_explorer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_simple_osmosis_explorer_proto_goTypes,
		DependencyIndexes: file_rpc_simple_osmosis_explorer_proto_depIdxs,
		MessageInfos:      file_rpc_simple_osmosis_explorer_proto_msgTypes,
	}.Build()
	File_rpc_simple_osmosis_explorer_proto = out.File
	file_rpc_simple_osmosis_explorer_proto_rawDesc = nil
	file_rpc_simple_osmosis_explorer_proto_goTypes = nil
	file_rpc_simple_osmosis_explorer_proto_depIdxs = nil
}
