package getters

const (
	GetUserWithNameQuery     = `SELECT * FROM users WHERE name=(?)`
	GetUserWithRowId         = `SELECT * FROM users WHERE id=(?)`
	GetUserWithEmail         = `SELECT * FROM users WHERE email = ?`
		GetUserWithPhoneNumber         = `SELECT * FROM users WHERE phoneNumber = ?`

	UpdateVerificationStatus = `UPDATE users SET isVerified = ? WHERE email= ?`
	UpdateVerificationStatusWithPhoneNumber=`UPDATE users SET isVerified = ? WHERE phoneNumber= ?`
	UpdateVerificationStatusWithPhoneNumberInOTPTable=`UPDATE OTP SET isVerified = ? WHERE phoneNumber= ? AND otpNumber = ?`
	

)
