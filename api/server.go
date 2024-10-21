package api

import (
	"net/http"

	"github.com/davepartner/go-blockchain/blockchain"
	"github.com/davepartner/go-blockchain/contracts"
	"github.com/davepartner/go-blockchain/storage"
	"github.com/labstack/echo/v4"
)

/*
//name: realDavePartner@gmail.com : Braintemple Tutorial TV, github.com/davepartner/go-blockchain

1. Startserver: initializes the Echo Framework server and sets up routs for handling transactions
retrieving bocks, and deploying/executing smart contracts.

2. handleTransaction: Accepts transaction via POST and process them (the logic can be explabded to create and sign transactions)
3. handleGetBlock: Fetches a block by its hash
4. handleDeployContract: Deploys a smart contract to the blockchain by validating its code and simulating storage on the blockchain
5. handleExecuteContract: Executes a deployed smart contract on the blockchain with the provided input
*/

// Blockchain and storage references (global for simplicity)
var bc *blockchain.Blockchain
var db *storage.BlockchainDB

// StartServer: initializes and starts the API server (echo framework)
func StartServer(bcInstance *blockchain.Blockchain, dbInstance *storage.BlockchainDB) {
	bc = bcInstance
	db = dbInstance

	e := echo.New()

	//Define API routes
	e.POST("/transaction", handleTransaction) //create a new transaction
	e.GET("/block/:hash", handleGetBlock) //retrieve a block by its hash
	e.POST("/contract", handleDeployContract) //deploy a new smart contract
	e.POST("/contract/execute", handleExecuteContract) //execute a deployed smart contract

	//Start the server
	e.Logger.Fatal(e.Start(":1323"))
}

// handleTransaction: process a new transaction
func handleTransaction(c echo.Context) error {
	//For simplicity, this example assumes that transaction details are passed as query parameters
	from := c.QueryParam("from")
	to := c.QueryParam("to")
	amount := c.QueryParam("amount")

	//TODO: create and broadcast a new trasnaction (The below is just a placeholder)
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Transaction sumbitted",
		"from":    from,
		"to":      to,
		"amount":  amount,
	})

}

// handleGetBlock: retrieves a block by its hash
func handleGetBlock(c echo.Context) error {
	hash := c.Param("hash")

	//convert the hash from string to bytes and fetch the block from the blockchain
	block, err := bc.GetBlock([]byte(hash))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Block not found",
		})
	}
	return c.JSON(http.StatusOK, block)
}

// handleDeployContract: deploys a new smart contract
func handleDeployContract(c echo.Context) error {
	id := c.QueryParam("id")
	code := c.QueryParam("code")

	contract := contracts.NewSmartContract(id, code)
	err := contract.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	//TODO: save/deploy the contract to the blockchain (this is just a placeholder)
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Contract deployed",
		"id":      id,
	})
}

// handleExecuteContract: executes a deployed smart contract
func handleExecuteContract(c echo.Context) error {
	id := c.QueryParam("id")
	input := map[string]interface{}{}

	//TODO: fetch the contract from the blockchain and execute it
	//for now, this is just a placeholder
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Contract executed",
		"id":      id,
		"input":   input,
	})
}
