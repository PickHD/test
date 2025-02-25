package middleware

import (
	"errors"
	"strings"
	"test/internal/config"
	"test/internal/helper"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWTMiddleware(ctx *fiber.Ctx) error {
	err := validate(ctx)
	if err != nil {
		return helper.NewResponses[any](ctx, fiber.StatusUnauthorized, "unauthorized", nil, err)
	}

	return ctx.Next()
}

func validate(ctx *fiber.Ctx) error {
	config := config.NewConfig()

	header := ctx.Get("Authorization", "")
	if !strings.Contains(header, "Bearer ") {
		return errors.New("token not found")
	}

	getToken := strings.Replace(header, "Bearer ", "", -1)
	parseToken, err := jwt.Parse(getToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt.Secret), nil
	})
	if err != nil {
		return errors.New("invalid token")
	}

	claims := parseToken.Claims.(jwt.MapClaims)

	if expInt, ok := claims["exp"].(float64); ok {
		now := time.Now().Unix()
		if now > int64(expInt) {
			return errors.New("expired token")
		}
	}

	return nil
}
