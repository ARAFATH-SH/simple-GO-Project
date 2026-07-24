package domain

// model or entity -> existance ->
type Product struct {
	Id          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgURL      string  `json:"image_url" db:"image_url"`
}
