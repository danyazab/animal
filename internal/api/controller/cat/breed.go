package cat

import (
	"danyazab/animal/internal/infrastructure/network/petfinder"
	"github.com/labstack/echo"
	"go.uber.org/dig"
	"net/http"
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
