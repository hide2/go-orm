package model

import (
	. "go-orm/db"

	"fmt"
	"time"
)

type EventModel struct {
	Datasource string
	Table      string
	ID         int64

	Name string
	CreatedAt time.Time
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

		name VARCHAR(255),
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
	n := EventModel{Datasource: "default", Table: "event"}
	return &n
}

func (m *EventModel) Find(id int64) (*EventModel, error) {
	sql := "SELECT * FROM event WHERE id = ?"
	db := DBPool[m.Datasource]["r"]
	row := db.QueryRow(sql, id)
	if err := row.Scan(&m.ID, &m.Name); err != nil {
		fmt.Printf("Scan failed, err:%v\n", err)
		return nil, err
	}
	return m, nil
}

func (m *EventModel) Save() (*EventModel, error) {
	if m.ID > 0 {
		fmt.Println("--Update")
	} else {
		fmt.Println("--Save")
	}
	return m, nil
}

func (m *EventModel) Where(conds map[string]interface{}) []*EventModel {
	// todo
	ms := []*EventModel{}
	return ms
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