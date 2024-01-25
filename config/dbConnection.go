package config

import (
	"air-q/Models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=BACKend2023. dbname=postgres  sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("db connection failed")
	}

	DB = db
	fmt.Println("db connected successfully.")
	AutoMigrate(db)

}

func AutoMigrate(connection *gorm.DB) {
	connection.AutoMigrate(
		&Models.User{},
		&Models.Organization{},
		&Models.Device{},
		&Models.Device_Log{},
	)
}
