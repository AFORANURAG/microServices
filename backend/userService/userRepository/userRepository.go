package userRepository

import (
	"database/sql"
	"log"

	userSchema "github.com/AFORANURAG/microServices/backend/userService/models"
	getters "github.com/AFORANURAG/microServices/backend/userService/queries/getters"
	dbservice "github.com/AFORANURAG/microServices/backend/userService/services/dbService"
)

type IUserRepository interface {
	GetUserByName(name string) (*sql.Row, error)
	GetUserByRowId(rowID string) (*sql.Row, error)
	CreateUser(userData *userSchema.UserSchema) (*sql.Result, error)
}

type UserRepository struct {
	dbs *dbservice.MYSQLDBService
}

func (s *UserRepository) GetUserByName(name string) (*sql.Row, error) {
	db, Dberr := s.dbs.GetDb()
	if Dberr != nil {
		log.Fatalf("Error while getting DB  :%v", Dberr)
		return nil, Dberr
	}

	res := db.QueryRow(getters.GetUserWithNameQuery, name)

	return res, nil
}

func (s *UserRepository) GetUserByRowId(rowID string) (*sql.Row, error) {
	db, Dberr := s.dbs.GetDb()
	if Dberr != nil {
		log.Fatalf("Error while getting DB  :%v", Dberr)
		return nil, Dberr
	}

	res := db.QueryRow(getters.GetUserWithRowId, rowID)

	return res, nil
}
func (s *UserRepository) CreateUser(userData *userSchema.UserSchema) (*sql.Result, error) {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	// The `Exec` function takes in a query, as well as the values for
	// the parameters in the order they are defined
	// s.dbs.Exec
	db, err := s.dbs.GetDb()
	if err != nil {
		log.Fatalf("Error while getting db instance :%v ", err)
	}
	// Check Whether the user is there or not
	res, err := db.Exec(query, userData.Name, userData.Email)
	if err != nil {
		log.Fatal("(CreateUser) db.Exec", err)
	}

	return &res, err
}

func NewUserRepositoryProvider(db *dbservice.MYSQLDBService) *UserRepository {
	return &UserRepository{dbs: db}
}
