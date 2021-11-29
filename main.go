package main

import (
	"log"

	"github.com/nikitamirzani323/isbpanel_news/routers"
)

func main() {
	app := routers.Init()
	log.Fatal(app.Listen(":6061"))
}
