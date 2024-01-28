package logging

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var LoggingConfig = logger.Config{
	Format:     "${time} ${locals:requestid} ${status} | ${method} ${path} ${error}\n",
	TimeFormat: "15:04:05",
	TimeZone:   "UTC",
}
