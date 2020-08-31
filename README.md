# GO-ORM Features
- Auto create table
- Model & CRUD methods generator
- Pagination
- Connection Pool
- Write/Read Splitting
- Multi datasources
- Auto/Customized mapping of Model and datasource/table
- SQL log & Slow SQL log for profiling

# Generator
Define Models in model.yml
``` yml
models:
  - model: User
    name: string

  - model: Event
    event: string
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
  datasource string
  table      string
  name string
}

...

func (m *UserModel) Find(id int64) (*UserModel, error) {
	...
}

...

var User = UserModel{datasource: "default", table: "user"}
```
-- Generate model/Event.go
``` go
type EventModel struct {
  datasource string
  table      string
  event string
  created_at time.Time
}

...

func (m *EventModel) Find(id int64) (*EventModel, error) {
	...
}

...

var Event = EventModel{datasource: "default", table: "event"}
```
You can use your User/Event Model now:
``` go
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

// UPDATE
props := map[string]interface{}{"name": "Cat"}
conds := map[string]interface{}{"name": "Dog"}
User.Update(props, conds)

// WHERE
conds := map[string]interface{}{"name": "Cat"}
us := User.Where(conds)
fmt.Println("[Where]", us))
```