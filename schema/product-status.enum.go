package schema

type ProductStatus string

const (
	Display ProductStatus = "display"
	Hide    ProductStatus = "hide"
	Delete  ProductStatus = "delete"
)
