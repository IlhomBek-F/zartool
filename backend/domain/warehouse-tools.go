package domain

type WarehouseRepository interface {
	AddNewTool(tools *[]WarehouseTools) error
	GetWareHouseTools(page int, pageSize int) ([]WarehouseTools, MetaModel, error)
	UpdateWareHouseTool(tool *WarehouseTools) error
	DeleteWarehouseTool(id int) error
}

type WarehouseUsecase interface {
	AddNewTool(tools *[]WarehouseTools) error
	GetWareHouseTools(page int, pageSize int) ([]WarehouseTools, MetaModel, error)
	UpdateWareHouseTool(tool *WarehouseTools) error
	DeleteWarehouseTool(id int) error
}

type WarehouseTools struct {
	Base
	Name string `json:"name" validate:"required"`
	Size string `json:"size" validate:"required"`
}

type WarehouseToolsResponse = SuccessResponseWithMeta[[]WarehouseTools]
type WarehouseToolsCreateResponse = SuccessResponseWithData[[]WarehouseTools]
type WarehouseToolsUpdateResponse = SuccessResponseWithData[WarehouseTools]
