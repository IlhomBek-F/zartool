package domain

type WarehouseTools struct {
	Base
	Name string `json:"name" validate:"required"`
	Size string `json:"size" validate:"required"`
}

type WarehouseToolsResponse = SuccessResponseWithMeta[[]WarehouseTools]
type WarehouseToolsCreateResponse = SuccessResponseWithData[[]WarehouseTools]
type WarehouseToolsUpdateResponse = SuccessResponseWithData[WarehouseTools]
