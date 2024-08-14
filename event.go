package events

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	TicketsLeft int    `json:"ticketsLeft"`
	Price       int    `json:"price"`
}