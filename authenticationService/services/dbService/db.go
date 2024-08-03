package dbservice

import (
	"database/sql"
	"log"
	"sync"

	userservicetypes "github.com/AFORANURAG/microServices/authenticationService/types/userServiceTypes"
	_ "github.com/go-sql-driver/mysql"
)

type MYSQLDBService struct {
	db  *sql.DB
}

func (dbs *MYSQLDBService) GetDb() (*sql.DB, error) {
	return dbs.db, nil
}

func (dbs *MYSQLDBService) Exec(query string, args ...any) (*sql.Rows, error) {
	_, err :=
		dbs.db.Exec(query, args...)
	if err != nil {
		log.Printf("error while inserting value : %v", err)
	}
	res, err := dbs.db.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("Error while quering for all users:%v", err)
		return nil, err
	}
	return res, nil
}

var DbClient *MYSQLDBService
var once sync.Once

func NewDBServiceClientProvider(uri userservicetypes.UserServicePhrase) *MYSQLDBService {
	once.Do(func() {
		url:=(string)(uri)
		db, err := sql.Open("mysql", url)
		if err != nil {
			log.Printf("Error while creating DBServiceClient:%v", err)
		}
		log.Println("Successfully connected to PlanetScale!")
		DbClient = &MYSQLDBService{db: db}
	})
	return DbClient
}
