package Event

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	TicketsLeft int    `json:"ticketsLeft"`
	Price       int    `json:"price"`
}

// new editing
func New(ID, name, date, ticketsLeft, price string) *Event {
	return &Event{
		ID:          ID,
		Name:        name,
		Date:        date,
		TicketsLeft: ticketsLeft,
		Price:       price,
	}
}
