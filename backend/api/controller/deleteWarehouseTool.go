package controller

import (
	"net/http"
	"strconv"
	"zartool/models"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
)

func (c Controller) DeleteWarehouseTool(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	err = repositories.DeleteWarehouseTool(c.DB, id)

	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
	}

	return e.JSON(http.StatusOK, resp)
}
