package cat

import (
	"fmt"
	"github.com/danyazab/animal/internal/animal/model"
	"github.com/danyazab/animal/internal/animal/model/util"
	"github.com/danyazab/animal/internal/api/http/request"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

const birthdayLayout = "2006-01-02"

type Crud struct {
	dig.In
	Repository model.CatRepository
}

func (cntr *Crud) Create(ec echo.Context) error {
	var req request.CreateCatReq
	if err := ec.Bind(&req); err != nil {
		return err
	}

	birthday, _ := time.Parse(birthdayLayout, req.Birthday)
	cat := model.Cat{
		Name:        req.Name,
		Description: req.Description,
		Breed:       req.Breed,
		Birthday:    birthday,
		Sex:         util.TypeSex(req.Sex),
		TailLength:  req.TailLength,
		Color:       req.Color,
		WoolType:    util.TypeWool(req.WoolType),
		IsChipped:   req.IsChipped,
		Weight:      req.Weight,
	}

	res, err := cntr.Repository.Store(ec.Request().Context(), cat)
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusCreated, map[string]uint{
		"id": res.ID,
	})
}

func (*Crud) Edit(ec echo.Context) error {
	return ec.NoContent(http.StatusOK)
}

func (*Crud) List(ec echo.Context) error {
	res := []string{
		"kkk",
	}

	return ec.JSON(http.StatusOK, res)
}

func (cntr *Crud) Info(ec echo.Context) error {
	var id uint
	if err := echo.PathParamsBinder(ec).Uint("catId", &id).BindError(); err != nil {
		return err
	}

	entity, found, err := cntr.Repository.FindByID(ec.Request().Context(), id)
	if err != nil {
		return err
	}

	if !found {
		return echo.NewHTTPError(
			http.StatusNotFound,
			fmt.Sprintf("The Cat by id %d not found", id),
		)
	}

	return ec.JSON(http.StatusOK, entity)
}

func (*Crud) Delete(ec echo.Context) error {
	var id uint
	if err := echo.QueryParamsBinder(ec).Uint("catId", &id).BindError(); err != nil {
		return err
	}

	return ec.NoContent(http.StatusOK)
}
