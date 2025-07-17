package controller

import (
	"net/http"
	"net/url"
	"strconv"
	"zartool/models"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
)

// Create new rental godoc
//
//	@Summary        Create new rental
//	@Description    Create new rental
//	@Tags           zartool
//	@Accept         json
//	@Security       JWT
//	@Produce        json
//	@Param          payload  body models.User true "Create new rental"
//	@Router         /rental/create [post]
func (c Controller) CreateNewRental(e echo.Context) error {
	var newRental models.User

	if err := e.Bind(&newRental); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	err := repositories.CreateNewRental(c.DB, newRental)

	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	resp := models.SuccessResponse[models.User]{
		Status:  http.StatusCreated,
		Message: "Success",
	}

	return e.JSON(http.StatusCreated, resp)
}

// Update rental godoc
//
//	@Summary        Update rental
//	@Description    Update rental
//	@Tags           zartool
//	@Accept         json
//	@Produce        json
//	@Security       JWT
//	@Param          payload  body models.User  true    "Update rental"
//	@Router         /rental/update [put]
func (c Controller) UpdateRental(e echo.Context) error {
	var currentRental models.User

	if err := e.Bind(&currentRental); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	err := repositories.UpdateRental(c.DB, currentRental)

	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	resp := models.SuccessResponse[models.User]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    currentRental,
	}

	return e.JSON(http.StatusOK, resp)
}

// Delete rental godoc
//
//	@Summary        Delete rental
//	@Description    Delete rental
//	@Tags           zartool
//	@Accept         json
//	@Security       JWT
//	@Produce        json
//	@Param          id  path  int  true "rental id"
//	@Router         /rental/delete/{id} [delete]
func (c Controller) DeleteRental(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	if err := repositories.DeleteRental(c.DB, uint(id)); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal Server error"})
	}

	resp := models.SuccessResponse[models.User]{
		Status:  http.StatusOK,
		Message: "Success",
	}

	return e.JSON(http.StatusOK, resp)
}

// Complete rental godoc
//
//	@Summary        Complete rental
//	@Description    Complete rental
//	@Tags           zartool
//	@Accept         json
//	@Security       JWT
//	@Produce        json
//	@Param          id  path  int  true "rental id"
//	@Router         /rental/complete/{id} [post]
func (c Controller) CompleteRental(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	if err := repositories.CompleteRental(c.DB, uint(id)); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := models.SuccessResponse[models.User]{
		Status:  http.StatusOK,
		Message: "Success",
	}

	return e.JSON(http.StatusOK, resp)
}

// Rental report godoc
//
//	@Summary        Rental report
//	@Description    Rental  report
//	@Tags           zartool
//	@Accept         json
//	@Security       JWT
//	@Produce        json
//	@Param          page  query  int false "page"
//	@Param          page  query  int false "page_size"
//	@Success        200 {object} models.RentalsResponse
//	@Router         /rental/report [get]
func (c Controller) GetRentalReport(e echo.Context) error {
	var queryMap url.Values = e.QueryParams()

	page, _ := strconv.Atoi(queryMap.Get("page"))
	pageSize, _ := strconv.Atoi(queryMap.Get("page_size"))
	queryTerm := queryMap.Get("q")

	reportData, meta, err := repositories.GetRentalReport(c.DB, page, pageSize, queryTerm)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := models.SuccessResponse[models.RentalReport]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    reportData,
		Meta:    meta,
	}

	return e.JSON(http.StatusOK, resp)
}

// Get rental list godoc
//
//	@Summary        Get rental list
//	@Description    Get rental list
//	@Tags           zartool
//	@Accept         json
//	@Security       JWT
//	@Produce        json
//	@Param          page  query  int false "page"
//	@Param          page_size  query  int false "page_size"
//	@Success        200 {object} models.RentalsResponse
//	@Router         /rentals [get]
func (c Controller) GetRentals(e echo.Context) error {
	var queryMap url.Values = e.QueryParams()

	page, _ := strconv.Atoi(queryMap.Get("page"))
	pageSize, _ := strconv.Atoi(queryMap.Get("page_size"))
	queryTerm := queryMap.Get("q")

	rentals, metaData, err := repositories.GetRentals(c.DB, page, pageSize, queryTerm)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := models.SuccessResponse[[]models.User]{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    rentals,
		Meta:    metaData,
	}

	return e.JSON(http.StatusOK, resp)
}
