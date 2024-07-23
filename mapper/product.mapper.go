package mapper

import (
	"github.com/your/repo/model"
	"github.com/your/repo/schema"
)

type ProductMapper struct {
}

func (m *ProductMapper) toSchema(product *model.Product) *schema.ProductSchema {
	return &schema.ProductSchema{
		Name:   product.GetName(),
		Price:  product.GetPrice(),
		Status: product.GetStatus(),
	}
}

func (m *ProductMapper) toModel(schema *schema.ProductSchema) *model.Product {
	return &model.Product{
		ID:        schema.ID,
		Name:      schema.Name,
		Price:     schema.Price,
		Status:    schema.Status,
		CreatedAt: schema.CreatedAt,
		UpdatedAt: schema.UpdatedAt,
	}
}
