package main

import (
	"log"

	"github.com/davepartner/go-blockchain/api"
	"github.com/davepartner/go-blockchain/blockchain"
	"github.com/davepartner/go-blockchain/storage"
)

func main() {
	//Initialize the blockchin
	genesisBlock := blockchain.NewBlock([]*blockchain.Transaction{}, []byte{}, []byte("genesis-validator"))
	bc := blockchain.NewBlockchain(genesisBlock)

	//Initialize the database
	db := storage.OpenDB("./tmp/badger")
	defer db.CloseDB()

	//Save the genesis block
	err := db.SaveBlock(genesisBlock)
	if err != nil {
		log.Panic(err)
	}

	//Start the API server
	api.StartServer(bc, db)
}
