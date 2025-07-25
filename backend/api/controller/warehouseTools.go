package controller

import (
	"net/http"
	"net/url"
	"strconv"
	"zartool/domain"
	"zartool/internal"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type WarehouseToolController struct {
	Db                  gorm.DB
	WarehouseRepository domain.WarehouseRepository
}

// GetWareHouseTools godoc
//
//		@Summary        GetWareHouseTools
//		@Description    GetWareHouseTools
//		@Tags           zartool
//		@Accept         json
//		@Produce        json
//		@Security       JWT
//		@Param          page query int false "page"
//		@Param          page_size query int false "page_size"
//	 @Success        200 {object} domain.WarehouseToolsResponse
//		@Router         /warehouse-tools [get]
func (wc WarehouseToolController) GetWareHouseTools(e echo.Context) error {
	var queries url.Values = e.QueryParams()

	page, _ := strconv.Atoi(queries.Get("page"))
	pageSize, _ := strconv.Atoi(queries.Get("page_size"))

	tools, meta, err := wc.WarehouseRepository.GetWareHouseTools(page, pageSize)

	if err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "Internal server error"})
	}

	resp := domain.WarehouseToolsResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    tools,
		Meta:    meta,
	}

	return e.JSON(http.StatusOK, resp)
}

// AddNewTools godoc
//
//		@Summary        AddNewTools
//		@Description    AddNewTools
//		@Tags           zartool
//		@Accept         json
//		@Produce        json
//		@Security       JWT
//		@Param          payload body domain.WarehouseTools false "body"
//	 @Success        200 {object} domain.WarehouseToolsCreateResponse
//		@Router         /warehouse-tool/create [post]
func (wc WarehouseToolController) AddNewTools(e echo.Context) error {
	var newTool = new([]domain.WarehouseTools)

	if err := e.Bind(&newTool); err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "Internal server error"})
	}

	for _, tool := range *newTool {
		if err := e.Validate(tool); err != nil {
			return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: err.Error()})
		}
	}

	err := wc.WarehouseRepository.AddNewTool(newTool)

	if err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: err.Error()})
	}

	resp := domain.WarehouseToolsCreateResponse{
		Status:  http.StatusCreated,
		Message: "Succes",
		Data:    *newTool,
	}

	return e.JSON(http.StatusCreated, resp)
}

// DeleteWarehouseTool godoc
//
//		@Summary        DeleteWarehouseTool
//		@Description    DeleteWarehouseTool
//		@Tags           zartool
//		@Accept         json
//		@Produce        json
//		@Security       JWT
//		@Param          id path int true "id"
//	 @Success        200 {object} domain.SuccessResponse
//		@Router         /warehouse-tool/delete/{id} [delete]
func (wc WarehouseToolController) DeleteWarehouseTool(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))

	if err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "Internal server error"})
	}

	err = wc.WarehouseRepository.DeleteWarehouseTool(id)

	if err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: err.Error()})
	}

	resp := domain.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success",
	}

	return e.JSON(http.StatusOK, resp)
}

// UpdateWareHouseTool godoc
//
//		@Summary        UpdateWareHouseTool
//		@Description    UpdateWareHouseTool
//		@Tags           zartool
//		@Accept         json
//		@Produce        json
//		@Security       JWT
//		@Param          payload body domain.WarehouseToolsUpdateResponse true "payload"
//	 @Success        200 {object} domain.SuccessResponse
//		@Router         /warehouse-tool/update/{id} [put]
func (wc WarehouseToolController) UpdateWareHouseTool(e echo.Context) error {
	var tool = new(domain.WarehouseTools)

	if err := e.Bind(&tool); err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: "Internal server error"})
	}

	if err := e.Validate(tool); err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: err.Error()})
	}

	err := wc.WarehouseRepository.UpdateWareHouseTool(tool)

	if err != nil {
		return e.JSON(internal.GetErrorCode(err), domain.ErrorResponse{Status: internal.GetErrorCode(err), Message: err.Error()})
	}

	resp := domain.WarehouseToolsUpdateResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    *tool,
	}

	return e.JSON(http.StatusOK, resp)
}
