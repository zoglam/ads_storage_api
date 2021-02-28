package controllers

import (
    "errors"
    "regexp"
    "strings"
)

type dataValidationInterface interface {
    GetPrice(price string) (string, error)
    GetImages(images string) ([]string, error)
}

type dataValidation struct{}

// DataValidation ...
var DataValidation dataValidationInterface

func init() {
    DataValidation = &dataValidation{}
}

func (d *dataValidation) GetPrice(price string) (string, error) {
    re := regexp.MustCompile(`^\d{1,10}(?:\.\d{1,2})?$`)
    matched := re.Match([]byte(price))
    if matched == false {
        return "", errors.New("Price not Found")
    }
    return string(re.Find([]byte(price))), nil
}

func (d *dataValidation) GetImages(images string) ([]string, error) {
    filteredImages := make([]string, 0)
    for _, ref := range strings.Split(images, ",") {
        if ref != "" {
            filteredImages = append(filteredImages, ref)
        }
    }
    if len(filteredImages) == 0 {
        return nil, errors.New("Images not found")
    }
    return filteredImages, nil
}
