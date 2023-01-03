package shop

type Product struct {
	Id          int     `json:"-"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      string  `json:"rating"`
}
