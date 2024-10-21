package blockchain

import "fmt"

// Block represents a block in the blockchain

type Blockchain struct {
	Blocks []*Block //in-memory chain of blocks
}

// GetBlocks retrieve a block from the in-memory chain by its hash
func (bc *Blockchain) GetBlock(hash []byte) ([]*Block, error) {
	for _, block := range bc.Blocks {
		if string(block.Hash) == string(hash) {
			return []*Block{block}, nil
		}
	}
	return nil, fmt.Errorf("block not found")
}

//cosmos 
//udemy.com/user/davepartner

func NewBlockchain(genesisBlock *Block) *Blockchain {
	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}

// AddBlock: adds a new block to the chain
func (bc *Blockchain) AddBlock(transactions []*Transaction, validator []byte) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1] //
	newBlock := NewBlock(transactions, prevBlock.Hash, validator)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}
