package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// GetConfig 設定を取得する
func GetConfig() middleware.KeyAuthConfig {

	return middleware.KeyAuthConfig{
		Validator: getValidator(),
	}
}

// getValidator KeyAuthValidatorを取得する
func getValidator() middleware.KeyAuthValidator {
	return func(key string, c echo.Context) (bool, error) {

		return key == "208d97ff-0ab4-49b7-8eab-985ff27a11dc", nil
	}
}
