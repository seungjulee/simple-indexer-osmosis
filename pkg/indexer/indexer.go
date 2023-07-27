package indexer

import (
	"context"
	"time"

	// rpc "github.com/tendermint/tendermint/rpc/client"
	// ctypes "github.com/tendermint/tendermint/rpc/coretypes"
	// jsonrpcclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore"
	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore/model"
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

	// wsClient, err := jsonrpcclient.NewWS(config.RPCEndpoint, "/gateway/cos4/rpc/b3fbc22883fc3fe61254616b40fc568b")
	// if err != nil {
	// 	return nil, err
	// }
	// c, err := jsonrpcclient.DefaultHTTPClient(config.RPCEndpoint)
	// if err != nil {
	// 	return nil, err
	// }
	// //
	// wsOption := rpcHttp.DefaultWSOptions()
	// wsOption.Path = "/gateway/cos4/rpc/b3fbc22883fc3fe61254616b40fc568b"
	// wsClient, err := rpcHttp.NewWithClientAndWSOptions(config.RPCEndpoint, c, wsOption)
	// if err != nil {
	// 	return nil, err
	// }

	return &indexer{
		client: client,
		db: db,
	}
}

func (a *indexer) SchedulePeriodicIndex(interval time.Duration) error {
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

	// netInfo, err := a.client.NetInfo(ctx)
	// fmt.Println(netInfo.NPeers)

	// res, err := a.db.GetBlockByHeight(blk.Block.Header.Height)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(res)

	// query := "tm.event = 'Tx' AND tx.height = 3"
	// txs, err := a.wsClient.Subscribe(ctx, "test-client", query)
	// if err != nil {
	// 	return err
	// }

	// go func() {
	// for e := range txs {
	// fmt.Println("got ", e.Data.(types.EventDataTx))
	// }
	// }()

	// fmt.Println("!sub")
	// blockChannel, err := a.wsClient.Subscribe(ctx, "tm.event = 'NewBlock'")
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("!sub2")


	// txChannel, err := a.wsClient.Subscribe(ctx, "tm.event = 'Tx'")
	// if err != nil {
	// 	return err
	// }

	return nil
}

