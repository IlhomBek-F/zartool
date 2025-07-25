package controller

import (
	"net/http"
	"zartool/domain"
	"zartool/internal"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type OwnerController struct {
	Db              gorm.DB
	OwnerRepository domain.OwnerRepository
}

// Create owner godoc
//
//		@Summary        Create owner
//		@Description    Create new owner
//		@Tags           zartool
//		@Accept         json
//		@Produce        json
//		@Security       JWT
//		@Param          owner  body domain.OwnerPayload true  "Owner payload"
//	 @Success        200 {object} domain.SuccessResponse
//		@Router         /create-owner [post]
func (or *OwnerController) CreateOwner(e echo.Context) error {
	newOwnerPayload := new(domain.OwnerPayload)
	var createdOwner domain.Owner

	if err := e.Bind(&newOwnerPayload); err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "Internal server error"})
	}

	if err := e.Validate(newOwnerPayload); err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: err.Error()})
	}

	_, err := or.OwnerRepository.GetOwnerByLogin(newOwnerPayload.Login)

	if err == nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "owner exist with this login"})
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(newOwnerPayload.Password), bcrypt.DefaultCost)
	newOwnerPayload.Password = string(encryptPassword)

	if err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "Internal server error"})
	}

	if err := or.OwnerRepository.CreateOwner(createdOwner); err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "Internal server error"})
	}

	resp := domain.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
	}

	return e.JSON(http.StatusCreated, resp)
}
