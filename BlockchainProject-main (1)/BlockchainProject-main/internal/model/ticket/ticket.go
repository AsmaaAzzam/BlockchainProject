package ticket

import "eventticket/internal/model"

type Ticket struct {
	EventID string `json:"eventId"`
	ID      string `json:"id"`
	Owner   string `json:"owner"`
}

// NewTicket creates a new Ticket instance
func NewTicket(EventID, ID, Owner string) *Ticket {
	return &Ticket{
		EventID: EventID,
		ID:      ID,
		Owner:   Owner,
	}
}

// GetID returns the ID of the Ticket, satisfying the model.Model interface
func (t *Ticket) GetID() string {
	return t.ID
}

// New creates a new Ticket instance as a model.Model
func (t *Ticket) New() model.Model {
	return t
}
