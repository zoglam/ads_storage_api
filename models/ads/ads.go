package ads

// Ads ...
type Ads struct {
    ID          int64    `json:"-"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Price       float32  `json:"price"`
    Images      []string `json:"image"`
    DataCreate  string   `json:"-"`
}

type adsDaoInterface interface {
    GetAdsByPageNumber(pageNumber string, sortBy string, orderBy bool) ([]Ads, error)
    GetByAdTitle(adTitle string, hasDescription bool, hasAllImages bool) (Ads, error)
    CreateAd(title string, description string, images []string, price string) (int64, error)
}
