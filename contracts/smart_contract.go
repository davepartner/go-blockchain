package contracts

import (
	"errors"
	"fmt"
	"time"
)

//name: realDavePartner@gmail.com : Youtube: Braintemple Tutorial TV, github.com/davepartner/go-blockchain

//SmartContract : Represents a smart contract with an ID, code, state, and timestamp
//NewSmartContract: Initialize a new smart contract with the given ID and code. The state is  smiple key-value store
//Execute: Simulates executing the contract by logging the inpute and updating the state. in a production system, this would involve parsing and running contract code.
//Validate: a basic placeholder method that checks if the smart contract code is valid

// SmartContract : initilizes a new smart contract with the given ID and code. The state is a simple key-value store
type SmartContract struct {
	ID        string
	Code      string
	State     map[string]interface{}
	CreatedAt time.Time
}

// NewSmartContract: Initialize a new smart contract with the given ID and code. The state is  smiple key-value store
func NewSmartContract(id string, code string) *SmartContract {
	return &SmartContract{
		ID:        id,
		Code:      code,
		State:     make(map[string]interface{}),
		CreatedAt: time.Now(),
	}
}

// execute : runs the contract code and updates the state
func (sc *SmartContract) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	//This is a placeholder for contract execution logic
	//In a real implementation, you would parse and execute
	//the contract code (eg. using a scripting language)

	//For now, we will just log the inputes and simulate execution
	fmt.Println("Executing contract with input:", input)

	//update state (for example purposes, we will store the last input in the state)
	sc.State["lastExecution"] = input

	//simulate contract returning updated state
	return sc.State, nil

}

// Validate checks if the smart contrct code is valid (placeholder)
func (sc *SmartContract) Validate() error {
	//This is a placeholder for contract validation logic
	//In a real implementation, you would parse and validate
	//the contract code (eg. using a scripting language)
	if sc.Code == "" {
		return errors.New("contract code is empty")
	}
	return nil
}
