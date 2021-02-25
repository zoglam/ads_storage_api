package handler

import (
    "errors"
    "net/http"
    "strconv"
    "strings"

    ads "github.com/zoglam/ads_storage_api/models/ads"
    utils "github.com/zoglam/ads_storage_api/utils"
)

type adsHandler struct{}

// Ads ...
var Ads adsHandlerInterface

func init() {
    Ads = &adsHandler{}
}

// GetAds ...
func (a *adsHandler) GetAds(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            Error.getErrorPage(w, r, err)
        }
    }()

    params := r.URL.Query()
    pageNumber := params.Get("page")
    sortParam := params.Get("sort_by")
    orderIN, err := strconv.ParseBool(params.Get("order"))
    if err != nil {
        return
    }

    ads, err := ads.AdsDao.GetAdsByPageNumber(pageNumber, sortParam, orderIN)
    if err != nil {
        return
    }

    success.getStatusOKPage(w, r, ads)
}

// GetAd ...
func (a *adsHandler) GetAd(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            Error.getErrorPage(w, r, err)
        }
    }()

    var hasDescription, hasAllImages bool
    params := r.URL.Query()
    adID := params.Get("id")
    if adID == "" {
        err = errors.New("ID not Found")
        return
    }

    fields := params.Get("fields")

    for _, field := range strings.Split(fields, ",") {
        if field == "description" {
            hasDescription = true
        } else if field == "images" {
            hasAllImages = true
        }
    }

    ad, err := ads.AdsDao.GetByAdTitle(adID, hasDescription, hasAllImages)
    if err != nil {
        return
    }

    success.getStatusOKPage(w, r, ad)
}

// CreateAd ...
func (a *adsHandler) CreateAd(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            Error.getErrorPage(w, r, err)
        }
    }()

    r.ParseForm()
    title := r.PostForm.Get("title")
    if title == "" {
        err = errors.New("Title not Found")
        return
    }
    description := r.PostForm.Get("description")

    images, err := utils.DataValidation.GetValidatedImages([]string{
        r.PostForm.Get("img1"),
        r.PostForm.Get("img2"),
        r.PostForm.Get("img3"),
    })
    if err != nil {
        return
    }

    price, err := utils.DataValidation.GetValidatedPrice(r.PostForm.Get("price"))
    if err != nil {
        return
    }

    lid, err := ads.AdsDao.CreateAd(title, description, images, price)
    if err != nil {
        return
    }

    success.getStatusOKPage(w, r, lid)
}
