package request

type MenuRequest struct {
	CategoryID  string `json:"category_id"  validate:"required"`
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	PriceModal  int    `json:"price_modal" validate:"required"`
	Description string `json:"description"  validate:"required"`
}
