package datastore

import (
	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore/model"
)

type Datastore interface {
	SaveBlock(blk *model.Block) error
	GetBlockByHeight(height int64) (*model.Block, error)
}



