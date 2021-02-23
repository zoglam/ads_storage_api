package ads

import (
    "fmt"

    models "github.com/zoglam/ads_storage_api/models"
    images "github.com/zoglam/ads_storage_api/models/images"
)

type adsDao struct{}

// AdsDao - data access object
var AdsDao adsDaoInterface

func init() {
    AdsDao = &adsDao{}
}

func (a *adsDao) GetAdsByPageNumber(pageNumber string, sortBy string, orderBy bool) ([]Ads, error) {

    sortByField := "price"
    if sortBy == "data" {
        sortByField = "data_create"
    }

    sortOrder := "asc"
    if !orderBy {
        sortOrder = "desc"
    }

    query := fmt.Sprintf(`
        SELECT
            a.title,
            (
                SELECT img.ref
                FROM IMAGES as img
                WHERE a.id = img.ads_id
                ORDER BY data_create
                LIMIT 1
            ) as image,
            price
        FROM ADS as a
        ORDER BY %s %s LIMIT %s, 10
    `, sortByField, sortOrder, pageNumber)
    rows, err := models.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ads []Ads
    for rows.Next() {
        var ad Ads
        var img string
        err := rows.Scan(&ad.Title, &img, &ad.Price)
        if err != nil {
            return nil, err
        }
        ad.Images = append(ad.Images, img)
        ads = append(ads, ad)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return ads, nil
}

func (a *adsDao) GetByAdTitle(adsTitle string, description bool, allImages bool) (Ads, error) {
    return Ads{}, nil
}

func (a *adsDao) CreateAd(title string, description string, imagesList []string, price string) (int64, error) {

    statement, err := models.DB.Prepare(`
        INSERT INTO ADS(title, description, price)
        VALUES(?,?,?)
    `)
    if err != nil {
        return 0, err
    }

    res, err := statement.Exec(title, description, price)
    if err != nil {
        return 0, err
    }

    lid, err := res.LastInsertId()
    if err != nil {
        return 0, err
    }

    for _, image := range imagesList {
        _, err := images.ImagesDao.CreateImage(image, lid)
        if err != nil {
            return 0, err
        }
    }

    return lid, nil
}
