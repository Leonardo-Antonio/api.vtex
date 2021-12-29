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

type ImgSku struct{}

func NewImgSku() *ImgSku {
	return &ImgSku{}
}

func (i *ImgSku) Set(c echo.Context) error {
	body := new(entity.ImagesSkus)

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	for _, image := range *body {
		client := &http.Client{}

		bodyRequest, err := json.Marshal(image)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		url := fmt.Sprintf("https://%s.vtexcommercestable.com.br/api/catalog/pvt/stockkeepingunit/%s/file", enum.Env.AccountName, image.SkuId)
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyRequest))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("X-VTEX-API-AppKey", enum.Env.AppKey)
		req.Header.Add("X-VTEX-API-AppToken", enum.Env.AppToken)

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if resp.StatusCode != 200 {
			log.Println(fmt.Sprintf("| %s | SkuId -> %s", resp.Status, image.SkuId))
		}

		log.Println(fmt.Sprintf("| %s | SkuId -> %s", resp.Status, image.SkuId))
	}
	return c.JSON(http.StatusCreated, response.NewSatisfactory("ok", nil))
}

func (i *ImgSku) DeleteAllImgBySku(c echo.Context) error {
	skus := []string{}
	if err := c.Bind(&skus); err != nil {
		return c.JSON(http.StatusOK, response.NewErr("error", err))
	}

	for _, id := range skus {
		client := new(http.Client)

		url := fmt.Sprintf("https://%s.vtexcommercestable.com.br/api/catalog/pvt/stockkeepingunit/%s/file", enum.Env.AccountName, id)
		req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(nil))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.NewErr("error", err.Error()))
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("X-VTEX-API-AppKey", enum.Env.AppKey)
		req.Header.Add("X-VTEX-API-AppToken", enum.Env.AppToken)

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.NewErr("error", err.Error()))
		}

		if resp.StatusCode != 200 {
			log.Println(fmt.Sprintf("| %s | SkuId -> %s", resp.Status, id))
		}

		log.Println(fmt.Sprintf("| %s | SkuId -> %s", resp.Status, id))
	}

	return c.JSON(http.StatusOK, response.NewSatisfactory("eliminados", skus))
}
