# GO-ORM Features
- Auto create table
- Model & CRUD methods generator
- Pagination
- Connection Pool
- Write/Read Splitting
- Multi datasources
- Auto/Customized mapping of Model and datasource/table
- SQL log & Slow SQL log for profiling

# Usage
Define Datasources in datasource.yml
``` yml
datasources:
  - name: default
    write: root:root@tcp(127.0.0.1:3306)/my_db_0
    read: root:root@tcp(127.0.0.1:3306)/my_db_0

  - name: ds_2
    write: root:root@tcp(127.0.0.1:3306)/my_db_0
    read: root:root@tcp(127.0.0.1:3306)/my_db_0

sql_log: true
slow_sql_log: false
```
Define Models in model.yml
``` yml
models:
  - model: User
    name: string

  - model: Event
    name: string
    created_at: time.Time
```
Generate Model go files
```
go run generator/generator.go -file model.yml
```
Which will generate Model files

-- Generate model/User.go
``` go
type UserModel struct {
	Datasource string
	Table      string
	ID         int64

	Name string
}

...

func (m *UserModel) Find(id int64) (*UserModel, error) {
	...
}

...

var User = UserModel{Datasource: "default", Table: "user"}
```
-- Generate model/Event.go
``` go
type EventModel struct {
	Datasource string
	Table      string
	ID         int64

	Name string
	CreatedAt time.Time
}

...

func (m *EventModel) Find(id int64) (*EventModel, error) {
	...
}

...

var Event = EventModel{Datasource: "default", Table: "event"}
```
You can use your User/Event Model now:
``` go
// Exec SQL
User.Exec("DROP TABLE IF EXISTS user")

// Create Table
User.CreateTable()

// C
u := User.New()
u.name = "John"
u.Save()
fmt.Println("[Save]", u)

// R
u = User.Find(1)
fmt.Println("[Find]", u)

// U
u.name = "Calvin"
u.Save()
fmt.Println("[Update]", u)

// D
u.Delete()
User.Destroy(1)

// Create
props := map[string]interface{}{"name": "Dog"}
u = User.Create(props)
fmt.Println("[Create]", u)

// WHERE
conds := map[string]interface{}{"name": "Cat"}
us := User.Where(conds)
fmt.Println("[Where]", us)

// UPDATE
props2 := map[string]interface{}{"name": "Cat"}
conds2 := map[string]interface{}{"name": "Dog"}
User.Update(props2, conds2)
```