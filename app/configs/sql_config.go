package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var sqlConfig *gorm.DB = nil

func MigrationIntializer() {
	db, err := gorm.Open(mysql.Open(formatDBURL()))
	if err != nil {
		log.Fatalln("[DB] unable to connect to DB")
	}

	sqlConfig = db
}

func formatDBURL() string {
	dbConfig := GetEnv()
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConfig.DbUsername, dbConfig.DbPassword, dbConfig.DbURL, dbConfig.DbName)
}

func GetDBConnection() *gorm.DB {
	return sqlConfig
}
