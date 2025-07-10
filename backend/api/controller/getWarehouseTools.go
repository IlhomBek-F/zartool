package controller

import (
	"net/http"
	"zartool/models"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
)

func (s *Controller) GetWareHouseTools(e echo.Context) error {
	tools, err := repositories.GetWareHouseTools(s.DB)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    tools,
	}

	return e.JSON(http.StatusOK, resp)
}
