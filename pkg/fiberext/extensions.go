package fiberext

import "github.com/gofiber/fiber/v2"

func RealIP(c *fiber.Ctx) string {
	// client can set the X-Forwarded-For or the X-Real-IP header to any arbitrary value it wants,
	// unless you have a trusted reverse proxy, you shouldn't use any of those values.
	if ip := c.Get(fiber.HeaderXForwardedFor); ip != "" {
		ips := c.IPs()
		if len(ips) > 0 {
			return ips[0]
		}
	}
	if ip := c.Get("X-Real-IP"); ip != "" {
		return ip
	}
	// return remote addr
	return c.IP()
}
