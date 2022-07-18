package cat

import (
	"danyazab/animal/internal/animal/model"
	"danyazab/animal/internal/animal/model/util"
	"danyazab/animal/internal/api/http/request"
	"net/http"
	"time"

	"github.com/labstack/echo"
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

func (*Crud) List(ec echo.Context) error {
	res := []string{
		"kkk",
	}

	return ec.JSON(http.StatusOK, res)
}
