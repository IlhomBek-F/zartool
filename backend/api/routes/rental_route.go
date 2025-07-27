package routes

import (
	"zartool/api/controller"
	"zartool/repositories"
	"zartool/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRentalRoute(db gorm.DB, groupeRoute echo.Group) {
	rentalRepo := repositories.NewRentalRepository(db)
	rentalController := controller.RentalController{
		RentalUsecase:    usecase.NewRentalUsecase(rentalRepo),
		RentalRepository: rentalRepo,
	}

	groupeRoute.POST("/rental/create", rentalController.CreateNewRental)
	groupeRoute.PATCH("/rental/update", rentalController.UpdateRental)
	groupeRoute.DELETE("/rental/delete/:id", rentalController.DeleteRental)
	groupeRoute.POST("/rental/complete/:id", rentalController.CompleteRental)
	groupeRoute.GET("/rental/report", rentalController.GetRentalReport)
	groupeRoute.GET("/rentals", rentalController.GetRentals)
}
