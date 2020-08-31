package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	. "go-orm/db"
	. "go-orm/model"
)

func main() {
	fmt.Println("DBPool", DBPool)
	User.CreateTable()
	User.Find(1)
}
