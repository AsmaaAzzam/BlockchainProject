package memory

import (
	"eventticket/internal/model/event"
	"eventticket/internal/model/ticket"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type EventRepoMemory map[int]*event.Event
type TicketRepoMemory map[string]*ticket.Ticket

func NewEventRepo() *EventRepoMemory {
	er := make(EventRepoMemory)
	return &er
}

func NewTicketRepo() *TicketRepoMemory {
	tr := make(TicketRepoMemory)
	return &tr
}

func (erm *EventRepoMemory) GetEvent(_ contractapi.TransactionContextInterface, ID int) (*event.Event, error) {
	event, existing := (*erm)[ID]
	if !existing {
		return nil, fmt.Errorf("no event with ID: %d", ID)
	}
	return event, nil
}

func (erm *EventRepoMemory) GetAllEvents(_ contractapi.TransactionContextInterface) ([]*event.Event, error) {
	var events []*event.Event
	for _, v := range *erm {
		events = append(events, v)
	}
	if len(events) == 0 {
		return nil, fmt.Errorf("no events found")
	}
	return events, nil
}

func (erm *EventRepoMemory) PutEvent(_ contractapi.TransactionContextInterface, event *event.Event) error {
	(*erm)[event.ID] = event
	return nil
}

func (trm *TicketRepoMemory) GetTicket(_ contractapi.TransactionContextInterface, ID string) (*ticket.Ticket, error) {
	ticket, existing := (*trm)[ID]
	if !existing {
		return nil, fmt.Errorf("no ticket with ID: %s", ID)
	}
	return ticket, nil
}

func (trm *TicketRepoMemory) GetAllTickets(_ contractapi.TransactionContextInterface) ([]*ticket.Ticket, error) {
	var tickets []*ticket.Ticket
	for _, v := range *trm {
		tickets = append(tickets, v)
	}
	if len(tickets) == 0 {
		return nil, fmt.Errorf("no tickets found")
	}
	return tickets, nil
}

func (trm *TicketRepoMemory) PutTicket(_ contractapi.TransactionContextInterface, ticket *ticket.Ticket) error {
	(*trm)[ticket.ID] = ticket
	return nil
}
