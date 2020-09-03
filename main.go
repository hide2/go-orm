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

	u, _ = User.Find(1)
	fmt.Println("[Find(1)]", u)

	// D
	u.Delete()
	u, _ = User.Find(1)
	fmt.Println("[After Delete Find(1)]", u)

	// Exec SQL
	fmt.Println("[Drop Table user]")
	User.Exec("DROP TABLE IF EXISTS user")

	// Create Table
	fmt.Println("[Create Table user]")
	User.CreateTable()

	// Create
	for i := 0; i < 20; i++ {
		props := map[string]interface{}{"name": fmt.Sprintf("%s%d", "Dog", i+1)}
		u, _ = User.Create(props)
		fmt.Println("[Create]", u)
	}

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

	// COUNT
	// c := User.CountAll()
	// fmt.Println("[CountAll]", c)
	// c = User.Count(props2)
	// fmt.Println("[Count]", c)

	// OrderBy&Limit&Paginate
	// us3 := User.All()
	// us4 := User.OrderBy("ID DESC").All()
	// us5 := User.OrderBy("ID DESC, Name DESC").Where(props2)
	// us6 := User.OrderBy("ID DESC, Name DESC").Where(props2)
	// us7 := User.OrderBy("ID DESC, Name DESC").Offset(0).Limit(10).Where(props2)
	// us8 := User.OrderBy("ID DESC, Name DESC").Page(1, 10).Where(props2)

	// Tx-Commit
	// User.Begin()
	// for i := 0; i < 10; i++ {
	// 	props := map[string]interface{}{"name": fmt.Sprintf("%s%d", "Dog", i+1)}
	// 	User.Create(props)
	// }
	// User.Destroy(25)
	// User.Commit()
	// u, e = User.Find(30)
	// fmt.Println("[Find]", u, e)

	// // Tx-Rollback
	// User.Begin()
	// for i := 0; i < 10; i++ {
	// 	props := map[string]interface{}{"name": fmt.Sprintf("%s%d", "Dog", i+1)}
	// 	User.Create(props)
	// }
	// User.Rollback()
	// u, e = User.Find(40)
	// fmt.Println("[Find]", u, e)
}
