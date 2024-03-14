package main

import (
	"encoding/json"
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

	app.HTTPErrorHandler = appMiddleware.TemplHTTPErrorHandler

	handler := handler.NewHandler(db)
	app.Static("static/", "static")
	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORS())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "${status} | ${latency_human} | ${method} | ${uri} | ${error} \n",
		},
	))

	sm := appMiddleware.NewSyncMiddleware(db)
	app.Use(sm.SyncGameInfo)
	app.GET("/", handler.Index)

	game := app.Group("/games")
	game.GET("/new", handler.CreateGame)
	game.GET("/new/generate", handler.GenerateGame)
	game.GET("/join/:game_id", handler.JoinGame)
	game.GET("/delete/:game_id", handler.DeleteGame)

	dashboard := app.Group("/games/dashboard/:game_id")
	dashboard.GET("", handler.Dashboard)
	dashboard.GET("/menu", handler.PlayerMenu)

	playerUpdate := dashboard.Group("/menu/update/:player")
	playerUpdate.POST("/modifier", handler.UpdatePlayerLuckModifier)
	playerUpdate.POST("/alive", handler.UpdatePlayerDeathStatus)
	playerUpdate.POST("/seat", handler.UpdatePlayerSeating)
	playerUpdate.POST("/alliance", handler.UpdatePlayerAlliance)

	playerInsert := dashboard.Group("/players")
	playerInsert.GET("", handler.PlayerDashboard)
	playerInsert.POST("/add", handler.PlayerAdd)

	alliancesDashboard := dashboard.Group("/alliances")
	alliancesDashboard.GET("", handler.AllianceDashboard)
	alliancesDashboard.POST("/new", handler.AllianceCreate)
	// alliancesDashboard.POST("/update", handler.AllianceUpdate)

	seatingDashboard := dashboard.Group("/:game_id/seating")
	seatingDashboard.GET("", handler.SeatingDashboard)
	seatingDashboard.POST("/swap", handler.SwapSeats)

	luckDashboard := dashboard.Group("/:game_id/luck")
	luckDashboard.GET("", handler.Luck)
	luckDashboard.POST("/update", handler.LuckUpdate)

	if os.Getenv("PORT") == "3000" {
		data, err := json.MarshalIndent(app.Routes(), "", "  ")
		if err != nil {
			panic(err)
		}
		os.WriteFile("routes.json", data, 0644)
	}

	log.Fatal(app.Start(":" + os.Getenv("PORT")))
}
