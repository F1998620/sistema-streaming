package database

import (
	"log"

	"sistema-streaming/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarBD() {

	dsn := "root:1234567890@tcp(127.0.0.1:3306)/streaming?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error al conectar con MySQL:", err)
	}

	DB.AutoMigrate(
		&models.Usuario{},
		&models.Contenido{},
		&models.Suscripcion{},
	)

	log.Println("Conexión exitosa a MySQL")
}
