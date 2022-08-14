package infra

import (
	"github.com/demo/layer/adapter/presenter"
	"github.com/go-playground/validator/v10"
	"github.com/junyaU/evbus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func InitRouter(bus evbus.Bus) {
	popup := presenter.NewPopup(bus)

	e := echo.New()
	e.Validator = NewValidator()

	e.Use(middleware.Logger())

	e.POST("/popup", popup.Create)
	e.PATCH("popup/merchandise", popup.UpdateMerchandises)

	e.Logger.Fatal(e.Start(":8000"))
}
