package model

import (
	. "go-orm/db"

	"fmt"
	"time"
)

type EventModel struct {
	datasource string
	table      string

	event string
	created_at time.Time
}

func (m *EventModel) CreateTable() {
	sql := `CREATE TABLE event (
		id BIGINT AUTO_INCREMENT,

		event VARCHAR(255),
		created_at DATETIME,
		PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	fmt.Println(sql)
	fmt.Println("datasource", DBPool[m.datasource])
	// todo
}

func (m *EventModel) New() *EventModel {
	// todo
	return m
}

func (m *EventModel) Find(id int64) (*EventModel, error) {
	sql := "SELECT * FROM event WHERE id = ?"
	fmt.Println(sql)
	return m, nil
}

func (m *EventModel) Where(conds map[string]interface{}) []*EventModel {
	// todo
	ms := []*EventModel{}
	return ms
}

func (m *EventModel) Save() (*EventModel, error) {
	// todo
	return m, nil
}

func (m *EventModel) Create(props map[string]interface{}) (*EventModel, error) {
	// todo
	return m, nil
}

func (m *EventModel) Delete() error {
	// todo
	return nil
}

func (m *EventModel) Destroy(id int64) error {
	// sql := "DELETE FROM event WHERE id = ?"
	// todo
	return nil
}

func (m *EventModel) Update(props map[string]interface{}, conds map[string]interface{}) error {
	// todo
	return nil
}

var Event = EventModel{datasource: "default", table: "event"}