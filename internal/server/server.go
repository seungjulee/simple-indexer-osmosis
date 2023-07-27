package server

import (
	"context"

	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore"
	pb "github.com/seungjulee/simple-indexer-osmosis/rpc"
	"github.com/twitchtv/twirp"
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

func (s *Server) GetBlocksByValidator(ctx context.Context, req *pb.GetBlocksByValidatorRequest) (*pb.GetBlocksByValidatorResponse, error) {
    if req.Address == "" {
        return nil, twirp.InvalidArgumentError("address", "'address' is empty")
    }
    // return &pb.Hat{
    //     Inches:  size.Inches,
    //     Color: []string{"white", "black", "brown", "red", "blue"}[rand.Intn(5)],
    //     Name:  []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(4)],
    // }, nil
	// s.datastore.GetBlockByHeight()
	return nil, nil
}