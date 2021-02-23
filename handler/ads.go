package handler

import (
    "errors"
    "net/http"
    "strconv"

    ads "github.com/zoglam/ads_storage_api/models/ads"
)

// GetAds ...
func GetAds(w http.ResponseWriter, r *http.Request) {
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

// GetAd ...
func GetAd(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            getErrorPage(w, r, err)
        }
    }()

}

// CreateAd ...
func CreateAd(w http.ResponseWriter, r *http.Request) {
    var err error
    defer func() {
        if err != nil {
            getErrorPage(w, r, err)
        }
    }()

    r.ParseForm()
    title := r.PostForm.Get("title")
    description := r.PostForm.Get("description")

    images := []string{
        r.PostForm.Get("img1"),
        r.PostForm.Get("img2"),
        r.PostForm.Get("img3"),
    }
    filteredImages := images[:0]
    for _, ref := range images {
        if ref != "" {
            filteredImages = append(filteredImages, ref)
        }
    }
    images = filteredImages
    if len(images) == 0 {
        err = errors.New("Images not found")
        return
    }

    price := r.PostForm.Get("price")

    _, err = ads.AdsDao.CreateAd(title, description, images, price)
    if err != nil {
        return
    }

    getStatusOKPage(w, r, nil)
}
