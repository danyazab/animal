package cat

import (
	"github.com/labstack/echo"
	"go.uber.org/dig"
	"net/http"
)

type Controller struct {
	dig.In
}

func (cntr *Controller) Create(ec echo.Context) error {
	return ec.NoContent(http.StatusCreated)
}

func (cntr *Controller) List(ec echo.Context) error {
	res := []string{
		"kkk",
	}

	return ec.JSON(http.StatusOK, res)
}
