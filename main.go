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
)

func main() {
	app := echo.New()
	// Connect to DB
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("error opening database,", err)
	}

	app.Use(middleware.CORS())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "${status} | ${latency_human} | ${method} | ${uri} | ${error} \n",
		},
	))

	// app.Use(middleware.Recover())

	// trailling slash

	app.HTTPErrorHandler = appMiddleware.TemplHTTPErrorHandler

	handler := handler.NewHandler(db)
	app.Static("static/", "static")
	app.Pre(middleware.RemoveTrailingSlash())

	app.GET("/", handler.Index)

	game := app.Group("/games")
	game.GET("/new", handler.CreateGame)
	game.GET("/new/generate", handler.GenerateGame)
	game.GET("/join/:game_id", handler.JoinGame)
	game.GET("/delete/:game_id", handler.DeleteGame)

	dashboard := app.Group("/games/dashboard")
	dashboard.GET("/:game_id", handler.Dashboard)

	log.Fatal(app.Start(":" + os.Getenv("PORT")))
}
