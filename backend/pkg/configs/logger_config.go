package configs

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerConfig() logger.Config {
	wd, err := os.Getwd()
	if err != nil {
		wd = "../../"
	}
	file, err := os.OpenFile(fmt.Sprintf("%s/logs/reqres.log", wd), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	return logger.Config{
		Format:     "${time} - ${ip}:${port} => ${status} - ${latency}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
		Output:     file,
	}
}
