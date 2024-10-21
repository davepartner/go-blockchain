package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//name: Dave Partner: Braintemple Tutorial TV, github.com/davepartner/go-blockchain

//Block: A block constst of timestamp, transactions, previous block and v.public key
//Transaction: Contains an input and output. The input holds the sender's public key and signature and the output contains the recipient's public key and and value
//NewTransaction: Create a new transaction, signs it with ECDSA, and assign an ID
//hashTransaction: hashes the trasnaction data to create a unique ID
// serialize and deserialize: for db storage and retrieval

// Transaction resprents blockchain transaction
type Transaction struct {
	ID     []byte
	Input  []TxInput
	Output []TxOutput
}

// TxInput is the input of a transaction
type TxInput struct {
	Signature []byte
	PublicKey []byte
}

// TxOutput is the output of a transaction
type TxOutput struct {
	Value     int
	PublicKey []byte
}

// NewTransaction: Create a new transaction, signs it with ECDSA, and assign an ID
func NewTransaction(privateKey ecdsa.PrivateKey, recipient []byte, amount int) *Transaction {
	txIn := TxInput{}
	txOut := TxOutput{Value: amount, PublicKey: recipient}

	tx := Transaction{
		Input:  []TxInput{txIn},
		Output: []TxOutput{txOut},
	}

	tx.ID = tx.hashTransaction()

	//sign the transaction with the sender's private key
	r, s, err := ecdsa.Sign(rand.Reader, &privateKey, tx.ID)
	//check for errors
	if err != nil {
		log.Panic(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)
	txIn.Signature = signature
	return &tx
}

// hashTransaction: hashes the transaction data to create a unique ID
func (tx *Transaction) hashTransaction() []byte {
	var hash [32]byte
	hash = sha256.Sum256(bytes.Join([][]byte{
		tx.Input[0].PublicKey,
		tx.Output[0].PublicKey,
		[]byte(string(tx.Output[0].Value)),
	}, []byte{}))

	return hash[:]
}

// serialize(): it serializes a transaction into a byte array
func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return encoded.Bytes()
}

//Deserialize(): Deserialize a transaction from a byte array
func DeserializeTransaction(data []byte) *Transaction{
	var transaction Transaction
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&transaction)
	if err != nil {
		log.Panic(err)
	}
	return &transaction
}
