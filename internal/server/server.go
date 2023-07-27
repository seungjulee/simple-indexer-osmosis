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
			Txs: b.TXs,
			NumTxs: int64(b.NumTXs),
		})
	}

	return &pb.GetBlocksByProposerResponse{
		Blocks: respBlocks,
	}, nil
}