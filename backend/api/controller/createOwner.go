package controller

import (
	"net/http"
	"zartool/domain"
	"zartool/internal"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type OwnerController struct {
	CreateOwnerUsecase domain.CreateOwnerUsecase
	OwnerRepository    domain.OwnerRepository
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
func (oc *OwnerController) CreateOwner(e echo.Context) error {
	newOwnerPayload := new(domain.OwnerPayload)
	var createdOwner domain.Owner

	if err := e.Bind(&newOwnerPayload); err != nil {
		errorCode, message := internal.GetErrorCode(err)
		return e.JSON(errorCode, domain.ErrorResponse{Status: errorCode, Message: message})
	}

	if err := e.Validate(newOwnerPayload); err != nil {
		errorCode, message := internal.GetErrorCode(err)
		return e.JSON(errorCode, domain.ErrorResponse{Status: errorCode, Message: message})
	}

	_, err := oc.CreateOwnerUsecase.GetOwnerByLogin(newOwnerPayload.Login)

	if err == nil {
		errorCode, message := internal.GetErrorCode(err)
		return e.JSON(errorCode, domain.ErrorResponse{Status: errorCode, Message: message})
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(newOwnerPayload.Password), bcrypt.DefaultCost)
	createdOwner.Password = string(encryptPassword)
	createdOwner.Login = newOwnerPayload.Login
	if err != nil {
		errorCode, message := internal.GetErrorCode(err)
		return e.JSON(errorCode, domain.ErrorResponse{Status: errorCode, Message: message})
	}

	if err := oc.CreateOwnerUsecase.CreateOwner(createdOwner); err != nil {
		errorCode, message := internal.GetErrorCode(err)
		return e.JSON(errorCode, domain.ErrorResponse{Status: errorCode, Message: message})
	}

	resp := domain.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
	}

	return e.JSON(http.StatusCreated, resp)
}
