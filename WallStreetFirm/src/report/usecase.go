package report

import (
	"github.com/F3r-n4nd0/WallStreetFirm/src/businesslogic/entity"
)

type UseCase interface {
	Fetch(stockChange entity.StockChange) error
}
