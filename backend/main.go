package main

import (
	"backend/pkg/configs"
	middleware "backend/pkg/middlewares"
	"backend/pkg/routes"
	"backend/pkg/utils"
	"fmt"

	"os"

	"io"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Println(os.Getenv("HOST"))
	fiberConfig := configs.FiberConfig()
	app := fiber.New(fiberConfig)
	middleware.FiberMiddleware(app)
	file, err := os.OpenFile("./logs/app-out.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening the file in main file %v", err)
	}
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)
	log.SetLevel(log.LevelInfo)
	log.SetLevel(log.LevelError)
	log.SetOutput(mw)
	routeFuctions := routes.AllRoutes()
	for _, route := range routeFuctions {
		route(app, "api", "v1")
	}

	// Rate limiter , dynamic rates
	app.Use(limiter.New(limiter.Config{
		Max:               30,
		Expiration:        time.Minute * 1,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
