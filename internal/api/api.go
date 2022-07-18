package api

import (
	"danyazab/animal/internal/api/controller/cat"
	"danyazab/animal/internal/api/controller/swagger"
	"fmt"

	"github.com/labstack/echo"
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
	app.GET("/pet/cat/breads", a.Bread.List)
}
