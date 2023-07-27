package datastore

import (
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

	// // Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// // Read
	// var product Product
	// // db.First(&product, 1) // find product with integer primary key
	// tx := db.First(&product, "code = ?", "D42") // find product with code D42
	// tx.Commit();

	// fmt.Println(product)
	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)

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