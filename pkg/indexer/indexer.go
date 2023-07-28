package indexer

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	// rpc "github.com/tendermint/tendermint/rpc/client"
	// ctypes "github.com/tendermint/tendermint/rpc/coretypes"
	// jsonrpcclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"

	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore"
	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore/model"
	"github.com/seungjulee/simple-indexer-osmosis/pkg/logger"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

type Indexer interface {
	IndexLatestBlock(context.Context) error
	SchedulePeriodicIndex(time.Duration) error
}

type indexer struct {
	client *rpchttp.HTTP
	db datastore.Datastore
}



func New(client *rpchttp.HTTP, db datastore.Datastore) Indexer {
	return &indexer{
		client: client,
		db: db,
	}
}

func (a *indexer) SchedulePeriodicIndex(interval time.Duration) error {
	logger.Info(fmt.Sprintf("Schedule periodic index every %s", interval))

	blockTicker := time.NewTicker(interval)
	ctx := context.Background()
	for {
		select {
		case <-blockTicker.C:
			a.IndexLatestBlock(ctx)
		}
	}
}

func (a *indexer) IndexLatestBlock(ctx context.Context) error {
	blk, err := a.client.Block(ctx, nil)
	if err != nil {
		return err
	}

	if err := a.db.SaveBlock(&model.Block{
		BlockID: blk.BlockID.String(),
		Hash: blk.BlockID.Hash.String(),
		Height: blk.Block.Header.Height,
		ProposerAddress: blk.Block.ProposerAddress.String(),
		Time: blk.Block.Time,
		LastBlockID: blk.Block.LastBlockID.String(),
		TXs: string(blk.Block.Data.Txs.Hash()),
		NumTXs: len(blk.Block.Data.Txs),
	}); err != nil {
		return err
	}

	netInfo, err := a.client.NetInfo(ctx)

	niModel := &model.NetInfo{
		NPeer: netInfo.NPeers,
		BlockHeight: int(blk.Block.Header.Height),
	}

	var peerModels []model.Peer
	for _, peer := range netInfo.Peers {
		peerModels = append(peerModels, model.Peer{
			BlockHeight: int(blk.Block.Header.Height),
			RemoteIP: peer.RemoteIP,
			DefaultNodeID: string(peer.NodeInfo.DefaultNodeID),
			Score: getPeerScore(),
		})
	}

	if err := a.db.SaveNetInfoAndPeer(niModel, peerModels); err != nil {
		return err
	}

	return nil
}


// TODO: derive an algorithm to get the peer score
func getPeerScore() int{
    rand.Seed(time.Now().UnixNano())
    min := 0
    max := 100
    return rand.Intn(max - min + 1) + min
}