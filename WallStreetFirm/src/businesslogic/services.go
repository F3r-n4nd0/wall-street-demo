package businesslogic

import (
	entity2 "github.com/F3r-n4nd0/WallStreetFirm/src/businesslogic/entity"
)

type NotificationService interface {
	Fetch(stockChange entity2.StockChange) error
}
