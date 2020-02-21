package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DataSourceName
const (
	host     = "localhost"
	port     = 5432
	user     = "dev"
	password = "password"
	dbname   = "dev"
	sslmode  = "disable"
)

// IGormClient exposes the database connection
type IGormClient interface {
	//SetupDB connection
	SetupDB()
	// Check verifies that the database was injected
	Check()
	// Close allows for the database to be closed
	Close()
}

// SetupDB connects to the database
func (gc *GormClient) SetupDB() {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	var err error
	gc.crDB, err = gorm.Open("postgres", dataSourceName)
	if err != nil {
	}
}

func (gc *GormClient) Check() bool {
	return gc.crDB != nil
}

func (gc *GormClient) Close() {
	gc.crDB.Close()
}
