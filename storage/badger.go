package storage

import (
	"log"

	"github.com/davepartner/go-blockchain/blockchain"
	"github.com/dgraph-io/badger/v3"
)

//name: realDavePartner@gmail.com : Braintemple Tutorial TV, github.com/davepartner/go-blockchain
/*
1.OpenDB: This will open a BadgerDB instance at the specified path for storing the blockchain data
2. CloseDB: This will close the BadgerDB instance safely to ensure there is no data corruption
3. SaveBlock: This will save a serialized block to the BadgerDB instance using its hash as key
4. GetBlock: This will get a block from the BadgerDB instance by its hash, deserialize it, and return the block
*/

// BlockchainDB manages the blockchain stroage using BadgerDB
type BlockchainDB struct {
	DB *badger.DB
}

// OpenDB opens a BadgerDB instance at the specified path for storing the blockchain data
func OpenDB(path string) *BlockchainDB {
	opts := badger.DefaultOptions(path)
	//
	db, err := badger.Open(opts) //open the database
	//
	if err != nil {
		log.Panic(err)
	}
	return &BlockchainDB{DB: db}
}

// SaveBlock: store a block in the badgerDB
func (bdb *BlockchainDB) SaveBlock(block *blockchain.Block) error {
	return bdb.DB.Update(func(txn *badger.Txn) error { //update-only transaction
		err := txn.Set(block.Hash, block.Serialize())
		if err != nil {
			return err
		}
		return nil
	})
}

// GetBlock: get a block from the badgerDB
func (bdb *BlockchainDB) GetBlock(hash []byte) (*blockchain.Block, error) {
	var block *blockchain.Block
	err := bdb.DB.View(func(txn *badger.Txn) error { //read-only transaction
		item, err := txn.Get(hash)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			block = blockchain.DeserialzeBlock(val) //deserialize the block
			return nil
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return block, nil
}

// CloseDB: close the BadgerDB instance safely to ensure there is no data corruption
func (bdb *BlockchainDB) CloseDB() {
	err := bdb.DB.Close() //close the database
	if err != nil {
		log.Panic(err)
	}
}
