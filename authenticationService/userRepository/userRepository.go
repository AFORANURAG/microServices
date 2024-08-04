package userRepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	userSchema "github.com/AFORANURAG/microServices/authenticationService/models"
	getters "github.com/AFORANURAG/microServices/authenticationService/queries/getters"
	dbservice "github.com/AFORANURAG/microServices/authenticationService/services/dbService"
)

type IUserRepository interface {
	GetUserByName(name string) (*sql.Row, error)
	GetUserByRowId(rowID string) (*sql.Row, error)
	GetUserWithEmail(email string) (*userSchema.NewUser, error)
		GetUserWithPhoneNumber(phoneNumber string) (*userSchema.NewUser, error)

	CreateUser(userData *userSchema.UserSchema) (*sql.Result, error)
	UpdateVerificationStatus(status bool, email *string,phoneNumber *string,otp int) (bool, error)
	
}

type UserRepository struct {
	dbs *dbservice.MYSQLDBService
}

func (s *UserRepository) GetUserByName(name string) (*sql.Row, error) {
	db, dbErr := s.dbs.GetDb()
	if dbErr != nil {
		log.Printf("Error while getting DB: %v", dbErr)
		return nil, dbErr
	}

	// Prepare the query using a constant or variable for better code maintainability
	query := getters.GetUserWithNameQuery

	// Use QueryRowContext to handle timeouts or cancellations
	row := db.QueryRowContext(context.TODO(), query, name)

	return row, nil
}

func (s *UserRepository) GetUserWithEmail(email string) (*userSchema.NewUser, error) {
	db, Dberr := s.dbs.GetDb()
	log.Printf("<-----------------userName is %v", email)
	if Dberr != nil {
		log.Printf("Error while getting DB  :%v", Dberr)
		return &userSchema.NewUser{}, Dberr
	}

	res := db.QueryRowContext(context.TODO(), getters.GetUserWithEmail, email)
	// Row,err
	var response userSchema.NewUser
	var createdAtString string;
	scanErr := res.Scan(&response.UserID, &response.Email, &response.Name, &response.PhoneNumber,&response.IsVerified,&createdAtString)

	return &response, scanErr
}

func (s *UserRepository) GetUserWithPhoneNumber(phoneNumber string) (*userSchema.NewUser, error) {
	db, Dberr := s.dbs.GetDb()
	log.Printf("<-----------------phoneNumber is is %v", phoneNumber)
	if Dberr != nil {
		log.Printf("Error while getting DB  :%v", Dberr)
		return &userSchema.NewUser{}, Dberr
	}

	res := db.QueryRowContext(context.TODO(), getters.GetUserWithPhoneNumber, phoneNumber)
	// Row,err
	// Row,err
	var response userSchema.NewUser
	var createdAtString string;
	scanErr := res.Scan(&response.UserID, &response.Email, &response.Name, &response.PhoneNumber,&response.IsVerified,&createdAtString)

	return &response, scanErr
}

func (s *UserRepository) GetUserByRowId(rowID string) (*sql.Row, error) {
	db, Dberr := s.dbs.GetDb()
	if Dberr != nil {
		log.Printf("Error while getting DB  :%v", Dberr)
		return nil, Dberr
	}

	res := db.QueryRow(getters.GetUserWithRowId, rowID)

	return res, nil
}
func (s *UserRepository) CreateUser(userData *userSchema.UserSchema) (*sql.Result, error) {
	query := `INSERT INTO users (name, email,isVerified,phoneNumber) VALUES (?, ?,?,?)`
	// The `Exec` function takes in a query, as well as the values for
	// the parameters in the order they are defined
	// s.dbs.Exec
	db, err := s.dbs.GetDb()
	if err != nil {
		log.Printf("Error while getting db instance :%v ", err)
	}
	// Check Whether the user is there or not
	res, err := db.Exec(query, userData.Name, userData.Email, false,userData.PhoneNumber)
	if err != nil {
		log.Printf("(CreateUser) db.Exec: Error is %v", err)
		return nil, err
	}

	return &res, err
}


func (s *UserRepository) UpdateVerificationStatus(status bool, email *string, phoneNumber *string,otp int) (bool, error) {
	fmt.Printf("<------------\notp is %d --------->",otp)
	var query string
	var updateisVerifiedInOTPTableQuery string=getters.UpdateVerificationStatusWithPhoneNumberInOTPTable

	if phoneNumber != nil {
		query = getters.UpdateVerificationStatusWithPhoneNumber
	} else {
		query = getters.UpdateVerificationStatus
	}

	db, err := s.dbs.GetDb()
	if err != nil {
		log.Printf("Error while getting db instance: %v", err)
		return false, err
	}

	var updateIsVerifiedErr error
	_, updateOTPVerifiedInOTPTableError:=db.Exec(updateisVerifiedInOTPTableQuery,status,phoneNumber,otp)

	if updateOTPVerifiedInOTPTableError!=nil{
		return false,err
	}
	if phoneNumber != nil {
		_, updateIsVerifiedErr = db.Exec(query, status, phoneNumber)
	} else {
		_, updateIsVerifiedErr = db.Exec(query, status, email)
	}

	if updateIsVerifiedErr != nil {
		log.Printf("Error while updating isVerified status: %v", updateIsVerifiedErr)
		return false, updateIsVerifiedErr
	}

	return true, nil
}

func NewUserRepositoryProvider(db *dbservice.MYSQLDBService) IUserRepository {
	return &UserRepository{dbs: db}
}
