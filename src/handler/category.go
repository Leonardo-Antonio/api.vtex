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

type Category struct{}

func (c *Category) Create(ctx echo.Context) error {
	var categories []entity.Category
	if err := ctx.Bind(&categories); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewErr("la estructura no es valida", nil))
	}

	if len(categories) == 0 {
		return ctx.JSON(http.StatusBadRequest, response.NewErr("debe contener al menos un item", nil))
	}

	for i, category := range categories {
		client := &http.Client{}
		b, _ := json.Marshal(category)
		req, err := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf("https://%s.vtexcommercestable.com.br/api/catalog/pvt/category", enum.Env.AccountName),
			bytes.NewBuffer(b),
		)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.NewErr(err.Error(), category.Name))
		}
		req.Header.Add("X-VTEX-API-AppKey", enum.Env.AppKey)
		req.Header.Add("X-VTEX-API-AppToken", enum.Env.AppToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json; charset=utf-8")

		resp, err := client.Do(req)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError,
				response.NewErr(
					fmt.Sprintf("la peticion no  se logro realizar, code: %s", resp.Status), category.Name))
		}

		if resp.StatusCode != 200 {
			return ctx.JSON(http.StatusInternalServerError,
				response.NewErr(
					fmt.Sprintf("la peticion no  se logro realizar, code: %s", resp.Status), category.Name))
		}

		log.Println(fmt.Sprintf("%d.- Code: | %s | Name -> | %s |", i, resp.Status, category.Name))
	}

	return ctx.JSON(http.StatusCreated, response.NewSatisfactory("las categorias se crearon correctamente", nil))
}
