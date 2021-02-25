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
            ad.id,
            ad.title,
            (
                SELECT img.ref
                FROM IMAGES as img
                WHERE ad.id = img.ads_id
                ORDER BY data_create
                LIMIT 1
            ) as image,
            price
        FROM ADS as ad
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
        err := rows.Scan(
            &ad.ID,
            &ad.Title,
            &ad.Image,
            &ad.Price,
        )
        if err != nil {
            return nil, err
        }
        ads = append(ads, ad)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return ads, nil
}

func (a *adsDao) GetByAdTitle(adID string, hasDescription bool, hasAllImages bool) (Ads, error) {
    rows, err := models.DB.Query(`
        SELECT 
            ad.id,
            ad.title,
            IFNULL(ad.description, ""),
            IFNULL(img1.ref, ""),
            IFNULL(img2.ref, ""),
            IFNULL(img3.ref, ""),
            ad.price
        FROM ADS as ad
            LEFT JOIN IMAGES as img1
                ON ad.id=img1.ads_id
            LEFT JOIN IMAGES as img2
                ON ad.id=img2.ads_id AND img2.id!=img1.id
            LEFT JOIN IMAGES as img3
                ON ad.id=img3.ads_id AND img3.id!=img1.id AND img3.id!=img2.id
        WHERE ad.id=?
        LIMIT 1
    `, adID)
    if err != nil {
        return Ads{}, err
    }
    defer rows.Close()

    var ad Ads
    for rows.Next() {
        var img1, img2, img3 string
        var description string
        err := rows.Scan(&ad.ID, &ad.Title, &description, &img1, &img2, &img3, &ad.Price)
        if err != nil {
            return Ads{}, err
        }
        ad.Image = img1

        if hasAllImages {
            for _, ref := range []string{img1, img2, img3} {
                ad.Images = append(ad.Images, ref)
            }
        }

        if hasDescription {
            if description == "" {
                ad.Description = "empty"
            } else {
                ad.Description = description
            }
        }
    }
    if err = rows.Err(); err != nil {
        return Ads{}, err
    }
    return ad, nil
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
