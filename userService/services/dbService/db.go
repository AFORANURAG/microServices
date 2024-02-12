package dbservice

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type MYSQLDBService struct {
	db  *sql.DB
	dsn string
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

func NewDBServiceClientProvider(uri string) *MYSQLDBService {
	once.Do(func() {
		db, err := sql.Open("mysql", uri)
		if err != nil {
			log.Printf("Error while creating DBServiceClient:%v", err)
		}
		log.Println("Successfully connected to PlanetScale!")
		DbClient = &MYSQLDBService{db: db}
	})
	return DbClient
}
