package main

import (
	"time"

	"github.com/seungjulee/simple-indexer-osmosis/pkg/app"
	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
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
	client, err := rpchttp.New(cfg.RPCEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	db, err := datastore.NewSqllite(cfg.SqliteConfig)
	if err != nil {
		panic(err)
	}

	a := app.New(client, db)

	err = a.SchedulePeriodicIndex(2 * time.Second)
	if err != nil {
		panic(err)
	}
}