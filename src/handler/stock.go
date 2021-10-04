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

type Stock struct{}

func (s *Stock) Set(ctx echo.Context) error {
	var stocks []entity.Stock
	if err := ctx.Bind(&stocks); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewErr(err.Error(), nil))
	}

	for i, stock := range stocks {
		client := &http.Client{}
		b, _ := json.Marshal(stock)
		req, err := http.NewRequest(
			http.MethodPut,
			fmt.Sprintf(
				"https://%s.vtexcommercestable.com.br/api/logistics/pvt/inventory/skus/%s/warehouses/%s",
				enum.Env.AccountName, stock.SkuId, stock.WarehousesId,
			),
			bytes.NewBuffer(b),
		)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.NewErr(err.Error(), stock.SkuId))
		}
		req.Header.Add("X-VTEX-API-AppKey", enum.Env.AppKey)
		req.Header.Add("X-VTEX-API-AppToken", enum.Env.AppToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(req)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError,
				response.NewErr(
					fmt.Sprintf("la peticion no  se logro realizar, code: %s", resp.Status), stock.SkuId))
		}

		log.Println(fmt.Sprintf("%d.- skuId: %s - code: %s", i, stock.SkuId, resp.Status))

		if resp.StatusCode != 200 {
			return ctx.JSON(http.StatusInternalServerError,
				response.NewErr(
					fmt.Sprintf("la peticion no  se logro realizar, code: %s", resp.Status), stock.SkuId))
		}
	}

	return ctx.JSON(http.StatusOK, response.NewSatisfactory(
		"el stock de los productos se actualizo correctamente", stocks))
}
