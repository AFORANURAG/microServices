package userRepository

import (
	"context"
	"database/sql"
	"log"

	userSchema "github.com/AFORANURAG/microServices/authenticationService/models"
	getters "github.com/AFORANURAG/microServices/authenticationService/queries/getters"
	dbservice "github.com/AFORANURAG/microServices/authenticationService/services/dbService"
)

type IUserRepository interface {
	GetUserByName(name string) (*sql.Row, error)
	GetUserByRowId(rowID string) (*sql.Row, error)
	GetUserWithEmail(email string) (*userSchema.UserSchema, error)
	CreateUser(userData *userSchema.UserSchema) (*sql.Result, error)
	UpdateVerificationStatus(status bool, email string) (bool, error)
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

func (s *UserRepository) GetUserWithEmail(email string) (*userSchema.UserSchema, error) {
	db, Dberr := s.dbs.GetDb()
	log.Printf("<-----------------userName is %v", email)
	if Dberr != nil {
		log.Printf("Error while getting DB  :%v", Dberr)
		return &userSchema.UserSchema{}, Dberr
	}

	res := db.QueryRowContext(context.TODO(), getters.GetUserWithEmail, email)
	// Row,err
	var response userSchema.UserSchema
	scanErr := res.Scan(&response.Id, &response.Name, &response.Email, &response.IsVerified)

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
	query := `INSERT INTO users (name, email,isVerified) VALUES (?, ?, ?)`
	// The `Exec` function takes in a query, as well as the values for
	// the parameters in the order they are defined
	// s.dbs.Exec
	db, err := s.dbs.GetDb()
	if err != nil {
		log.Printf("Error while getting db instance :%v ", err)
	}
	// Check Whether the user is there or not
	res, err := db.Exec(query, userData.Name, userData.Email, false)
	if err != nil {
		log.Printf("(CreateUser) db.Exec", err)
		return nil, err
	}

	return &res, err
}

func (s *UserRepository) UpdateVerificationStatus(status bool, email string) (bool, error) {
	query := getters.UpdateVerificationStatus
	db, err := s.dbs.GetDb()
	if err != nil {
		log.Printf("Error while getting db instance :%v ", err)
	}
	_, updateIsVerifiedErr := db.Exec(query, status, email)
	if updateIsVerifiedErr != nil {
		log.Printf("Error While updating isVerifiedStatus : %v\n", err)
		return false, nil
	}
	return true, nil
}

func NewUserRepositoryProvider(db *dbservice.MYSQLDBService) IUserRepository {
	return &UserRepository{dbs: db}
}
