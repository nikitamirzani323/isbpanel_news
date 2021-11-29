package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/isbpanel_news/controllers"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	app.Post("api/login", controllers.Login)
	app.Post("api/valid", controllers.Checkpage)
	app.Post("api/news", controllers.News)
	app.Post("api/newssave", controllers.Newssave)
	app.Post("api/newsdelete", controllers.Newsdelete)
	app.Post("api/categorynews", controllers.Categorynews)
	return app
}
