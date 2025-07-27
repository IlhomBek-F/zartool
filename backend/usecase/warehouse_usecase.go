package usecase

import "zartool/domain"

type warehouseUsecase struct {
	warehouseRepository domain.WarehouseRepository
}

func NewWarehouseUsecase(wr domain.WarehouseRepository) domain.WarehouseUsecase {
	return warehouseUsecase{warehouseRepository: wr}
}

func (wu warehouseUsecase) AddNewTool(tools *[]domain.WarehouseTools) error {
	return wu.warehouseRepository.AddNewTool(tools)
}

func (wu warehouseUsecase) GetWareHouseTools(page int, pageSize int) ([]domain.WarehouseTools, domain.MetaModel, error) {
	return wu.warehouseRepository.GetWareHouseTools(page, pageSize)
}

func (wu warehouseUsecase) UpdateWareHouseTool(tool *domain.WarehouseTools) error {
	return wu.warehouseRepository.UpdateWareHouseTool(tool)
}

func (wu warehouseUsecase) DeleteWarehouseTool(id int) error {
	return wu.warehouseRepository.DeleteWarehouseTool(id)
}
