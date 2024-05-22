package main

import (
	"fmt"
	"kitchen-service/app"
	"kitchen-service/config"
	"kitchen-service/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	// Initialize database
	db, err := config.InitDB(cfg)
	if err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		return
	}
	defer db.Close()

	handler := app.SetupApp(db)
	e := echo.New()
	e.Use(middleware.Logger())

	routes.RoutesApi(e, handler)

	e.Use(middleware.Recover())

	// Start server
	e.Logger.Fatal(e.Start(":8081"))

}
