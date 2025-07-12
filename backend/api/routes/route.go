package routes

import (
	"net/http"
	"os"
	"zartool/api/controller"

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
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(10))))
	e.Use(configureCORS())

	secretKey := os.Getenv("ACCESS_TOKEN_SECRET")

	publicRoute := e.Group("/api")
	protectedRoute := e.Group("/api")

	server := controller.Controller{
		DB: *s.DB,
	}

	protectedRoute.Use(echojwt.JWT([]byte(secretKey)))

	{
		publicRoute.POST("/create-owner", server.CreateOwner)
		publicRoute.POST("/auth/login", server.Login)
	}

	{
		protectedRoute.POST("/rental/create", server.CreateNewRental)
		protectedRoute.PUT("/rental/update", server.UpdateRental)
		protectedRoute.DELETE("/rental/delete/:id", server.DeleteRental)
		protectedRoute.POST("/rental/complete/:id", server.CompleteRental)
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
