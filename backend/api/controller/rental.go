package controller

import (
	"net/http"
	"net/url"
	"strconv"
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

func (c Controller) DeleteRental(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	if err := repositories.DeleteRental(c.DB, uint(id)); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal Server error"})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
	}

	return e.JSON(http.StatusOK, resp)
}

func (c Controller) CompleteRental(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	if err := repositories.CompleteRental(c.DB, uint(id)); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Meta:    models.MetaModel{},
	}

	return e.JSON(http.StatusOK, resp)
}

func (c Controller) GetRentalReport(e echo.Context) error {
	var queryMap url.Values = e.QueryParams()

	page, _ := strconv.Atoi(queryMap.Get("page"))
	pageSize, _ := strconv.Atoi(queryMap.Get("page_size"))
	queryTerm := queryMap.Get("q")

	reportData, meta, err := repositories.GetRentalReport(c.DB, page, pageSize, queryTerm)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    reportData,
		Meta:    meta,
	}

	return e.JSON(http.StatusOK, resp)
}

func (c Controller) GetRentals(e echo.Context) error {
	var queryMap url.Values = e.QueryParams()

	page, _ := strconv.Atoi(queryMap.Get("page"))
	pageSize, _ := strconv.Atoi(queryMap.Get("page_size"))
	queryTerm := queryMap.Get("q")

	rentals, metaData, err := repositories.GetRentals(c.DB, page, pageSize, queryTerm)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Interna; server error"})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    rentals,
		Meta:    metaData,
	}

	return e.JSON(http.StatusOK, resp)
}
