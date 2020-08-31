package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type Datasources struct {
	Datasources []Datasource `yaml:"datasources,flow"`
}

type Datasource struct {
	Name  string `yaml:"name"`
	Write string `yaml:"write"`
	Read  string `yaml:"read"`
}

var DBPool = make(map[string]map[string]*sql.DB)

func InitDBPool() {
	f, _ := ioutil.ReadFile("datasource.yml")
	dss := Datasources{}
	err := yaml.Unmarshal(f, &dss)
	if err != nil {
		fmt.Println("error:", err)
	}
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

		DBPool[ds.Name] = make(map[string]*sql.DB)
		DBPool[ds.Name]["w"] = wdb
		DBPool[ds.Name]["r"] = rdb
	}
}
