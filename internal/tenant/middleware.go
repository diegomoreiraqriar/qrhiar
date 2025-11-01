package tenant

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Middleware(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantID := c.Get("X-Tenant-ID")
		if tenantID == "" {
			tenantID = "default-tenant"
		}
		c.Locals("tenant_id", tenantID)
		logger.Info("request", zap.String("tenant", tenantID), zap.String("path", c.Path()))
		return c.Next()
	}
}
