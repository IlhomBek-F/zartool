package controller

import (
	"net/http"
	"os"
	"strconv"
	_ "zartool/docs"
	"zartool/domain"
	"zartool/internal"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase    domain.LoginOwnerUsecase
	OwnerRepository domain.OwnerRepository
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
func (lc LoginController) Login(e echo.Context) error {
	ownerCredential := new(domain.OwnerPayload)

	if err := e.Bind(&ownerCredential); err != nil {
		errorCode, message := internal.GetErrorCode(err)
		return e.JSON(errorCode, domain.ErrorResponse{Status: errorCode, Message: message})
	}

	if err := e.Validate(ownerCredential); err != nil {
		errorCode, message := internal.GetErrorCode(err)
		return e.JSON(errorCode, domain.ErrorResponse{Status: errorCode, Message: message})
	}

	owner, err := lc.LoginUsecase.GetOwnerByLogin(ownerCredential.Login)

	if err != nil {
		errorCode, message := internal.GetErrorCode(err)
		return e.JSON(errorCode, domain.ErrorResponse{Status: errorCode, Message: message})
	}

	if bcrypt.CompareHashAndPassword([]byte(owner.Password), []byte(ownerCredential.Password)) != nil {
		return e.JSON(http.StatusUnauthorized, domain.ErrorResponse{Status: http.StatusUnauthorized, Message: "Invalid credentials"})
	}

	jwtPrivateKey := os.Getenv("ACCESS_TOKEN_SECRET")
	accessTokenExpiryHour, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	accessToken, err := lc.LoginUsecase.GeneretaAccessToken(owner, jwtPrivateKey, accessTokenExpiryHour)

	if err != nil {
		errorCode, message := internal.GetErrorCode(err)
		return e.JSON(errorCode, domain.ErrorResponse{Status: errorCode, Message: message})
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
