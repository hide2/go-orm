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
```
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
```
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
```
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