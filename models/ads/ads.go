package ads

// Ads ...
type Ads struct {
    ID          int64    `json:"id,omitempty"`
    Title       string   `json:"title"`
    Description string   `json:"description,omitempty"`
    Price       float32  `json:"price"`
    Image       string   `json:"main_image,omitempty"`
    Images      []string `json:"images,omitempty"`
    DataCreate  string   `json:"-"`
}

type adsDaoInterface interface {
    GetAdsByPageNumber(pageNumber string, sortBy string, orderBy bool) ([]Ads, error)
    GetByAdTitle(adID string, hasDescription bool, hasAllImages bool) (Ads, error)
    CreateAd(title string, description string, images []string, price string) (int64, error)
}
