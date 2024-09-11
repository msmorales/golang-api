package models

// User representa la tabla de usuarios en la base de datos
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"` // Clave primaria
	Name  string `json:"name" binding:"required"`            // Campo obligatorio
	Email string `json:"email" binding:"required"`           // Campo obligatorio
}
