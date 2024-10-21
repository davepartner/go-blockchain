package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

// Block resprents each block in the blockchain
// A block consists of a timestamp, transactions, previous block hash and validator's public key
type Block struct {
	Timestamp     time.Time
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Validator     []byte //validator's public key
	Nonce         int
}

// NewBlock: Create and return a new block
// This generates a new block by hashing the previous block and the current transactions
func NewBlock(transactions []*Transaction, prevBlockHash []byte, validator []byte) *Block {
	block := &Block{
		Timestamp:     time.Now(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Validator:     validator,
	}
	block.Hash = block.calculateHash()
	return block
}

// calculateHash(): generate the hash of the block
// This calculates a SHA-256 hash o the block's data: previous hash, transactions, and timestamp
func (b *Block) calculateHash() []byte {
	var txHashes []byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.hashTransaction()...)
	}
	hash := sha256.Sum256(bytes.Join([][]byte{
		b.PrevBlockHash,
		txHashes,
		[]byte(b.Timestamp.String()),
	}, []byte{}))

	return hash[:]
}

// serialize() converts a block into a byte slice for storage
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	//if there is an error
	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}

	return result.Bytes()

}

// DeserialzeBlock(): Convert byte slice back into block
func DeserialzeBlock(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	//if there is an error
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}

	return &block
}
