package controller

import (
	"net/http"
	"zartool/models"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
)

func (c Controller) CreateNewRental(e echo.Context) error {
	var newRental models.User

	if err := e.Bind(&newRental); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	err := repositories.CreateNewRental(c.DB, newRental)

	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
	}

	return e.JSON(http.StatusCreated, resp)
}

func (c Controller) UpdateRental(e echo.Context) error {
	var currentRental models.User

	if err := e.Bind(&currentRental); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	err := repositories.UpdateRental(c.DB, currentRental)

	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    currentRental,
	}

	return e.JSON(http.StatusOK, resp)
}

func (c Controller) GetRentals(e echo.Context) error {

	rentals, err := repositories.GetRentals(c.DB)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Interna; server error"})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    rentals,
	}

	return e.JSON(http.StatusOK, resp)
}
