package main

import (
	"time"

	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore"
	"github.com/seungjulee/simple-indexer-osmosis/pkg/indexer"
	"github.com/seungjulee/simple-indexer-osmosis/pkg/logger"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"go.uber.org/zap"
)

type Config struct {
	RPCEndpoint string
	SqliteConfig *datastore.SqliteConfig
}

func main() {
	cfg := &Config{
		RPCEndpoint: "https://rpc-osmosis.ecostake.com:443",
		SqliteConfig: &datastore.SqliteConfig{
			SqlitePath: "./test.db",
		},
	}

	logger.Info("initializing app with config", zap.Any("config", cfg))
	client, err := rpchttp.New(cfg.RPCEndpoint, "/websocket")
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	db, err := datastore.NewSqllite(cfg.SqliteConfig)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	a := indexer.New(client, db)
	err = a.SchedulePeriodicIndex(2 * time.Second)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}
}