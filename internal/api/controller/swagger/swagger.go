package swagger

import (
	"bytes"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"text/template"

	"go.uber.org/dig"
)

const SwaggerConfPath = "./api/api-v1.yaml"
const SwaggerConfUrl = "/api/api-v1.yaml"

type TplParams struct {
	Url string
}

type Swagger struct {
	dig.In
}

func (*Swagger) View(ec echo.Context) error {
	p := TplParams{Url: SwaggerConfUrl}

	res, err := parse(Tpl, &p)

	if err != nil {
		return err
	}

	ec.Response().Header().Set("Content-type", "text/html; charset=UTF-8")

	return ec.String(http.StatusOK, res)
}

func parse(t string, data interface{}) (string, error) {
	n := uuid.New().String()
	tmpl, err := template.New(n).Parse(t)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
