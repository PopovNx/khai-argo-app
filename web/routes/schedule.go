package routes

import (
	"github.com/gofiber/fiber/v2"
	"lab-argo-app/pkg/schedule"
	"lab-argo-app/web/handlers"
)

const AppVersion = "1.0.1"

func ScheduleRouter(app fiber.Router, service schedule.Service) {
	app.Get("/", handlers.GetSchedule(service))
	app.Get("/lessons/:day", handlers.GetDay(service))
	app.Get("/version", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app":     "khai-argo-app",
			"version": AppVersion,
		})
	})
}
