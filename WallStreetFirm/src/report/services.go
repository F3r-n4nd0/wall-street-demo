package report

import (
	entity2 "github.com/F3r-n4nd0/WallStreetFirm/src/businesslogic/entity"
)

type StockChangeService interface {
	NotifyChange(stockChange entity2.StockChange) error
}
