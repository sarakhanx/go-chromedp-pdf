package helper

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sarakhanx/go-try-something/config/chromedp"
)

func Hifolks(c fiber.Ctx) error {
	return c.SendString("hi folks")
}

func GetHifolks(c fiber.Ctx) error {
	chromedp.PdfInit()

	return c.SendString("PDF was saved in root folder")

}
