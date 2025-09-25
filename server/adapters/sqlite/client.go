package sqlite

import (
	"log"
	"path"

	"github.com/rubenalves-dev/taskuik/adapters"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var FILE_PATH string = path.Join("data", "taskuik.db")
var CONFIG gorm.Config = gorm.Config{}

type SQLiteClient struct {
	Filepath string
	Config   gorm.Config
	DB       *gorm.DB
}

var _ adapters.Database = &SQLiteClient{}

func NewClient() adapters.Database {
	db, err := gorm.Open(sqlite.Open(FILE_PATH), &CONFIG)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return &SQLiteClient{
		Filepath: FILE_PATH,
		Config:   CONFIG,
		DB:       db,
	}
}

func (c *SQLiteClient) Connect() error {

	return nil
}

func (c *SQLiteClient) Disconnect() error {
	// Implementation for disconnecting from the database
	return nil
}

func (c *SQLiteClient) GetDB() *gorm.DB {
	return c.DB
}
