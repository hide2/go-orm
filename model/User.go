package model

import (
	. "go-orm/db"

	"fmt"
)

type UserModel struct {
	Datasource string
	Table      string

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
	// todo
	return m
}

func (m *UserModel) Find(id int64) (*UserModel, error) {
	sql := "SELECT * FROM user WHERE id = ?"
	fmt.Println(sql)
	return m, nil
}

func (m *UserModel) Where(conds map[string]interface{}) []*UserModel {
	// todo
	ms := []*UserModel{}
	return ms
}

func (m *UserModel) Save() (*UserModel, error) {
	// todo
	return m, nil
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