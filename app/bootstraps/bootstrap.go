package bootstraps

import (
	"log"

	"github.com/gokul656/product-service/app/configs"
	"github.com/gokul656/product-service/app/dao"
)

var conn = &configs.DBConnection{}

func GetConnection() *configs.DBConnection {
	return conn
}

func Setup() {
	conn.MigrationIntializer()
	setupDAO()
}

func setupDAO() {
	db := conn.GetDBConnection()

	// Do not remove any struct unless it's not necessary
	err := db.AutoMigrate(&dao.User{})
	if err != nil {
		log.Fatalln("[DB] error", err)
	}

}
