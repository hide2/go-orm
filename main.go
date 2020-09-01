package main

import (
	"fmt"
	. "go-orm/model"
)

func main() {
	// Exec SQL
	User.Exec("DROP TABLE IF EXISTS user")

	// Create Table
	User.CreateTable()

	// C
	u := User.New()
	u.Name = "John"
	u.Save()
	fmt.Println("[Save]", u)

	// R
	u, _ = User.Find(1)
	fmt.Println("[Find]", u)

	// U
	u.Name = "Calvin"
	u.Save()
	fmt.Println("[Update]", u)

	// D
	u.Delete()
	User.Destroy(1)

	// Create
	props := map[string]interface{}{"name": "Dog"}
	u, _ = User.Create(props)
	fmt.Println("[Create]", u)

	// WHERE
	conds := map[string]interface{}{"name": "Cat"}
	us := User.Where(conds)
	fmt.Println("[Where]", us)

	// UPDATE
	props2 := map[string]interface{}{"name": "Cat"}
	conds2 := map[string]interface{}{"name": "Dog"}
	User.Update(props2, conds2)
}
