package controller

import (
	"net/http"
	"zartool/models"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Create owner godoc
//
//		@Summary        Create owner
//		@Description    Create new owner
//		@Tags           zartool
//		@Accept         json
//		@Produce        json
//		@Security       JWT
//		@Param          owner  body models.OwnerPayload true  "Owner payload"
//	 @Success        200 {object} models.SuccessResponse
//		@Router         /create-owner [post]
func (s *Controller) CreateOwner(e echo.Context) error {
	newOwnerPayload := new(models.OwnerPayload)
	var createdOwner models.Owner

	if err := e.Bind(&newOwnerPayload); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	if err := e.Validate(newOwnerPayload); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, models.ErrorResponse{Status: http.StatusUnprocessableEntity, Message: err.Error()})
	}

	_, err := repositories.GetOwnerByLogin(s.DB, newOwnerPayload.Login)

	if err == nil {
		return e.JSON(http.StatusFound, models.ErrorResponse{Status: http.StatusFound, Message: "owner exist with this login"})
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(newOwnerPayload.Password), bcrypt.DefaultCost)
	newOwnerPayload.Password = string(encryptPassword)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	if err := repositories.CreateOwner(s.DB, createdOwner); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := models.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
	}

	return e.JSON(http.StatusCreated, resp)
}
