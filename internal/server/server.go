package server

import (
	"context"
	"fmt"

	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore"
	pb "github.com/seungjulee/simple-indexer-osmosis/rpc"
	"github.com/twitchtv/twirp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Server implements SimpleOsmosisExplorer rpc service
type Server struct {
	Datastore datastore.Datastore
}

func NewExplorerServer(ds datastore.Datastore) pb.SimpleOsmosisExplorer {
	return &Server{
		Datastore: ds,
	}
}

func (s *Server) GetBlocksByProposer(ctx context.Context, req *pb.GetBlocksByProposerRequest) (*pb.GetBlocksByProposerResponse, error) {
    if req.Address == "" {
        return nil, twirp.InvalidArgumentError("address", "'address' is empty")
    }

	blocks, err := s.Datastore.GetBlocksByProposer(req.Address)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	if len(blocks) == 0 {
		return nil, twirp.NotFoundError(fmt.Sprintf("could not find '%s'", req.Address))
	}

	var respBlocks []*pb.Block

	for _, b := range blocks {
		respBlocks = append(respBlocks, &pb.Block{
			Hash: b.Hash,
			Height: b.Height,
			ProposerAddress: b.ProposerAddress,
			Time: timestamppb.New(b.Time),
			// Txs: b.TXs,
			NumTxs: int64(b.NumTXs),
		})
	}

	return &pb.GetBlocksByProposerResponse{
		Blocks: respBlocks,
	}, nil
}

func (s *Server) GetNumberOfTXsInLastNBlocks(ctx context.Context, req *pb.GetNumberOfTXsInLastNBlocksRequest) (*pb.GetNumberOfTXsInLastNBlocksResponse, error) {
	if req.N == 0 {
		return nil, twirp.RequiredArgumentError("n")
	}

	blocks, err := s.Datastore.GetNumberOfTXsInLastNBlocks(int(req.N))
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}

	var totalNumTxs int64
	var respBlocks []*pb.Block
	for _, b := range blocks {
		totalNumTxs += int64(b.NumTXs)
		respBlocks = append(respBlocks, &pb.Block{
			Hash: b.Hash,
			Height: b.Height,
			ProposerAddress: b.ProposerAddress,
			Time: timestamppb.New(b.Time),
			// Txs: b.TXs,
			NumTxs: int64(b.NumTXs),
		})
	}

	return &pb.GetNumberOfTXsInLastNBlocksResponse{
		TotalNumTxs: totalNumTxs,
		Blocks: respBlocks,
	}, nil
}

func (s *Server) GetTopNPeersByScoreInLastNBlocks(ctx context.Context, req *pb.GetTopNPeersByScoreInLastNBlocksRequest) (*pb.GetTopNPeersByScoreInLastNBlocksResponse, error) {
	if req.NBlock == 0 {
		return nil, twirp.RequiredArgumentError("n_block")
	}

	if req.NPeer == 0 {
		return nil, twirp.RequiredArgumentError("n_peer")
	}

	peersByBlock, err := s.Datastore.GetTopNPeersByScoreInLastNBlocks(int(req.NPeer), int(req.NBlock))
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}

	var peersByHeight map[int64]*pb.Peers
	peersByHeight = make(map[int64]*pb.Peers)
	for k, v := range peersByBlock.PeersByBlockHeight {
		peersByHeight[int64(k)] = &pb.Peers{
			Peers: []*pb.Peer{},
		}

		for _, p := range v {
			peersByHeight[int64(k)].Peers = append(peersByHeight[int64(k)].Peers, &pb.Peer{
				BlockHeight: int64(k),
				RemoteIp: p.RemoteIP,
				DefaultNodeId: p.DefaultNodeID,
				Score: int64(p.Score),
			})
		}
	}

	return &pb.GetTopNPeersByScoreInLastNBlocksResponse{
		TopNPeersByBlockHeight: peersByHeight,
	}, nil
}