package datastore

import "github.com/seungjulee/simple-indexer-osmosis/pkg/datastore/model"

type PeersByBlockHeight struct {
	PeersByBlockHeight map[int][]model.Peer
}