type EventModel struct {
	datasource string
	table      string

	id int64

	event string

	created_at time.Time

}

func (m *EventModel) _createtable() {
	// todo
	return sql
}

func (m *EventModel) _find(id int64) (string, int64)  {
	sql := "SELECT * FROM event WHERE id = ?"
	return sql, id
}

func (m *EventModel) _where(conds map[string]interface{}) (string, ...interface{}) {
	// todo
	return sql, args
}

func (m *EventModel) _save() (string, ...interface{}) {
	// todo
	return sql, args
}

func (m *EventModel) _create(props map[string]interface{}) (string, ...interface{}) {
	// todo
	return sql, args
}

func (m *EventModel) _delete() (string, int64) {
	// todo
	return sql, id
}

func (m *EventModel) _destroy(id int64) (string, int)  {
	sql := "DELETE FROM event WHERE id = ?"
	return sql, id
}

func (m *EventModel) _update(props map[string]interface{}, conds map[string]interface{}) (string, ...interface{}) {
	// todo
	return sql, args
}

func (m *EventModel) createtable() {
	// todo
}

func (m *EventModel) new() *EventModel {
	// todo
	return m
}

func (m *EventModel) find(id int64) (*EventModel, error) {
	// todo
	return m, err
}

func (m *EventModel) where(conds map[string]interface{}) []*EventModel {
	// todo
	ms := []*EventModel{}
	return ms
}

func (m *EventModel) save() (*EventModel, error) {
	// todo
	return m, err
}

func (m *EventModel) create(props map[string]interface{}) (*EventModel, error) {
	// todo
	return m, err
}

func (m *EventModel) delete() error {
	// todo
	return err
}

func (m *EventModel) destroy(id int64) error {
	// todo
	retrun err
}

func (m *EventModel) update(props map[string]interface{}, conds map[string]interface{}) error {
	// todo
	return err
}

var Event = EventModel{datasource: "default", table: "event"}