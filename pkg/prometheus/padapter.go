package prometheus

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func FastHTTPAdapter() fiber.Handler {
	return adaptor.HTTPHandler(promhttp.Handler())
}
