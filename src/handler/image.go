package handler

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/Leonardo-Antonio/api.pdq-prices/src/util/enum"
	"github.com/Leonardo-Antonio/api.pdq-prices/src/util/response"
	"github.com/labstack/echo/v4"
)

type ImageSku struct{}

func (i *ImageSku) DeleteBySkuId(c echo.Context) error {
	var skuIds []string
	if err := c.Bind(&skuIds); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErr(err.Error(), nil))
	}

	if len(skuIds) == 0 {
		return c.JSON(http.StatusBadRequest, response.NewErr("debe contener al menos un item", nil))
	}

	for i, skuId := range skuIds {
		client := &http.Client{}
		req, err := http.NewRequest(
			http.MethodDelete,
			fmt.Sprintf("https://%s.vtexcommercestable.com.br/api/catalog/pvt/stockkeepingunit/%s/file",
				enum.Env.AccountName, skuId),
			bytes.NewBuffer(nil),
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.NewErr(err.Error(), skuId))
		}
		req.Header.Add("X-VTEX-API-AppKey", enum.Env.AppKey)
		req.Header.Add("X-VTEX-API-AppToken", enum.Env.AppToken)

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				response.NewErr(
					fmt.Sprintf("la peticion no  se logro realizar, code: %s", resp.Status), skuId))
		}

		if resp.StatusCode != 200 {
			return c.JSON(http.StatusInternalServerError,
				response.NewErr(
					fmt.Sprintf("la peticion no  se logro realizar, code: %s", resp.Status), skuId))
		}

		log.Println(fmt.Sprintf("%d.- Code: | %s | skuId -> | %s |", i, resp.Status, skuId))
	}

	return c.JSON(http.StatusOK, response.NewSatisfactory("se eliminaron las todas imagenes de los skuId ingresasdos", nil))
}

// https://editorialmacroperu.vtexcommercestable.com.br/api/catalog/pvt/stockkeepingunit/13834/file
