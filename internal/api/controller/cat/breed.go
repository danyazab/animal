package cat

import (
	"github.com/danyazab/animal/internal/infrastructure/network/petfinder"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Breed struct {
	dig.In

	Client petfinder.Client
}

func (cntr *Breed) List(ec echo.Context) error {
	breads, err := cntr.Client.GetCatsBreeds(ec.Request().Context())
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, breads)
}
