package controller

import (
	"net/http"
	"zartool/models"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
)

func (c Controller) UpdateWareHouseTool(e echo.Context) error {
	var tool models.WarehouseTools

	if err := e.Bind(&tool); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	err := repositories.UpdateWareHouseTool(c.DB, tool)

	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    tool,
	}

	return e.JSON(http.StatusOK, resp)
}
