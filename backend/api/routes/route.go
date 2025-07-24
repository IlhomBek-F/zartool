package routes

import (
	"context"
	"net/http"
	"os"
	"time"
	"zartool/api/controller"
	"zartool/domain"

	"github.com/go-playground/validator"
	echoSwagger "github.com/swaggo/echo-swagger"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type Server struct {
	Port int
	DB   *gorm.DB
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(configureCORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(10))))

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "custom timeout error message returns to client",
		Timeout:      2 * time.Second,
	}))

	e.Validator = &domain.CustomValidator{Validator: validator.New()}
	secretKey := os.Getenv("ACCESS_TOKEN_SECRET")

	publicRoute := e.Group("/api")
	protectedRoute := e.Group("/api")

	server := controller.Controller{
		DB: *s.DB,
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	protectedRoute.Use(echojwt.JWT([]byte(secretKey)))

	{
		publicRoute.POST("/create-owner", server.CreateOwner)
		publicRoute.POST("/auth/login", server.Login)
	}

	{
		protectedRoute.POST("/rental/create", server.CreateNewRental)
		protectedRoute.PATCH("/rental/update", server.UpdateRental)
		protectedRoute.DELETE("/rental/delete/:id", server.DeleteRental)
		protectedRoute.POST("/rental/complete/:id", server.CompleteRental)
		protectedRoute.GET("/rental/report", server.GetRentalReport)
		protectedRoute.GET("/rentals", server.GetRentals)
	}

	{
		protectedRoute.GET("/warehouse-tools", server.GetWareHouseTools)
		protectedRoute.PUT("/warehouse-tool/update", server.UpdateWareHouseTool)
		protectedRoute.DELETE("/warehouse-tool/delete/:id", server.DeleteWarehouseTool)
		protectedRoute.POST("/warehouse-tool/create", server.AddNewTools)
	}

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

// SetRequestContextWithTimeout will set the request context with timeout for every incoming HTTP Request
func setRequestContextWithTimeout(d time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, cancel := context.WithTimeout(c.Request().Context(), d)
			defer cancel()

			newRequest := c.Request().WithContext(ctx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}
