package models

type Contenido struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Titulo string `gorm:"size:100" json:"titulo"`
	Genero string `gorm:"size:50" json:"genero"`
	Tipo   string `gorm:"size:50" json:"tipo"`
}
