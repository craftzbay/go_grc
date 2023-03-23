package helpers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Response(statusCode uint, c *fiber.Ctx, responseDetails ...interface{}) error {

	if len(responseDetails) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{})
	}

	return c.Status(fiber.StatusOK).JSON(responseDetails[0])
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
