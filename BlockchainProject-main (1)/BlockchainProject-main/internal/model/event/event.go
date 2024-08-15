package event

import (
	"eventticket/internal/model"
	"fmt"
)

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	TicketsLeft int    `json:"ticketsLeft"`
	Price       int    `json:"price"`
}

func NewEvent(ID int, name string, date string, ticketsLeft, price int) *Event {
	return &Event{
		ID:          ID,
		Name:        name,
		Date:        date,
		TicketsLeft: ticketsLeft,
		Price:       price,
	}
}

// Implement GetID method to satisfy the model.Model interface
func (e *Event) GetID() string {
	return fmt.Sprintf("%d", e.ID)
}

// New returns the Event as a model.Model interface
func (e *Event) New() model.Model {
	return e
}
