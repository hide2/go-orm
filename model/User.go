package model

import (
	. "go-orm/db"

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
	if err := row.Scan(&m.ID, &m.Name); err != nil {
		fmt.Printf("Scan failed, err:%v\n", err)
		return nil, err
	}
	return m, nil
}

func (m *UserModel) Save() (*UserModel, error) {
	if m.ID > 0 {
		fmt.Println("--Update")
	} else {
		fmt.Println("--Save")
	}
	return m, nil
}

func (m *UserModel) Where(conds map[string]interface{}) []*UserModel {
	// todo
	ms := []*UserModel{}
	return ms
}

func (m *UserModel) Create(props map[string]interface{}) (*UserModel, error) {
	// todo
	return m, nil
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