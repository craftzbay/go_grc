package helpers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Response(statusCode uint, c *fiber.Ctx, res interface{}) error {
	return c.Status(fiber.StatusOK).JSON(res)
}

func ResponseForbidden(c *fiber.Ctx, message string) error {
	zap.L().Warn(fiber.ErrForbidden.Message, zap.String("error", message))
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": message})
}

func ResponseUnauthorized(c *fiber.Ctx, message string) error {
	zap.L().Warn(fiber.ErrUnauthorized.Message, zap.String("error", message))
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": message})
}

func ResponseBadRequest(c *fiber.Ctx, message string) error {
	zap.L().Warn(fiber.ErrBadRequest.Message, zap.String("error", message))
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": message})
}

func ResponseErr(c *fiber.Ctx, message string) error {
	zap.L().Warn(fiber.ErrInternalServerError.Message, zap.String("error", message))
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": message})
}
