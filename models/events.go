package models


type Event struct {
	ID int 
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime string `binding:"required"`
	UserID int
}

var events []Event = []Event{

}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}