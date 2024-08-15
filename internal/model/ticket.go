package ticket

type Ticket struct {
	EventID int    `json:"eventId"`
	ID      int    `json:"id"`
	Owner   string `json:"owner"`
}

// new editing

func New(EventID, ID, Owner string) *tickets {
	return &tickets{
		EventID: eventID,
		ID:      ID,
		Owner:   Owner,
	}
}
