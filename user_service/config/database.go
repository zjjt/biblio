package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //importing the dialect postgres
)

//CreatePostgresDBConnection returns an oppen connection to the database
func CreatePostgresDBConnection() (*gorm.DB, error) {
	var host, user, dbname, password string
	if len([]byte(os.Getenv("DB_HOST"))) > 0 {
		//if we enter this if then we are in production
		host = os.Getenv("DB_HOST")
		user = os.Getenv("DB_USER")
		dbname = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	} else {
		host = GetEnvVariable("DB_HOST")
		user = GetEnvVariable("DB_USER")
		dbname = GetEnvVariable("DB_NAME")
		password = GetEnvVariable("DB_PASSWORD")
	}

	return gorm.Open(
		"postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname),
	)

}
