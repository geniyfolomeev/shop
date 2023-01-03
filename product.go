package shop

type Product struct {
	Id          int     `json:"id" db:"id"`
	Title       string  `json:"title" binding:"required" db:"title"`
	Price       float64 `json:"price" binding:"required" db:"price"`
	Description string  `json:"description" db:"description"`
	Category    string  `json:"category" binding:"required" db:"category"`
	Image       string  `json:"image" db:"image"`
	Rating      string  `json:"rating" binding:"required" db:"rating"`
}
