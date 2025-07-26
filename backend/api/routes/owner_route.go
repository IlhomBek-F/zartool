package routes

import (
	"zartool/api/controller"
	"zartool/repositories"
	"zartool/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewOwnerRoute(db gorm.DB, groupeRoute echo.Group) {
	or := repositories.NewOwnerRepository(db)
	ownerController := controller.OwnerController{
		CreateOwnerUsecase: usecase.NewCreateOwnerusecase(or),
		OwnerRepository:    or,
	}

	groupeRoute.POST("/create-owner", ownerController.CreateOwner)
}
