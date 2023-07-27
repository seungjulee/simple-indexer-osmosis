package model

import (
	"time"

	"gorm.io/gorm"
)

type Block struct {
	gorm.Model

	BlockID string `gorm:"primaryKey"`
	Hash string `gorm:"uniqueIndex"`
	Height int64 `gorm:"index:,sort:desc"`
	ProposerAddress string `gorm:"index"`
	Time time.Time
	LastBlockID string

	TXs string
	NumTXs int


	// // hashes of block data
	// LastCommitHash []byte `json:"last_commit_hash"` // commit from validators from the last block
	// DataHash       []byte `json:"data_hash"`        // transactions

	// // hashes from the app output from the prev block
	// ValidatorsHash     []byte `json:"validators_hash"`      // validators for the current block
	// NextValidatorsHash []byte `json:"next_validators_hash"` // validators for the next block
	// ConsensusHash      []byte `json:"consensus_hash"`       // consensus params for current block
	// AppHash            []byte `json:"app_hash"`             // state after txs from the previous block
}