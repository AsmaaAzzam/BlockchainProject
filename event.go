package events

type Event struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Date        string `json:"date"`
    TicketsLeft int    `json:"ticketsLeft"`
    Price       int    `json:"price"`
}




func NewEvent(id int, name, date string, ticketsLeft, price int) *Event {
    return &Event{
        ID:          id,
        Name:        name,
        Date:        date,
        TicketsLeft: ticketsLeft,
        Price:       price,
    }
}
