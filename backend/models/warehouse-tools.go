package models

type WarehouseTools struct {
	Base
	Name string `json:"name"`
	Size string `json:"size"`
}

type WarehouseToolsResponse = SuccessResponseWithMeta[[]WarehouseTools]
type WarehouseToolsCreateResponse = SuccessResponseWithData[[]WarehouseTools]
type WarehouseToolsUpdateResponse = SuccessResponseWithData[WarehouseTools]
