package route_helper

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sarakhanx/go-try-something/controller/debug"
)

func Hifolks(app *fiber.App) {
	app.Get("/hifolks", helper.Hifolks)
	app.Post("/gethifolks", helper.GetHifolks)

}
