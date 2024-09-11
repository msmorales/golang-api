package models

// Video representa la tabla de videos en la base de datos
type Video struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"` // Clave primaria
	Title       string `json:"title" binding:"required"`           // Campo obligatorio
	Content     string `json:"content" binding:"required"`         // Campo obligatorio
	Description string `json:"description"`                        // Campo no obligatorio
	UserID      uint   `json:"user_id"`                            // Campo no obligatorio, clave foránea
}
