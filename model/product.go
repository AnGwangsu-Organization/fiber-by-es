package model

import (
	"github.com/your/repo/schema"
	"time"
)

type Product struct {
	ID        *int
	Name      string
	Price     int
	Status    schema.ProductStatus
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (p *Product) GetId() *int {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() int {
	return p.Price
}

func (p *Product) GetStatus() schema.ProductStatus {
	return p.Status
}

func (p *Product) GetCreatedAt() *time.Time {
	return p.CreatedAt
}

func (p *Product) GetUpdatedAt() *time.Time {
	return p.UpdatedAt
}
