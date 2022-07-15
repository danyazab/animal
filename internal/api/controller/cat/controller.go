package cat

import (
	"danyazab/animal/internal/animal/model"
	"danyazab/animal/internal/animal/model/util"
	"danyazab/animal/internal/api/http/request"
	"danyazab/animal/internal/api/http/response"
	"github.com/labstack/echo"
	"go.uber.org/dig"
	"net/http"
	"time"
)

const birthdayLayout = "2006-01-02"

type Controller struct {
	dig.In
	Repository model.CatRepository
}

func (cntr *Controller) Create(ec echo.Context) error {
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

func (cntr *Controller) List(ec echo.Context) error {
	cats, total, err := cntr.Repository.GetAll(ec.Request().Context())
	if err != nil {
		return err
	}

	items := make([]response.CatResp, len(cats))
	for i, cat := range cats {
		items[i] = response.CatResp{
			Id:    cat.ID,
			Name:  cat.Name,
			Age:   23, // FIXME Danylo
			Breed: cat.Breed,
			Sex:   string(cat.Sex),
			Color: cat.Color,
		}
	}

	return ec.JSON(http.StatusOK, response.CatRespList{
		Items: items,
		Meta: response.MetaData{
			Total: total,
		},
	})
}
