package models

import "github.com/nitesh-mhatre/go-rest-api/db"


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

func (e Event) Save() error{
	quary := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt , err := db.DB.Prepare(quary)
	
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = int(id)

	return nil

}

func GetAllEvents() ([]Event, error) {
	quary := `SELECT id, name, description, location, date_time, user_id FROM events`
	rows, err := db.DB.Query(quary)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var events []Event	
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}