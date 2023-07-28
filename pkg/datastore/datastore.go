package datastore

import (
	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore/model"
)

type Datastore interface {
	SaveBlock(blk *model.Block) error
	SaveNetInfoAndPeer(netinfo *model.NetInfo, peers []model.Peer) error

	GetBlockByHeight(height int64) (*model.Block, error)
	GetBlocksByProposer(address string) ([]model.Block, error)
	GetNumberOfTXsInLastNBlocks(n int) ([]model.Block, error)
	GetTopNPeersByScoreInLastNBlocks(nPeers, nBlocks int) (*PeersByBlockHeight, error)
}