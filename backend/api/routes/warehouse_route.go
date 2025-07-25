package routes

import (
	"zartool/api/controller"
	"zartool/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewWarehouseRoute(db gorm.DB, groupRoute echo.Group) {
	wr := repositories.NewWarehouseRepository(db)
	wc := controller.WarehouseToolController{
		Db:                  db,
		WarehouseRepository: wr,
	}

	groupRoute.GET("/warehouse-tools", wc.GetWareHouseTools)
	groupRoute.PUT("/warehouse-tool/update", wc.UpdateWareHouseTool)
	groupRoute.DELETE("/warehouse-tool/delete/:id", wc.DeleteWarehouseTool)
	groupRoute.POST("/warehouse-tool/create", wc.AddNewTools)
}
