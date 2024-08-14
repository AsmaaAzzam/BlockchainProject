package tickets

type Ticket struct {
	EventID int    `json:"eventId"`
	ID      int    `json:"id"`
	Owner   string `json:"owner"`
}