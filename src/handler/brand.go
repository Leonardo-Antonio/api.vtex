package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Leonardo-Antonio/api.pdq-prices/src/entity"
	"github.com/Leonardo-Antonio/api.pdq-prices/src/util/enum"
	"github.com/Leonardo-Antonio/api.pdq-prices/src/util/response"
	"github.com/labstack/echo/v4"
)

type brand struct{}

func NewBrand() *brand {
	return &brand{}
}

func (b *brand) Create(c echo.Context) error {
	body := new([]*entity.Brand)
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErr(err.Error(), nil))
	}

	for _, value := range *body {
		data, err := json.Marshal(value)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.NewErr(err.Error(), fmt.Sprintf("Id: %s", value.Id)))
		}

		req, err := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf("https://%s.vtexcommercestable.com.br/api/catalog/pvt/brand", enum.Env.AccountName),
			bytes.NewBuffer(data),
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.NewErr(err.Error(), fmt.Sprintf("Id: %s", value.Id)))
		}

		req.Header.Add("X-VTEX-API-AppKey", enum.Env.AppKey)
		req.Header.Add("X-VTEX-API-AppToken", enum.Env.AppToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json; charset=utf-8")

		resp, err := new(http.Client).Do(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.NewErr(err.Error(), fmt.Sprintf("Id: %s", value.Id)))
		}

		log.Println(fmt.Sprintf("%s | ID: %s", resp.Status, value.Id))
	}

	return c.JSON(http.StatusCreated, response.NewSatisfactory("ok", body))
}
