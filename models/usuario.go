package models

type Usuario struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Nombre string `gorm:"size:100" json:"nombre"`
	Correo string `gorm:"size:100;unique" json:"correo"`
}
