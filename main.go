package main

import (
	"context"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/mccune1224/betrayal-widget/handler"
	"github.com/mccune1224/betrayal-widget/route"

	//appMiddleware "github.com/mccune1224/betrayal-widget/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	app := echo.New()
	// Connect to DB
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("error opening database,", err)
	}

	handler := handler.NewHandler(db)
	app.Pre(middleware.RemoveTrailingSlash())

	//TODO: Setup CORS for frontend domain
	// app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodHead, http.MethodTrace},
	// }))

	// CORS allow ALL origins and ALL methods and headers
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "${status} | ${latency_human} | ${method} | ${uri} | ${error} \n",
		},
	))

	route.Routes(app, handler)
	log.Fatal(app.Start(":" + os.Getenv("PORT")))
}
