package repository

import (
	"event_ticket/internal/model/event"
	"event_ticket/internal/model/ticket"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type EventTicketRepository interface {
	// Event-related methods
	GetEvent(contractapi.TransactionContextInterface, int) (*event.Event, error)
	GetAllEvents(contractapi.TransactionContextInterface) ([]*event.Event, error)
	PutEvent(contractapi.TransactionContextInterface, *event.Event) error

	// Ticket-related methods
	GetTicket(contractapi.TransactionContextInterface, string) (*ticket.Ticket, error)
	GetAllTickets(contractapi.TransactionContextInterface) ([]*ticket.Ticket, error)
	PutTicket(contractapi.TransactionContextInterface, *ticket.Ticket) error
}
