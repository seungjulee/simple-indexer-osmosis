package datastore

import (
	"errors"

	"github.com/seungjulee/simple-indexer-osmosis/pkg/datastore/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteConfig struct {
	SqlitePath string
}

func NewSqllite(cfg *SqliteConfig) (Datastore, error) {
	db, err := gorm.Open(sqlite.Open(cfg.SqlitePath), &gorm.Config{})
	if err != nil {
	  return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&model.Block{})
	db.AutoMigrate(&model.NetInfo{})
	db.AutoMigrate(&model.Peer{})

	return &sqliteDB{
		db: db,
	}, nil
}

type sqliteDB struct {
	db *gorm.DB
}

func (s *sqliteDB) SaveBlock(blk *model.Block) error {
	return s.db.Create(blk).Error
}

func (s *sqliteDB) SaveNetInfoAndPeer(netinfo *model.NetInfo, peers []model.Peer) error {
	if err := s.db.Create(netinfo).Error; err != nil {
		return err
	}

	if err := s.db.CreateInBatches(peers, len(peers)).Error; err != nil {
		return err
	}

	return nil
}

func (s *sqliteDB) GetBlockByHeight(height int64) (*model.Block, error) {
	var block model.Block
	if err := s.db.First(&block, "height = ?", height).Error; err != nil {
		// fmt.Println(err)
		return nil, err
	}
	return &block, nil
}

func (s *sqliteDB) GetBlocksByProposer(address string) ([]model.Block, error) {
	var blocks []model.Block

	if err := s.db.Where("proposer_address = ?", address).Find(&blocks).Error; err != nil {
		return nil, err
	}

	return blocks, nil
}

func (s *sqliteDB) GetNumberOfTXsInLastNBlocks(n int) ([]model.Block, error) {
	var blocks []model.Block

	if err := s.db.Order("height desc").Limit(n).Find(&blocks).Error; err != nil {
		return nil, err
	}

	return blocks, nil
}

func (s *sqliteDB) GetTopNPeersByScoreInLastNBlocks(nPeers, nBlocks int) (*PeersByBlockHeight, error) {
	var blocks []model.Block

	// TOdO: optimize the query to only output number of last
	if err := s.db.Order("height desc").Limit(nBlocks).Find(&blocks).Error; err != nil {
		return nil, err
	}
	if len(blocks) == 0 {
		return nil, errors.New("there is no recent block")
	}

	var peersByHeight map[int][]model.Peer
	peersByHeight = make(map[int][]model.Peer)
	for _, b := range blocks {
		peersByHeight[int(b.Height)] = []model.Peer{}
		var peers []model.Peer
		if err := s.db.Order("score desc").Limit(nPeers).Where("block_height = ?", b.Height).Find(&peers).Error; err != nil {
			return nil, err
		}
		peersByHeight[int(b.Height)] = peers
	}

	return &PeersByBlockHeight{
		PeersByBlockHeight: peersByHeight,
	}, nil
}