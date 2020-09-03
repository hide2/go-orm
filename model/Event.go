package model

import (
	. "go-orm/db"
	"strings"

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
	// todo
	if err := row.Scan(&m.ID, &m.Name); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *EventModel) Save() (*EventModel, error) {
	db := DBPool[m.Datasource]["w"]
	// Update
	if m.ID > 0 {
		// todo
	// Create
	} else {
		sql := "INSERT INTO event(name,created_at) VALUES(?,?)"
		result, err := db.Exec(sql, m.Name, m.CreatedAt)
		if err != nil {
			fmt.Printf("Insert data failed, err:%v", err)
			return nil, err
		}
		lastInsertID, err := result.LastInsertId() //获取插入数据的自增ID
		if err != nil {
			fmt.Printf("Get insert id failed, err:%v", err)
			return nil, err
		}
		m.ID = lastInsertID
	}
	return m, nil
}

func (m *EventModel) Where(conds map[string]interface{}) []*EventModel {
	// todo
	ms := []*EventModel{}
	return ms
}

func (m *EventModel) Create(props map[string]interface{}) (*EventModel, error) {
	db := DBPool[m.Datasource]["w"]

	keys := make([]string, 0)
	values := make([]interface{}, 0)
	for k, v := range props {
		keys = append(keys, k)
		values = append(values, v)
	}
	cstr := strings.Join(keys, ",")
	phs := make([]string, 0)
	for i := 0; i < len(keys); i++ {
		phs = append(phs, "?")
	}
	ph := strings.Join(phs, ",")
	sql := fmt.Sprintf("INSERT INTO event(%s) VALUES(%s)", cstr, ph)

	result, err := db.Exec(sql, values...)
	if err != nil {
		fmt.Printf("Insert data failed, err:%v", err)
		return nil, err
	}
	lastInsertID, err := result.LastInsertId() //获取插入数据的自增ID
	if err != nil {
		fmt.Printf("Get insert id failed, err:%v", err)
		return nil, err
	}
	return m.Find(lastInsertID)
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