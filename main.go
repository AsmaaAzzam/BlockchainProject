package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Smart Contract
type SmartContract struct {
	contractapi.Contract
}

// Event struct
type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	TicketsLeft int    `json:"ticketsLeft"`
	Price       int    `json:"price"`
}

// Ticket struct
type Ticket struct {
	EventID int    `json:"eventId"`
	ID      int    `json:"id"`
	Owner   string `json:"owner"`
}

// Database simulation (using maps)
var eventsMap = make(map[int]*Event)
var ticketsMap = make(map[int]*Ticket)
var ticketCounter = 1

// Create a new event
func (s *SmartContract) CreateEvent(ctx contractapi.TransactionContextInterface, name string, date string, totalTickets int, price int) error {
	newEvent := &Event{
		ID:          len(eventsMap) + 1,
		Name:        name,
		Date:        date,
		TicketsLeft: totalTickets,
		Price:       price,
	}
	eventsMap[newEvent.ID] = newEvent

	// Simulate writing to blockchain (in real-world, you'd use ctx.GetStub().PutState())
	eventJSON, _ := json.Marshal(newEvent)
	ctx.GetStub().PutState(strconv.Itoa(newEvent.ID), eventJSON)

	return nil
}

// Purchase a ticket for an event
func (s *SmartContract) PurchaseTicket(ctx contractapi.TransactionContextInterface, eventID int, owner string) error {
	event, exists := eventsMap[eventID]
	if !exists {
		return fmt.Errorf("Event with ID %d not found", eventID)
	}

	if event.TicketsLeft > 0 {
		event.TicketsLeft--

		// Simulate writing to blockchain (in real-world, you'd use ctx.GetStub().PutState())
		eventJSON, _ := json.Marshal(event)
		ctx.GetStub().PutState(strconv.Itoa(event.ID), eventJSON)

		newTicket := &Ticket{
			EventID: eventID,
			ID:      ticketCounter,
			Owner:   owner,
		}
		ticketsMap[ticketCounter] = newTicket
		ticketCounter++

		// Simulate writing to blockchain (in real-world, you'd use ctx.GetStub().PutState())
		ticketJSON, _ := json.Marshal(newTicket)
		ctx.GetStub().PutState(strconv.Itoa(newTicket.ID), ticketJSON)

		return nil
	} else {
		return fmt.Errorf("No tickets left for event %d", eventID)
	}
}

// Get event details by ID
func (s *SmartContract) GetEvent(ctx contractapi.TransactionContextInterface, eventID int) (*Event, error) {
	eventJSON, err := ctx.GetStub().GetState(strconv.Itoa(eventID))
	if err != nil {
		return nil, fmt.Errorf("Failed to get event with ID %d: %w", eventID, err)
	}
	if eventJSON == nil {
		return nil, fmt.Errorf("Event with ID %d not found", eventID)
	}

	var event Event
	err = json.Unmarshal(eventJSON, &event)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal event JSON: %w", err)
	}

	return &event, nil
}

// Get ticket details by ID
func (s *SmartContract) GetTicket(ctx contractapi.TransactionContextInterface, ticketID int) (*Ticket, error) {
	ticketJSON, err := ctx.GetStub().GetState(strconv.Itoa(ticketID))
	if err != nil {
		return nil, fmt.Errorf("Failed to get ticket with ID %d: %w", ticketID, err)
	}
	if ticketJSON == nil {
		return nil, fmt.Errorf("Ticket with ID %d not found", ticketID)
	}

	var ticket Ticket
	err = json.Unmarshal(ticketJSON, &ticket)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal ticket JSON: %w", err)
	}

	return &ticket, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %s", err.Error())
		return
	}
}
