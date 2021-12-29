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

type Order struct{}
type reason struct {
	Reason string `json:"reason"`
}

func (p *Order) Cancel(c echo.Context) error {
	var listIdOrder []string
	if err := c.Bind(&listIdOrder); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErr(err.Error(), nil))
	}

	if len(listIdOrder) == 0 {
		return c.JSON(http.StatusBadRequest, response.NewErr("debe contener al menos un item", nil))
	}

	for i, idOrder := range listIdOrder {
		log.Println(fmt.Sprintf("https://%s.vtexcommercestable.com.br/api/oms/pvt/orders/%s/cancel", enum.Env.AccountName, idOrder))
		client := &http.Client{}
		req, err := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf("https://%s.vtexcommercestable.com.br/api/oms/pvt/orders/%s/cancel", enum.Env.AccountName, idOrder),
			bytes.NewBuffer([]byte{}),
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.NewErr(err.Error(), idOrder))
		}
		req.Header.Add("X-VTEX-API-AppKey", enum.Env.AppKey)
		req.Header.Add("X-VTEX-API-AppToken", enum.Env.AppToken)

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				response.NewErr(
					fmt.Sprintf("la peticion no  se logro realizar, code: %s", resp.Status), idOrder))
		}

		if resp.StatusCode != 200 {
			return c.JSON(http.StatusInternalServerError,
				response.NewErr(
					fmt.Sprintf("la peticion no  se logro realizar, code: %s", resp.Status), idOrder))
		}

		log.Println(fmt.Sprintf("%d.- Code: | %s | skuId -> | %s |", i, resp.Status, idOrder))
	}

	return c.JSON(http.StatusOK, response.NewSatisfactory("pricios actualizados y agregados correctamente", listIdOrder))
}
