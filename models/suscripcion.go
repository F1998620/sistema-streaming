package models

type Suscripcion struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	NombrePlan  string  `gorm:"size:100" json:"nombre_plan"`
	Precio      float64 `json:"precio"`
	Descripcion string  `gorm:"type:text" json:"descripcion"`
}
