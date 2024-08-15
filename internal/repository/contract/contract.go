package contract

import (
	"encoding/json"
	"fmt"
	"eventticket/internal/model/event"
	"eventticket/internal/model/ticket"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RepoContract struct {
	contractapi.Contract
}

// Event-related methods

func (r *RepoContract) GetEvent(ctx contractapi.TransactionContextInterface, ID int) (*event.Event, error) {
	eventJSON, err := ctx.GetStub().GetState(fmt.Sprintf("%d", ID))
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: error %v", err)
	}
	if eventJSON == nil {
		return nil, fmt.Errorf("event %d does not exist", ID)
	}

	var event event.Event
	err = json.Unmarshal(eventJSON, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *RepoContract) GetAllEvents(ctx contractapi.TransactionContextInterface) ([]*event.Event, error) {
	eventIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: error %v", err)
	}
	defer eventIterator.Close()

	var events []*event.Event

	for eventIterator.HasNext() {
		result, err := eventIterator.Next()
		if err != nil {
			return nil, err
		}
		var event event.Event
		err = json.Unmarshal(result.Value, &event)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	return events, nil
}

func (r *RepoContract) PutEvent(ctx contractapi.TransactionContextInterface, event *event.Event) error {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(fmt.Sprintf("%d", event.ID), eventJSON)
}

// Ticket-related methods

func (r *RepoContract) GetTicket(ctx contractapi.TransactionContextInterface, ID string) (*ticket.Ticket, error) {
	ticketJSON, err := ctx.GetStub().GetState(ID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: error %v", err)
	}
	if ticketJSON == nil {
		return nil, fmt.Errorf("ticket %s does not exist", ID)
	}

	var ticket ticket.Ticket
	err = json.Unmarshal(ticketJSON, &ticket)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (r *RepoContract) GetAllTickets(ctx contractapi.TransactionContextInterface) ([]*ticket.Ticket, error) {
	ticketIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: error %v", err)
	}
	defer ticketIterator.Close()

	var tickets []*ticket.Ticket

	for ticketIterator.HasNext() {
		result, err := ticketIterator.Next()
		if err != nil {
			return nil, err
		}
		var ticket ticket.Ticket
		err = json.Unmarshal(result.Value, &ticket)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, &ticket)
	}

	return tickets, nil
}

func (r *RepoContract) PutTicket(ctx contractapi.TransactionContextInterface, ticket *ticket.Ticket) error {
	ticketJSON, err := json.Marshal(ticket)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(ticket.ID, ticketJSON)
}
