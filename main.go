package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/mccune1224/betrayal-widget/handler"

	appMiddleware "github.com/mccune1224/betrayal-widget/middleware"
	"github.com/mccune1224/betrayal-widget/route"
)

func main() {
	app := echo.New()
	// Connect to DB
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("error opening database,", err)
	}

	handler := handler.NewHandler(db)
	app.Static("static/", "static")
	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORS())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "${status} | ${latency_human} | ${method} | ${uri} | ${error} \n",
		},
	))

	route.ViewRoutes(app, handler)

	sm := appMiddleware.NewSyncMiddleware(db)
	app.Use(sm.SyncGameInfo)
	log.Fatal(app.Start(":" + os.Getenv("PORT")))
}
