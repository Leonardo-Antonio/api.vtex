package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api.pdq-prices/src/entity"
	"github.com/Leonardo-Antonio/api.pdq-prices/src/util/response"
	"github.com/labstack/echo/v4"
)

type Slug struct{}

func (s *Slug) SetByNameProduct(ctx echo.Context) error {
	var slugs []*entity.Slug
	if err := ctx.Bind(&slugs); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewErr(err.Error(), nil))
	}

	for _, category := range slugs {
		s.replaceSpecialCharacters(&category.Name)
	}

	return ctx.JSON(http.StatusOK, response.NewSatisfactory("se hizo el formateo", slugs))
}

func (s *Slug) replaceSpecialCharacters(slug *string) {
	searchAndRemplace := func(old, new string) {
		*slug = strings.ReplaceAll(*slug, old, new)
	}

	*slug = strings.ToLower(*slug)

	specialCharacters := []string{"/", ":", ".", "#", "[", "]", "&", "/", "*", "+", "$", "'", "(", ")", "=", "|", "°", ",", "%"}
	for _, character := range specialCharacters {
		searchAndRemplace(character, "")
	}

	letters := map[string]string{
		"n": "ñ",
		"y": "&",
		"e": "é",
		"o": "ó",
		"i": "í",
		"u": "ú",
		"a": "á",
	}

	for new, old := range letters {
		searchAndRemplace(old, new)
	}

	searchAndRemplace(" ", "-")
	searchAndRemplace("ü", "u")

	*slug = strings.TrimRight(*slug, "-")

	log.Println(*slug)
}
