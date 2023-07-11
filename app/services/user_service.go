package services

import (
	"errors"

	"github.com/gokul656/product-service/app/bootstraps"
	db "github.com/gokul656/product-service/app/configs"
	"github.com/gokul656/product-service/app/dao"
	"gorm.io/gorm"
)

var conn *db.DBConnection = bootstraps.GetConnection()

func getDB() *gorm.DB {
	return conn.GetDBConnection()
}

func CreateUser(request *dao.UserDTO) (*dao.User, error) {
	user := &dao.User{
		Name: request.Name,
		Role: "USER",
		Age:  request.Age,
	}

	_, err := GetUserByName(request.Name)
	if err == nil {
		return nil, errors.New("username already exisits")
	}

	id := getDB().Create(user)
	if id == nil {
		return nil, errors.New("unable to create user")
	}

	return user, nil
}

func GetUser(id string) (*dao.User, error) {
	user := &dao.User{}
	getDB().Find(user, id)

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func GetUserByName(name string) (*dao.User, error) {
	user := &dao.User{}
	getDB().Find(user, "name = ?", name)
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func GetAllUsers() []dao.User {
	var users []dao.User
	conn.GetDBConnection().Find(&users)

	return users
}

func DeleteUserByName(name string) any {
	return conn.GetDBConnection().Where("name = ?", name).Delete(&dao.User{})
}
