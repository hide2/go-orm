package model

import (
	. "database/sql"
	. "go-orm/db"
	. "go-orm/lib"
	"strings"
	"time"

	"fmt"
)

type UserModel struct {
	Datasource string
	Table      string
	Trx        *Tx
	ID         int64

	Name string
	CreatedAt time.Time
}

func (m *UserModel) Begin() (*Tx, error) {
	db := DBPool[m.Datasource]["w"]
	sql := "BEGIN"
	if GoOrmSqlLog {
		fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql)
	}
	tx, err := db.Begin()
	m.Trx = tx
	return tx, err
}

func (m *UserModel) Commit() error {
	if m.Trx != nil {
		sql := "COMMIT"
		if GoOrmSqlLog {
			fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql)
		}
		return m.Trx.Commit()
	}
	m.Trx = nil
	return nil
}

func (m *UserModel) Rollback() error {
	if m.Trx != nil {
		sql := "ROLLBACK"
		if GoOrmSqlLog {
			fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql)
		}
		return m.Trx.Rollback()
	}
	m.Trx = nil
	return nil
}

func (m *UserModel) Exec(sql string) error {
	db := DBPool[m.Datasource]["w"]
	if GoOrmSqlLog {
		fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql)
	}
	if _, err := db.Exec(sql); err != nil {
		fmt.Println("Execute sql failed:", err)
		return err
	}
	return nil
}

func (m *UserModel) CreateTable() error {
	db := DBPool[m.Datasource]["w"]
	sql := `CREATE TABLE user (
		id BIGINT AUTO_INCREMENT,

		name VARCHAR(255),
		created_at DATETIME,
		PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	if GoOrmSqlLog {
		fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql)
	}
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
	db := DBPool[m.Datasource]["r"]
	sql := "SELECT * FROM user WHERE id = ?"
	if GoOrmSqlLog {
		fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql, id)
	}
	row := db.QueryRow(sql, id)
	if err := row.Scan(&m.ID, &m.Name, &m.CreatedAt); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *UserModel) Save() (*UserModel, error) {
	db := DBPool[m.Datasource]["w"]
	// Update
	if m.ID > 0 {
		props := StructToMap(*m)
		conds := map[string]interface{}{"id": m.ID}
		uprops := make(map[string]interface{})
		for k, v := range props {
			if k != "Datasource" && k != "Table" && k != "Trx" && k != "ID" {
				uprops[Underscore(k)] = v
			}
		}
		return m, m.Update(uprops, conds)
	// Create
	} else {
		sql := "INSERT INTO user(name,created_at) VALUES(?,?)"
		if GoOrmSqlLog {
			fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql, m.Name, m.CreatedAt)
		}
		result, err := db.Exec(sql, m.Name, m.CreatedAt)
		if err != nil {
			fmt.Printf("Insert data failed, err:%v\n", err)
			return nil, err
		}
		lastInsertID, err := result.LastInsertId() //获取插入数据的自增ID
		if err != nil {
			fmt.Printf("Get insert id failed, err:%v\n", err)
			return nil, err
		}
		m.ID = lastInsertID
	}
	return m, nil
}

func (m *UserModel) Where(conds map[string]interface{}) ([]*UserModel, error) {
	db := DBPool[m.Datasource]["r"]
	wherestr := make([]string, 0)
	cvs := make([]interface{}, 0)
	for k, v := range conds {
		wherestr = append(wherestr, k + "=?")
		cvs = append(cvs, v)
	}
	sql := fmt.Sprintf("SELECT * FROM user WHERE %s", strings.Join(wherestr, " AND "))
	if GoOrmSqlLog {
		fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql, cvs)
	}
	rows, err := db.Query(sql, cvs...)
	defer func() {
		if rows != nil {
			rows.Close() //关闭掉未scan的sql连接
		}
	}()
	if err != nil {
		fmt.Printf("Query data failed, err:%v\n", err)
		return nil, err
	}
	ms := make([]*UserModel, 0)
	for rows.Next() {
		m = new(UserModel)
		err = rows.Scan(&m.ID, &m.Name, &m.CreatedAt) //不scan会导致连接不释放
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil
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

	if GoOrmSqlLog {
		fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql, values)
	}
	var result Result
	var err error
	if m.Trx != nil {
		result, err = m.Trx.Exec(sql, values...)
	} else {
		result, err = db.Exec(sql, values...)
	}
	if err != nil {
		fmt.Printf("Insert data failed, err:%v\n", err)
		return nil, err
	}
	lastInsertID, err := result.LastInsertId() //获取插入数据的自增ID
	if err != nil {
		fmt.Printf("Get insert id failed, err:%v\n", err)
		return nil, err
	}
	return m.Find(lastInsertID)
}

func (m *UserModel) Delete() error {
	return m.Destroy(m.ID)
}

func (m *UserModel) Destroy(id int64) error {
	db := DBPool[m.Datasource]["w"]
	sql := "DELETE FROM user WHERE id = ?"
	if GoOrmSqlLog {
		fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql, id)
	}
	var err error
	if m.Trx != nil {
		_, err = m.Trx.Exec(sql, id)
	} else {
		_, err = db.Exec(sql, id)
	}
	if err != nil {
		fmt.Printf("Delete data failed, err:%v\n", err)
		return err
	}
	m.ID = 0
	return nil
}

func (m *UserModel) Update(props map[string]interface{}, conds map[string]interface{}) error {
	db := DBPool[m.Datasource]["w"]
	setstr := make([]string, 0)
	wherestr := make([]string, 0)
	cvs := make([]interface{}, 0)
	for k, v := range props {
		setstr = append(setstr, k + "=?")
		cvs = append(cvs, v)
	}
	for k, v := range conds {
		wherestr = append(wherestr, k + "=?")
		cvs = append(cvs, v)
	}
	sql := fmt.Sprintf("UPDATE user SET %s WHERE %s", strings.Join(setstr, ", "), strings.Join(wherestr, " AND "))
	if GoOrmSqlLog {
		fmt.Println("["+time.Now().Format("2006-01-02 15:04:05")+"][SQL]", sql, cvs)
	}
	var err error
	if m.Trx != nil {
		_, err = m.Trx.Exec(sql, cvs...)
	} else {
		_, err = db.Exec(sql, cvs...)
	}
	if err != nil {
		fmt.Printf("Update data failed, err:%v\n", err)
		return err
	}
	return nil
}

func (m *UserModel) CountAll() int64 {
	// todo
	return 0
}

func (m *UserModel) Count(conds map[string]interface{}) int64 {
	// todo
	return 0
}

func (m *UserModel) All() ([]*UserModel, error) {
	// todo
	return nil, nil
}

func (m *UserModel) OrderBy(o string) *UserModel {
	// todo
	return nil
}

func (m *UserModel) Offset(o int64) *UserModel {
	// todo
	return nil
}

func (m *UserModel) Limit(l int64) *UserModel {
	// todo
	return nil
}

func (m *UserModel) Page(page int64, size int64) ([]*UserModel, error) {
	// todo
	return nil, nil
}

var User = UserModel{Datasource: "default", Table: "user"}