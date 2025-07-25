package routes

import (
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

	server := controller.LoginController{
		DB: *s.DB,
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	protectedRoute.Use(echojwt.JWT([]byte(secretKey)))

	NewLoginRoute(*s.DB, *publicRoute)
	NewOwnerRoute(*s.DB, *protectedRoute)
	NewRentalRoute(*s.DB, *protectedRoute)

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
