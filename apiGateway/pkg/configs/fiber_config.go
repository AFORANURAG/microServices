package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/AFORANURAG/microServices/apiGateway/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
)

func FiberConfig() fiber.Config {
	// seeach for the env and if you don't find it , give some error
	wd, err := os.Getwd()
	if err != nil {
		wd = "../../"
	}
	engine := html.New(fmt.Sprintf("%s/app/views", wd), ".html")
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		// Global Error Handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Info(err)
			return c.Status(fiber.StatusInternalServerError).JSON(utils.GlobalErrorHandlerResp{Success: false, Message: "Internal Server Error"})
		},
		Views: engine,
	}
}
