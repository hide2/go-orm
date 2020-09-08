GO-ORM is a simple ORM framework based on code generator

# GO-ORM Features
- Auto create table
- Model & CRUD methods generator
- Transaction
- Pagination
- Connection Pool
- Write/Read Splitting
- Multi datasources
- Auto/Customized mapping of Model and datasource/table
- SQL log & Slow SQL log for profiling

# Install Go
```
sudo rm -fr /usr/local/go
Download & Install MacOS pkg from https://golang.org/dl/
export PATH=$PATH:/usr/local/go/bin
```

# Usage
Define Datasources in datasource.yml
``` yml
datasources:
  - name: default
    write: root:root@tcp(127.0.0.1:3306)/my_db_0?charset=utf8mb4&parseTime=True
    read: root:root@tcp(127.0.0.1:3306)/my_db_0?charset=utf8mb4&parseTime=True

  - name: ds_2
    write: root:root@tcp(127.0.0.1:3306)/my_db_0?charset=utf8mb4&parseTime=True
    read: root:root@tcp(127.0.0.1:3306)/my_db_0?charset=utf8mb4&parseTime=True

sql_log: false
slow_sql_log: 2
```
Define Models in model.yml
``` yml
models:
  - model: User
    name: string
    created_at: time.Time

  - model: Event
    name: string
    created_at: time.Time
```
Generate Model go files
```
go run gen.go
```
Which will generate Model files

-- Generate model/User.go
``` go
type UserModel struct {
	OdB string
	Lmt int
	Ofs int

	Datasource string
	Table      string
	Trx        *Tx
	ID         int64

	Name      string
	CreatedAt time.Time
}

...

func (m *UserModel) Find(id int64) (*UserModel, error) {
	...
}

...

var User = UserModel{Datasource: "default", Table: "user"}
```
You can use your User Model now:
``` go
package main

import (
	"fmt"
	"time"

	. "github.com/hide2/go-orm/model"
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
	u.CreatedAt = time.Now()
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
		props := map[string]interface{}{"name": "Dog", "created_at": time.Now()}
		u, _ = User.Create(props)
		fmt.Println("[Create]", u)
	}

	// WHERE
	conds := map[string]interface{}{"name": "Dog"}
	us, _ := User.Where(conds)
	for _, v := range us {
		fmt.Println("[Where]", *v)
	}

	// UPDATE
	props2 := map[string]interface{}{"name": "Cat"}
	conds2 := map[string]interface{}{"name": "Dog"}
	User.Update(props2, conds2)
	us2, _ := User.Where(props2)
	fmt.Println("[Update]", len(us2))

	// COUNT
	c, _ := User.CountAll()
	fmt.Println("[CountAll]", c)
	c, _ = User.Count(props2)
	fmt.Println("[Count]", c)

	// OrderBy&Limit&Paginate
	us3, _ := User.All()
	fmt.Println("[All]", len(us3))
	us4, _ := User.OrderBy("ID DESC").Offset(2).Limit(3).All()
	for _, v := range us4 {
		fmt.Println("[OrderBy]", *v)
	}
	us5, _ := User.OrderBy("ID ASC, Name ASC").Offset(0).Limit(5).Where(props2)
	for _, v := range us5 {
		fmt.Println("[Offset/Limit]", *v)
	}
	us6, _ := User.OrderBy("ID ASC, Name ASC").Page(1, 10).Where(props2)
	for _, v := range us6 {
		fmt.Println("[Page]", *v)
	}

	// Tx-Commit
	User.Begin()
	for i := 0; i < 10; i++ {
		props := map[string]interface{}{"name": fmt.Sprintf("%s%d", "Dog", i+1), "created_at": time.Now()}
		User.Create(props)
	}
	User.Destroy(25)
	User.Commit()
	u, e = User.Find(30)
	fmt.Println("[Find]", u, e)

	// Tx-Rollback
	User.Begin()
	for i := 0; i < 10; i++ {
		props := map[string]interface{}{"name": fmt.Sprintf("%s%d", "Dog", i+1), "created_at": time.Now()}
		User.Create(props)
	}
	User.Rollback()
	u, e = User.Find(40)
	fmt.Println("[Find]", u, e)
}
```