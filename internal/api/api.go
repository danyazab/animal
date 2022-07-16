package api

import (
	"danyazab/animal/internal/api/controller/cat"
	"danyazab/animal/internal/api/controller/swagger"
	"fmt"
	"github.com/labstack/echo"
	"go.uber.org/dig"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type API struct {
	dig.In

	cat.Controller
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

	app.POST("/pet/cat", a.Create)
	app.GET("/pet/cat", a.List)
}
