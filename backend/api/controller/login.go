package controller

import (
	"net/http"
	"os"
	"strconv"
	_ "zartool/docs"
	"zartool/domain"
	"zartool/internal"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Controller struct {
	DB gorm.DB
}

// Login godoc
//
//		@Summary        Login
//		@Description    Login to app
//		@Tags           zartool
//		@Accept         json
//		@Produce        json
//		@Param          credential  body domain.OwnerPayload true  "Owner credential"
//	 @Success        200 {object} domain.OwnerResponse
//		@Router         /auth/login [post]
func (s Controller) Login(e echo.Context) error {
	ownerCredential := new(domain.OwnerPayload)

	if err := e.Bind(&ownerCredential); err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "Internal server error"})
	}

	if err := e.Validate(ownerCredential); err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: err.Error()})
	}

	owner, err := repositories.GetOwnerByLogin(s.DB, ownerCredential.Login)

	if err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "Owner not found"})
	}

	if bcrypt.CompareHashAndPassword([]byte(owner.Password), []byte(ownerCredential.Password)) != nil {
		return e.JSON(http.StatusUnauthorized, domain.ErrorResponse{Status: http.StatusUnauthorized, Message: "Invalid credentials"})
	}

	jwtPrivateKey := os.Getenv("ACCESS_TOKEN_SECRET")
	accessTokenExpiryHour, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	accessToken, err := internal.GeneretaAccessToken(owner, jwtPrivateKey, accessTokenExpiryHour)

	if err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "Internal server error"})
	}

	resp := domain.OwnerResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: domain.Credential{
			ID:          owner.ID,
			CreatedAt:   owner.CreatedAt,
			AccessToken: accessToken,
		},
	}
	return e.JSON(http.StatusOK, resp)
}
