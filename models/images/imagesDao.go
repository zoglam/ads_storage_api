package images

import (
    models "github.com/zoglam/ads_storage_api/models"
)

type imagesDao struct{}

// ImagesDao - data access object
var ImagesDao imagesDaoInterface

func init() {
    ImagesDao = &imagesDao{}
}

func (a *imagesDao) CreateImage(ref string, adsID int64) (int64, error) {

    statement, err := models.DB.Prepare(`
        INSERT INTO IMAGES(ref, ads_id)
        VALUES(?,?)
    `)
    if err != nil {
        return 0, err
    }

    res, err := statement.Exec(ref, adsID)
    if err != nil {
        return 0, err
    }

    lid, err := res.LastInsertId()
    if err != nil {
        return 0, err
    }
    return lid, nil
}
