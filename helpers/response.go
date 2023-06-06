package helpers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Success буцаагч - httpStatusCode 200
//
// Жишээ1: Response(c) гэж дуудвал {"message": "Амжилттай"}
//
// Жишээ2: Response(c, "Амжилттай бүртгэгдлээ") гэж дуудвал {"message": "Амжилттай бүртгэгдлээ"}
//
// Жишээ3: Response(c, anyJson) гэж дуудвал anyJson буцаана
func Response(c *fiber.Ctx, responseDetails ...interface{}) error {

	if len(responseDetails) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Амжилттай"})
	}

	var res interface{}
	switch val := responseDetails[0].(type) {
	case string:
		res = fiber.Map{"message": val}
	default:
		res = val
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// Forbidden хариу буцаагч - httpStatusCode 403
// {"message": "Хандах эрхгүй"}
func ResponseForbidden(c *fiber.Ctx) error {
	msg := "Хандах эрхгүй"
	zap.L().Error(fiber.ErrForbidden.Message, zap.String("error", msg))
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": msg})
}

// Unauthorized хариу буцаагч - httpStatusCode 401
// {"message": "Токен хугацаа дууссан эсвэл буруу байна"}
func ResponseUnauthorized(c *fiber.Ctx) error {
	msg := "Токен хугацаа дууссан эсвэл буруу байна"
	zap.L().Error(fiber.ErrUnauthorized.Message, zap.String("error", msg))
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": msg})
}

// BadRequest хариу буцаагч - httpStatusCode 400
func ResponseBadRequest(c *fiber.Ctx, message string) error {
	zap.L().Error(fiber.ErrBadRequest.Message, zap.String("error", message))
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": message})
}

// InternalServerError хариу буцаагч - httpStatusCode 500
func ResponseErr(c *fiber.Ctx, message string) error {
	zap.L().Error(fiber.ErrInternalServerError.Message, zap.String("error", message))
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": message})
}
