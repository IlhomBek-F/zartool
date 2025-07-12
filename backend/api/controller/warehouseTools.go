package controller

import (
	"net/http"
	"net/url"
	"strconv"
	"zartool/models"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
)

func (s *Controller) GetWareHouseTools(e echo.Context) error {
	var queries url.Values = e.QueryParams()

	page, _ := strconv.Atoi(queries.Get("page"))
	pageSize, _ := strconv.Atoi(queries.Get("page_size"))

	tools, meta, err := repositories.GetWareHouseTools(s.DB, page, pageSize)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    tools,
		Meta:    meta,
	}

	return e.JSON(http.StatusOK, resp)
}

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
