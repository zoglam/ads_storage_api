package images

// Images ...
type Images struct {
    ID         int64   `json:"-"`
    Ref        string  `json:"ref"`
    DataCreate string  `json:"-"`
    AdsID      float32 `json:"-"`
}

type imagesDaoInterface interface {
    CreateImage(ref string, adsID int64) (int64, error)
}
