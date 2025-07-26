package routes

import (
	"zartool/api/controller"
	"zartool/repositories"
	"zartool/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewLoginRoute(db gorm.DB, groupRoute echo.Group) {
	ownerRepo := repositories.NewOwnerRepository(db)
	loginController := controller.LoginController{
		LoginUsecase:    usecase.NewLoginUsecase(ownerRepo),
		OwnerRepository: ownerRepo,
	}

	groupRoute.POST("/auth/login", loginController.Login)
}
