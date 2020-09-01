package model

import (
	. "go-orm/db"

	"fmt"
	"time"
)

type EventModel struct {
	Datasource string
	Table      string

	Event string
	Created_at time.Time
}

func (m *EventModel) Exec(sql string) error {
	db := DBPool[m.Datasource]["w"]
	if _, err := db.Exec(sql); err != nil {
		fmt.Println("Execute sql failed:", err)
		return err
	}
	return nil
}

func (m *EventModel) CreateTable() error {
	sql := `CREATE TABLE event (
		id BIGINT AUTO_INCREMENT,

		event VARCHAR(255),
		created_at DATETIME,
		PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	db := DBPool[m.Datasource]["w"]
	if _, err := db.Exec(sql); err != nil {
		fmt.Println("Create table failed:", err)
		return err
	}
	return nil
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

var Event = EventModel{Datasource: "default", Table: "event"}