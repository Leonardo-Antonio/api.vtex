package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api.pdq-prices/src/entity"
	"github.com/Leonardo-Antonio/api.pdq-prices/src/util/response"
	"github.com/labstack/echo/v4"
)

type Formating struct{}

func (f *Formating) urlImages(path string) ([]string, error) {
	fsInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var nameImages []string
	for _, file := range fsInfo {
		nameImages = append(nameImages, file.Name())
	}

	return nameImages, nil
}

func (f *Formating) SearchSkuEanAndRemplaceCoincidences(c echo.Context) error {
	var data []*entity.Image
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErr(err.Error(), nil))
	}

	skusEan, err := f.urlImages("portadas")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErr(err.Error(), nil))
	}

	for _, skuEanSearch := range skusEan {
		for _, dataTotal := range data {
			nameImg := strings.ReplaceAll(skuEanSearch, ".jpg", "")
			if strings.EqualFold(nameImg, dataTotal.SkuEan) {
				dataTotal.Label = "front view"
				dataTotal.Url = fmt.Sprintf("%s/%s", "https://esedor.com/macro/macro-3009", skuEanSearch)
			}
		}
	}

	return c.JSON(http.StatusOK, response.NewSatisfactory("ok", data))
}

func (f Formating) SearchAndRemplaceTypeProductDigitalBySpecification(c echo.Context) error {
	var data *entity.ProductsSpecification
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErr(err.Error(), nil))
	}

	var success []entity.Specification
	for _, search := range data.Search {
		for _, dataProduct := range data.DataSpecification {
			if strings.EqualFold(search, dataProduct.Name) {
				dataProduct.SpecificationValue = "Físico"
				success = append(success, *dataProduct)
			}
		}
	}
	log.Println(len(data.Search))
	return c.JSON(http.StatusOK, response.NewSatisfactory("ok", success))
}
