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
//	@Summary        Create owner
//	@Description    Create new owner
//	@Tags           zartool
//	@Accept         json
//	@Produce        json
//	@Security       JWT
//	@Param          owner  body models.Owners  true    "Owner payload"
//	@Router         /create-owner [post]
func (s *Controller) CreateOwner(e echo.Context) error {
	var newOwner models.Owners

	if err := e.Bind(&newOwner); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	_, err := repositories.GetOwnerByLogin(s.DB, newOwner.Login)

	if err == nil {
		return e.JSON(http.StatusFound, models.ErrorResponse{Status: http.StatusFound, Message: "owner exist with this login"})
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(newOwner.Password), bcrypt.DefaultCost)
	newOwner.Password = string(encryptPassword)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	if err := repositories.CreateOwner(s.DB, newOwner); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := models.SuccessResponse[models.Owners]{
		Status:  http.StatusCreated,
		Message: "Success",
	}

	return e.JSON(http.StatusCreated, resp)
}
