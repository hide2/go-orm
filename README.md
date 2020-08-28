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
    id: int64
    name: string

  - model: Event
    id: int64
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

	id int64

	name string

}

...

func (m *UserModel) find(id int64) (*UserModel, error) {
	...
}

...

var User = UserModel{datasource: "default", table: "user"}
```
-- Generate model/Event.go
``` go
type EventModel struct {

	id int64

	event string

	created_at time.Time

}

...

func (m *EventModel) find(id int64) (*EventModel, error) {
	...
}

...

var Event = EventModel{datasource: "default", table: "event"}
```
You can use your User/Event Model now:
``` go
// Create Table
User.createtable()

// C
u := User.new()
u.name = "John"
u.save()
fmt.Println("[Save]", u)

// R
u = User.find(1)
fmt.Println("[Find]", u)

// U
u.name = "Calvin"
u.save()
fmt.Println("[Update]", u)

// D
u.delete()
User.destroy(1)

// Create
props := map[string]interface{}{"name": "Dog"}
u = User.create(props)
fmt.Println("[Create]", u)

// UPDATE
props := map[string]interface{}{"name": "Cat"}
conds := map[string]interface{}{"name": "Dog"}
User.update(props, conds)

// WHERE
conds := map[string]interface{}{"name": "Cat"}
us := User.where(conds)
fmt.Println("[Where]", us))
```