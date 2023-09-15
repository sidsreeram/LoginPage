package database

import (
	"fmt"
	"project1/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect(connect_to string) {
	DB, err = gorm.Open(postgres.Open(connect_to),&gorm.Config{})
	if err!=nil {
		panic("cannot connect to the database")
	}
	fmt.Println("connected to the database")
}
//migrating users struct to users table
func Migrator()  {
	DB.AutoMigrate(&models.Users{})
	fmt.Println("tables created successfully")
}

//migrating admin struct to admin table
func Migrateadmin(){
	DB.AutoMigrate(&models.Admins{})
	fmt.Println("admin tables created successfully...")
}