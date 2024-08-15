// Ticket represents a ticket
type Ticket struct {
    ID        string `json:"id"`
    EventID   string `json:"eventId"`
    Status    string `json:"status"`
}
// NewTicket creates a new ticket
func NewTicket(id, eventID, status string) *Ticket {
    return &Ticket{
        ID:        id,
        EventID:   eventID,
        Status:    status,
    }
}
