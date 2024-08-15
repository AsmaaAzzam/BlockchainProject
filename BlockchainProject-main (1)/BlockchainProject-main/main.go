package main

import (
	"eventticket/internal/handler"
	"eventticket/internal/memory"
	"eventticket/internal/model/ticket"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func mainChaincode() {
	// Initialize the repository
	repo := memory.NewTicketRepo()

	// Initialize the handler with the repository
	handler := handler.TicketHandler{
		RepoContract: *repo,
	}

	// Initialize the chaincode with the handler
	chaincode, err := contractapi.NewChaincode(&handler)
	if err != nil {
		fmt.Printf("error creating ticket chaincode: %v", err)
		return
	}

	// Start the chaincode
	if err := chaincode.Start(); err != nil {
		fmt.Printf("error starting ticket chaincode: %v", err)
	}
}

func main() {
	// Example usage of Ticket
	t := ticket.NewTicket("event1", "ticket1", "owner1")
	fmt.Println("Ticket:", t)

	// Uncomment to start the chaincode
	// mainChaincode()
}
