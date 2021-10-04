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

type Price struct{}

func (p *Price) Set(c echo.Context) error {
	var prices []entity.Price
	if err := c.Bind(&prices); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErr("la estructura no es valida", nil))
	}

	if len(prices) == 0 {
		return c.JSON(http.StatusBadRequest, response.NewErr("debe contener al menos un item", nil))
	}

	for i, price := range prices {
		client := &http.Client{}
		b, _ := json.Marshal(price)
		req, err := http.NewRequest(
			http.MethodPut,
			fmt.Sprintf("https://api.vtex.com/%s/pricing/prices/%s", enum.Env.AccountName, price.SkuId),
			bytes.NewBuffer(b),
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.NewErr(err.Error(), price.SkuId))
		}
		req.Header.Add("X-VTEX-API-AppKey", enum.Env.AppKey)
		req.Header.Add("X-VTEX-API-AppToken", enum.Env.AppToken)

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				response.NewErr(
					fmt.Sprintf("la peticion no  se logro realizar, code: %s", resp.Status), price.SkuId))
		}

		if resp.StatusCode != 200 {
			return c.JSON(http.StatusInternalServerError,
				response.NewErr(
					fmt.Sprintf("la peticion no  se logro realizar, code: %s", resp.Status), price.SkuId))
		}

		log.Println(fmt.Sprintf("%d.- Code: | %s | skuId -> | %s |", i, resp.Status, price.SkuId))
	}

	return c.JSON(http.StatusOK, response.NewSatisfactory("pricios actualizados y agregados correctamente", nil))
}
