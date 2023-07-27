package main

import (
	"net/http"

	"github.com/seungjulee/simple-indexer-osmosis/internal/server"
	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore"
	"github.com/seungjulee/simple-indexer-osmosis/rpc"
)

type Config struct {
	SqliteConfig *datastore.SqliteConfig
}

func main() {
	cfg := &Config{
		SqliteConfig: &datastore.SqliteConfig{
			SqlitePath: "./test.db",
		},
	}
	db, err := datastore.NewSqllite(cfg.SqliteConfig)
	if err != nil {
		panic(err)
	}
	sv := server.NewExplorerServer(db)
	twirpHandler := rpc.NewSimpleOsmosisExplorerServer(sv)

	http.ListenAndServe(":8080", twirpHandler)
}