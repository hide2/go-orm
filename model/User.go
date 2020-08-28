type UserModel struct {
	datasource string
	table      string

	id int64

	name string

}

func (m *UserModel) _createtable() {
	// todo
	return sql
}

func (m *UserModel) _find(id int64) (string, int64)  {
	sql := "SELECT * FROM user WHERE id = ?"
	return sql, id
}

func (m *UserModel) _where(conds map[string]interface{}) (string, ...interface{}) {
	// todo
	return sql, args
}

func (m *UserModel) _save() (string, ...interface{}) {
	// todo
	return sql, args
}

func (m *UserModel) _create(props map[string]interface{}) (string, ...interface{}) {
	// todo
	return sql, args
}

func (m *UserModel) _delete() (string, int64) {
	// todo
	return sql, id
}

func (m *UserModel) _destroy(id int64) (string, int)  {
	sql := "DELETE FROM user WHERE id = ?"
	return sql, id
}

func (m *UserModel) _update(props map[string]interface{}, conds map[string]interface{}) (string, ...interface{}) {
	// todo
	return sql, args
}

func (m *UserModel) createtable() {
	// todo
}

func (m *UserModel) new() *UserModel {
	// todo
	return m
}

func (m *UserModel) find(id int64) (*UserModel, error) {
	// todo
	return m, err
}

func (m *UserModel) where(conds map[string]interface{}) []*UserModel {
	// todo
	ms := []*UserModel{}
	return ms
}

func (m *UserModel) save() (*UserModel, error) {
	// todo
	return m, err
}

func (m *UserModel) create(props map[string]interface{}) (*UserModel, error) {
	// todo
	return m, err
}

func (m *UserModel) delete() error {
	// todo
	return err
}

func (m *UserModel) destroy(id int64) error {
	// todo
	retrun err
}

func (m *UserModel) update(props map[string]interface{}, conds map[string]interface{}) error {
	// todo
	return err
}

var User = UserModel{datasource: "default", table: "user"}