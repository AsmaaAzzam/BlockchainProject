package handler

import (
	"fmt"
	"eventticket/internal/model/event"
	"eventticket/internal/model/ticket"
	"eventticket/internal/repository/contract"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type EventTicketHandler struct {
	contract.RepoContract
}

// Event-related methods

func (h *EventTicketHandler) GetEvent(ctx contractapi.TransactionContextInterface, ID int) (*event.Event, error) {
	return h.RepoContract.GetEvent(ctx, ID)
}

func (h *EventTicketHandler) GetAllEvents(ctx contractapi.TransactionContextInterface) ([]*event.Event, error) {
	return h.RepoContract.GetAllEvents(ctx)
}

func (h *EventTicketHandler) CreateEvent(ctx contractapi.TransactionContextInterface, ID int, name, date string, ticketsLeft, price int) error {
	_, err := h.GetEvent(ctx, ID)
	if err == nil {
		return fmt.Errorf("event with ID %d already exists", ID)
	}
	newEvent := event.NewEvent(ID, name, date, ticketsLeft, price)
	return h.RepoContract.PutEvent(ctx, newEvent)
}

// Ticket-related methods

func (h *EventTicketHandler) GetTicket(ctx contractapi.TransactionContextInterface, ID string) (*ticket.Ticket, error) {
	return h.RepoContract.GetTicket(ctx, ID)
}

func (h *EventTicketHandler) GetAllTickets(ctx contractapi.TransactionContextInterface) ([]*ticket.Ticket, error) {
	return h.RepoContract.GetAllTickets(ctx)
}

func (h *EventTicketHandler) PurchaseTicket(ctx contractapi.TransactionContextInterface, ticketID, studentID string, eventID int) error {
	// Check if the event exists and has available tickets
	event, err := h.GetEvent(ctx, eventID)
	if err != nil {
		return fmt.Errorf("event with ID %d does not exist", eventID)
	}
	if event.TicketsLeft <= 0 {
		return fmt.Errorf("no tickets available for event %d", eventID)
	}

	// Create a new ticket
	newTicket := ticket.NewTicket(ticketID, studentID, fmt.Sprintf("%d", eventID), "purchased")
	event.TicketsLeft--

	// Update the event and save the ticket
	err = h.RepoContract.PutEvent(ctx, event)
	if err != nil {
		return err
	}
	return h.RepoContract.PutTicket(ctx, newTicket)
}

func (h *EventTicketHandler) UpdateTicketStatus(ctx contractapi.TransactionContextInterface, ticketID, newStatus string) error {
	// Check if the ticket exists
	ticket, err := h.GetTicket(ctx, ticketID)
	if err != nil {
		return err
	}
	ticket.Status = newStatus
	return h.RepoContract.PutTicket(ctx, ticket)
}
