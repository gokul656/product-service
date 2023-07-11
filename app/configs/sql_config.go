package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnection struct {
	sqlConfig *gorm.DB
}

func (conn *DBConnection) MigrationIntializer() {
	log.Println("[DB] connection to MySQL..")
	db, err := gorm.Open(mysql.Open(formatDBURL()))
	if err != nil {
		log.Fatalln("[DB] unable to connect to DB")
	}

	conn.sqlConfig = db
}

func (conn *DBConnection) GetDBConnection() *gorm.DB {
	return conn.sqlConfig
}

func formatDBURL() string {
	dbConfig := GetEnv()
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbConfig.DbUsername, dbConfig.DbPassword, dbConfig.DbURL, dbConfig.DbName)
}
