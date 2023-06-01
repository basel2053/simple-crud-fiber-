package db

import (
	"log"

	"github.com/basel2053/simple-crud-fiber-/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(URL string) *gorm.DB{
	db,err :=gorm.Open(postgres.Open(URL),&gorm.Config{})

	if err !=nil{
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Book{})
	return db
}