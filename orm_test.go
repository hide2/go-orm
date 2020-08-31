package orm

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type UserModel struct {
	datasource string
	table      string

	id   int64
	name string
}

func (u *UserModel) CreateTable() {
	// todo
}

func (u *UserModel) find(id int64) *UserModel {
	u.id = id
	u.name = "Test"
	return u
}

func (u *UserModel) new() *UserModel {
	return u
}

func (u *UserModel) save() *UserModel {
	return u
}

func (u *UserModel) create(attrs map[string]interface{}) *UserModel {
	if _, ok := attrs["id"].(int); ok {
		u.id = int64(attrs["id"].(int))
	}
	if _, ok := attrs["name"].(string); ok {
		u.name = attrs["name"].(string)
	}
	return u
}

func (u *UserModel) delete() {
}

func (u *UserModel) destroy(id int64) {
}

func (u *UserModel) count(conds map[string]interface{}) int {
	return 0
}

func (u *UserModel) where(conds map[string]interface{}) []*UserModel {
	us := []*UserModel{}
	return us
}

func (u *UserModel) update(pros map[string]interface{}, conds map[string]interface{}) []*UserModel {
	us := []*UserModel{}
	return us
}

var User = UserModel{datasource: "default", table: "users"}

type Datasources struct {
	Datasources []Datasource `yaml:"datasources,flow"`
}

type Datasource struct {
	Name  string `yaml:"name"`
	Write string `yaml:"write"`
	Read  string `yaml:"read"`
}

type ModelsConfig struct {
	Models []ModelConfig `yaml:"models,flow"`
}

type ModelConfig map[string]string

func TestORM(t *testing.T) {
	f, _ := ioutil.ReadFile("datasource.yml")
	dss := Datasources{}
	err := yaml.Unmarshal(f, &dss)
	if err != nil {
		fmt.Println("error:", err)
	}
	DBs := make(map[string]map[string]*sql.DB)
	for _, ds := range dss.Datasources {
		wdb, err := sql.Open("mysql", ds.Write)
		if err != nil {
			fmt.Println("Connection to mysql failed:", err)
			return
		}
		wdb.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超时的连接就close
		wdb.SetMaxOpenConns(100)                  //设置最大连接数
		rdb, err := sql.Open("mysql", ds.Read)
		if err != nil {
			fmt.Println("Connection to mysql failed:", err)
			return
		}
		rdb.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超时的连接就close
		rdb.SetMaxOpenConns(100)                  //设置最大连接数

		DBs[ds.Name] = make(map[string]*sql.DB)
		DBs[ds.Name]["w"] = wdb
		DBs[ds.Name]["r"] = rdb
	}
	fmt.Println("DBs", DBs)

	// R
	u := User.find(123)
	fmt.Println("[Find]", u)

	// C
	u2 := User.new()
	u2.id = 111
	u2.name = "John"
	u2.save()
	fmt.Println("[Save]", u2)

	// U
	u3 := User.find(111)
	u3.name = "MAYUN"
	u3.save()
	fmt.Println("[Update]", u3)

	// D
	User.destroy(123)

	// Create
	attrs := map[string]interface{}{"id": 9223372036854775807, "name": "888"}
	u4 := User.create(attrs)
	fmt.Println("[Create]", u4)

	// WHERE
	conds := map[string]interface{}{"id": 123, "name": "Test"}
	us := User.where(conds)
	fmt.Println("[Where]", us)
}
