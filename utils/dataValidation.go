package utils

import (
    "errors"
    "regexp"
)

type dataValidationInterface interface {
    GetValidatedPrice(price string) (string, error)
    GetValidatedImages(images []string) ([]string, error)
}

type dataValidation struct{}

// DataValidation ...
var DataValidation dataValidationInterface

func init() {
    DataValidation = &dataValidation{}
}

func (d *dataValidation) GetValidatedPrice(price string) (string, error) {
    re := regexp.MustCompile(`^\d{1,10}(?:\.\d{1,2})?$`)
    matched := re.Match([]byte(price))
    if matched == false {
        return "", errors.New("Price not Found")
    }
    return string(re.Find([]byte(price))), nil
}

func (d *dataValidation) GetValidatedImages(images []string) ([]string, error) {
    filteredImages := images[:0]
    for _, ref := range images {
        if ref != "" {
            filteredImages = append(filteredImages, ref)
        }
    }
    if len(filteredImages) == 0 {
        return nil, errors.New("Images not found")
    }
    return filteredImages, nil
}
