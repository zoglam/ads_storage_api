package handler

import (
    "errors"
    "net/http"
    "strconv"
    "strings"

    controllers "github.com/zoglam/ads_storage_api/controllers"
    ads "github.com/zoglam/ads_storage_api/models/ads"
)

func getAds(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            getErrorPage(w, r, err)
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

    getStatusOKPage(w, r, ads)
}

func getAd(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            getErrorPage(w, r, err)
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

    getStatusOKPage(w, r, ad)
}

func createAd(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            getErrorPage(w, r, err)
        }
    }()

    r.ParseForm()
    title := r.PostForm.Get("title")
    if title == "" {
        err = errors.New("Title not Found")
        return
    }
    description := r.PostForm.Get("description")

    images, err := controllers.DataValidation.GetImages(r.PostForm.Get("images"))
    if err != nil {
        return
    }

    price, err := controllers.DataValidation.GetPrice(r.PostForm.Get("price"))
    if err != nil {
        return
    }

    lid, err := ads.AdsDao.CreateAd(title, description, images, price)
    if err != nil {
        return
    }

    getStatusOKPage(w, r, lid)
}
