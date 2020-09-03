package model

import (
	. "go-orm/db"
	"strings"

	"fmt"
)

type UserModel struct {
	Datasource string
	Table      string
	ID         int64

	Name string
}

func (m *UserModel) Exec(sql string) error {
	db := DBPool[m.Datasource]["w"]
	if _, err := db.Exec(sql); err != nil {
		fmt.Println("Execute sql failed:", err)
		return err
	}
	return nil
}

func (m *UserModel) CreateTable() error {
	sql := `CREATE TABLE user (
		id BIGINT AUTO_INCREMENT,

		name VARCHAR(255),
		PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	db := DBPool[m.Datasource]["w"]
	if _, err := db.Exec(sql); err != nil {
		fmt.Println("Create table failed:", err)
		return err
	}
	return nil
}

func (m *UserModel) New() *UserModel {
	n := UserModel{Datasource: "default", Table: "user"}
	return &n
}

func (m *UserModel) Find(id int64) (*UserModel, error) {
	sql := "SELECT * FROM user WHERE id = ?"
	db := DBPool[m.Datasource]["r"]
	row := db.QueryRow(sql, id)
	// todo
	if err := row.Scan(&m.ID, &m.Name); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *UserModel) Save() (*UserModel, error) {
	db := DBPool[m.Datasource]["w"]
	// Update
	if m.ID > 0 {
		// todo
	// Create
	} else {
		sql := "INSERT INTO user(name) VALUES(?)"
		result, err := db.Exec(sql, m.Name)
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

func (m *UserModel) Where(conds map[string]interface{}) []*UserModel {
	// todo
	ms := []*UserModel{}
	return ms
}

func (m *UserModel) Create(props map[string]interface{}) (*UserModel, error) {
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
	sql := fmt.Sprintf("INSERT INTO user(%s) VALUES(%s)", cstr, ph)

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

func (m *UserModel) Delete() error {
	// todo
	return nil
}

func (m *UserModel) Destroy(id int64) error {
	// sql := "DELETE FROM user WHERE id = ?"
	// todo
	return nil
}

func (m *UserModel) Update(props map[string]interface{}, conds map[string]interface{}) error {
	// todo
	return nil
}

var User = UserModel{Datasource: "default", Table: "user"}