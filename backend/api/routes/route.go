package routes

import (
	"net/http"
	"time"
	"zartool/domain"
	"zartool/internal/database"

	"github.com/go-playground/validator"
	echoSwagger "github.com/swaggo/echo-swagger"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

func RegisterRoutes(db gorm.DB, config database.Config) http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(configureCORS())
	e.Use(configureCORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(10))))

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "custom timeout error message returns to client",
		Timeout:      2 * time.Second,
	}))

	e.Validator = &domain.CustomValidator{Validator: validator.New()}

	publicRoute := e.Group("/api")
	protectedRoute := e.Group("/api")

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	protectedRoute.Use(echojwt.JWT([]byte(config.AccessTokenSecret)))

	NewLoginRoute(db, *publicRoute)
	NewOwnerRoute(db, *protectedRoute)
	NewRentalRoute(db, *protectedRoute)
	NewWarehouseRoute(db, *protectedRoute)

	return e
}

func configureCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           30,
	})
}
