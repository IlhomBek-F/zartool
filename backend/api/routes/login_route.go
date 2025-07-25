package routes

import (
	"zartool/api/controller"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewLoginRoute(db gorm.DB, groupRoute echo.Group) {
	ownerRepo := repositories.NewOwnerRepository(db)
	loginController := controller.LoginController{
		DB:              db,
		OwnerRepository: ownerRepo,
	}

	groupRoute.POST("/auth/login", loginController.Login)
}
