package schema

import "time"

type ProductSchema struct {
	ID        *int          `gorm:"type:int unsigned;auto_increment;primaryKey"`
	Name      string        `gorm:"type:varchar(255);not null"`
	Price     int           `gorm:"type:int unsigned;default:0"`
	Status    ProductStatus `gorm:"type:enum('display', 'hide', 'delete');default:'display'"`
	CreatedAt *time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;omitempty"`
	UpdatedAt *time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime;omitempty"`
}

func (ProductSchema) TableName() string {
	return "products"
}
