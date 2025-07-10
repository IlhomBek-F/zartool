package controller

import (
	"net/http"
	"zartool/models"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
)

func (s Controller) AddNewTools(e echo.Context) error {
	var newTool []models.WarehouseTools

	if err := e.Bind(&newTool); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	err := repositories.AddNewTool(s.DB, newTool)

	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Succes",
		Data:    newTool,
	}

	return e.JSON(http.StatusCreated, resp)
}
