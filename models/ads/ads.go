package models

// Ads ...
type Ads struct {
	ID          int64    `json:"-"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	Images      []string `json:"image"`
}

// daoInterface ...
type adsDaoInterface interface {
	All(pageNumber int64, sortParam string) ([]Ads, error)
	Get(adsTitle string, hasDescription bool, hasAllImages bool) (Ads, error)
	Create(title string, description string, images []string, price string) (int64, error)
}
