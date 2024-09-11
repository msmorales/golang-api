package models

// Challenge representa la tabla de desafíos en la base de datos
type Challenge struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"` // Clave primaria
	Title       string `json:"title" binding:"required"`           // Campo obligatorio
	Description string `json:"description"`                        // Campo no obligatorio
	UserID      uint   `json:"user_id"`                            // Campo no obligatorio, clave foránea
}
