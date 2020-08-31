package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	. "go-orm/db"
	. "go-orm/model"
)

func main() {
	InitDBPool()
	fmt.Println("DBPoolDBs", DBPool)
	User.CreateTable()
	User.Find(1)
}
