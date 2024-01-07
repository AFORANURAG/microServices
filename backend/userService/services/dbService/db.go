package dbservice

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type MYSQLDBService struct {
	db  *sql.DB
	DSN string
}

var instance *MYSQLDBService
var once sync.Once

func NewDBServiceClientProvider(uri string) *MYSQLDBService {
	once.Do(func() {
		db, err := sql.Open("mysql", uri)
		if err != nil {
			log.Fatalf("Error while creating DBServiceClient:%v", err)
		}
		log.Println("Successfully connected to PlanetScale!")
		instance = &MYSQLDBService{db: db}
	})
	return instance
}
