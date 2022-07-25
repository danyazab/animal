package api

import (
	"fmt"
	"github.com/danyazab/animal/internal/api/controller/cat"
	"github.com/danyazab/animal/internal/api/controller/swagger"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type API struct {
	dig.In

	CatCrud cat.Crud
	Bread   cat.Breed
	swagger.Swagger
}

func RunServer(api API) error {
	port := 8000

	app := echo.New()
	api.initRoutes(app)

	return app.Start(fmt.Sprintf(":%d", port))
}

func (a *API) initRoutes(app *echo.Echo) {
	app.Static(swagger.SwaggerConfUrl, swagger.SwaggerConfPath)
	app.GET("/swagger", a.View)

	app.POST("/pet/cat", a.CatCrud.Create)
	app.GET("/pet/cat", a.CatCrud.List)
	app.GET("/pet/cat/:catId", a.CatCrud.Info)
	app.PUT("/pet/cat/:catId", a.CatCrud.Edit)
	app.DELETE("/pet/cat/:catId", a.CatCrud.Delete)

	app.GET("/pet/cat/breads", a.Bread.List)
}
