package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AFORANURAG/microServices/backend/apiGateway/pkg/routes"
	"github.com/AFORANURAG/microServices/backend/apiGateway/pkg/utils"

	middleware "github.com/AFORANURAG/microServices/backend/apiGateway/pkg/middlewares"

	"github.com/AFORANURAG/microServices/backend/apiGateway/pkg/configs"

	"github.com/gofiber/fiber/v2"
	log "github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file:%v", err)
	}
	fmt.Println(os.Getenv("HOST"))
	fiberConfig := configs.FiberConfig()
	app := fiber.New(fiberConfig)
	middleware.FiberMiddleware(app)
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("%s/logs/app-out.log", wd))
	file, err := os.OpenFile(fmt.Sprintf("%s/logs/app-out.log", wd), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening the file in main file %v", err)
	}
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)
	log.SetLevel(log.LevelInfo)
	log.SetOutput(mw)

	requestResponseLoggerFile, err_ := os.OpenFile(fmt.Sprintf("%s/logs/reqres.log", wd), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err_ != nil {
		fmt.Printf("error opening file: %v", err)
	}
	app.Use(logger.New(logger.Config{
		Format:        "${time} - ${ip}:${port} => ${status} - ${latency}\n",
		TimeFormat:    "02-Jan-2006",
		TimeZone:      "UTC",
		Output:        requestResponseLoggerFile,
		DisableColors: false,
	}))

	defer requestResponseLoggerFile.Close()
	routeFuctions := routes.AllRoutes()
	for _, route := range routeFuctions {
		route(app, "api", "v1")
	}
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}

}
