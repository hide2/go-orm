package main

import (
	"fmt"
	. "go-orm/model"
)

func main() {
	// Exec SQL
	fmt.Println("[Drop Table user]")
	User.Exec("DROP TABLE IF EXISTS user")

	// Create Table
	fmt.Println("[Create Table user]")
	User.CreateTable()

	// C
	u := User.New()
	fmt.Println("[New]", u)
	u.Name = "John"
	u.Save()
	fmt.Println("[Save]", u)

	// R
	u, e := User.Find(123)
	fmt.Println("[Find(123)]", u, e)

	u, _ = User.Find(1)
	fmt.Println("[Find(1)]", u)

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
	u, _ = User.Create(props)
	u, _ = User.Create(props)
	fmt.Println("[Create]", u)

	// WHERE
	conds := map[string]interface{}{"name": "Dog"}
	us := User.Where(conds)
	fmt.Println("[Where]", us)

	// UPDATE
	props2 := map[string]interface{}{"name": "Cat"}
	conds2 := map[string]interface{}{"name": "Dog"}
	User.Update(props2, conds2)
	us2 := User.Where(props2)
	fmt.Println("[Update]", us2)
}
