package datastore

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// GetDBConnection : GET DB Connection
func GetDBConnection() *gorm.DB {
	portNum := os.Getenv("DBPORT")
	dbUser := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	db, err := gorm.Open("postgres", "port="+portNum+" user="+dbUser+" dbname="+dbName)
	if err != nil {
		panic(err)
	}
	return db
}
