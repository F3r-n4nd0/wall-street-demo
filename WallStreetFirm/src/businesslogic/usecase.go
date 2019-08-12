package businesslogic

import "github.com/F3r-n4nd0/WallStreetFirm/src/businesslogic/entity"

type UseCase interface {
	UpdateStockChange(stockChange entity.StockChange) error
}
