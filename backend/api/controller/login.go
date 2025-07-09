package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller struct {
	DB gorm.DB
}

func (s Controller) Login(e echo.Context) error {
	return nil
}
