package blockchain

import (
	"log"
	"crypto/rand"
	"math/big"
	//"math/rand"
	//"time"
)

//name: realDavePartner@gmail.com : Braintemple Tutorial TV, github.com/davepartner/go-blockchain

/*
1. PoSValidator: represents a validator in the proof of stake system with a public key and stake
2. ProofOfStake(): Selects a validator based on their stake. The more stake a validator has,
the higher the probability of being chosen to validate a block
*/

// PoSValidator: represents a validator in the proof of stake system with a public key and stake
type PoSValidator struct {
	PublicKey []byte
	Stake     int
}

/*
ProofOfStake(): Selects a validator based on their stake. The more stake a validator has,
the higher the probability of being chosen to validate a block
*/
func ProofOfStake(validators map[string]*PoSValidator) string {
	totalStake := 0
	for _, validaor := range validators {
		totalStake += validaor.Stake
	}

	//select a validator randomly based on their stake
	randomBig, err := rand.Int(rand.Reader, big.NewInt(int64(totalStake)))

	if err != nil {
		log.Panic(err)
	}
	random := randomBig.Int64()

	//select a validator based on their stake
	for _, validator := range validators {
		random -= int64(validator.Stake)
		if random <= 0 {
			return string(validator.PublicKey)
		}
	}
   //	log.Panic("Unable to find a validator")
	log.Panic("Unable to find a validator")
	return ""
}
