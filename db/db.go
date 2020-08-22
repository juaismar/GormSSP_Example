package db

import (
	"GormSSP_Example/models"
	_ "GormSSP_Example/routers"
	"fmt"

	"github.com/jinzhu/gorm"
)

func InitORM() {
	if models.Conn == nil {
		models.Conn = openDB("localhost",
			"5432",
			"postgres",
			"GormSSP_Example",
			"postgres")

		CreateSchema()
		PopulateUsers()
	}
}

func openDB(host, port, user, dbname, pwd string) *gorm.DB {
	sslMode := "disable"

	dbSource := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host,
		port,
		user,
		dbname,
		pwd,
		sslMode,
	)
	gormDB, err := gorm.Open("postgres", dbSource)

	if err != nil {
		panic(err)
	}
	return gormDB
}

func CreateSchema() {
	createIfNotExist(&models.User{})
}
func createIfNotExist(m interface{}) {
	tblExist := models.Conn.HasTable(m)
	if !tblExist {
		err := models.Conn.CreateTable(m).Error
		if err != nil {
			panic(err)
		}
	}
}
