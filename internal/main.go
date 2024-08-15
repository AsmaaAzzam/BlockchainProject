package main

import (
	"fmt"
	"eventticket/internal/handler" // Replace with the actual path to your handler package

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func mainChaincode() {
	chaincode, err := contractapi.NewChaincode(new(handler.TicketHandler))

	if err != nil {
		fmt.Printf("error creating ticket chaincode: %v", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("error starting ticket chaincode: %v", err)
	}
}

func main() {
	mainChaincode()
	// handler := handler.TicketHandler{
	// 	RepoContract: *memory.New(), // Use appropriate initialization for your handler
	// }
}

