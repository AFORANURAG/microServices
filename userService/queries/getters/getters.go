package getters

const (
	GetUserWithNameQuery     = `SELECT * FROM users WHERE name=(?)`
	GetUserWithRowId         = `SELECT * FROM users WHERE id=(?)`
	GetUserWithEmail         = `SELECT * FROM users WHERE email = ?`
	UpdateVerificationStatus = `UPDATE users SET isVerified = ? WHERE email= ?`
)
