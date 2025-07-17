package controller

import (
	"net/http"
	"os"
	"strconv"
	_ "zartool/docs"
	"zartool/internal"
	"zartool/models"
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
//	@Summary        Login
//	@Description    Login to app
//	@Tags           zartool
//	@Accept         json
//	@Produce        json
//	@Param          credential  body models.Owners  true    "Owner credential"
//	@Router         /auth/login [post]
func (s Controller) Login(e echo.Context) error {
	var ownerCredential models.Owners

	if err := e.Bind(&ownerCredential); err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	owner, err := repositories.GetOwnerByLogin(s.DB, ownerCredential.Login)

	if err != nil {
		return e.JSON(http.StatusNotFound, models.ErrorResponse{Status: http.StatusNotFound, Message: "Owner not found"})
	}

	if bcrypt.CompareHashAndPassword([]byte(owner.Password), []byte(ownerCredential.Password)) != nil {
		return e.JSON(http.StatusUnauthorized, models.ErrorResponse{Status: http.StatusUnauthorized, Message: "Invalid credentials"})
	}

	jwtPrivateKey := os.Getenv("ACCESS_TOKEN_SECRET")
	accessTokenExpiryHour, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	accessToken, err := internal.GeneretaAccessToken(owner, jwtPrivateKey, accessTokenExpiryHour)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}

	resp := models.SuccessResponse[any]{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]any{
			"id":           owner.ID,
			"created_at":   owner.CreatedAt,
			"access_token": accessToken,
		},
	}
	return e.JSON(http.StatusOK, resp)
}
