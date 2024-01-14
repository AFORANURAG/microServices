package getters

const (
	GetUserWithNameQuery = `SELECT * FROM users WHERE name=(?)`
	GetUserWithRowId     = `SELECT * FROM users WHERE id=(?)`
)
