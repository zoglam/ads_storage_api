package models

type adsDao struct{}

// AdsDao - data access object
var AdsDao adsDaoInterface

func init() {
	AdsDao = &adsDao{}
}

func (a *adsDao) All(pageNumber int64, sortParam string) ([]Ads, error) {

}

func (a *adsDao) Get(adsTitle string, description bool, allImages bool) (Ads, error) {

}

func (a *adsDao) Create(title string, description string, images []string, price string) (int64, error) {

}
