package logging

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

var LoggingConfig = logger.Config{
	Format:       "[${time}] ${locals:requestid} ${status} | ${latency} | ${method} ${path} ${error}\n",
	TimeFormat:   "15:04:05 15:04:05.000",
	TimeInterval: time.Millisecond,
	TimeZone:     "UTC",
}
