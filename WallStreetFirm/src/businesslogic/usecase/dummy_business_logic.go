package usecase

import (
	"math/rand"
	"time"

	"github.com/F3r-n4nd0/WallStreetFirm/src/businesslogic"
	"github.com/F3r-n4nd0/WallStreetFirm/src/businesslogic/entity"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type dummyUseCase struct {
	notificationService businesslogic.NotificationService
	contextTimeout      time.Duration
}

func NewDummyBusinessLogicUseCase(service businesslogic.NotificationService, timeout time.Duration) *dummyUseCase {
	return &dummyUseCase{
		notificationService: service,
		contextTimeout:      timeout,
	}
}

func (uc *dummyUseCase) UpdateStockChange(stockChange entity.StockChange) error {
	err := uc.notificationService.Fetch(stockChange)
	if err != nil {
		return err
	}
	return nil
}

func (uc *dummyUseCase) StartDummyTest() {

	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(20)
		logrus.Info("Sleeping %d seconds...", n)
		time.Sleep(time.Duration(n) * time.Second)

		newUUID, err := uuid.NewUUID()
		if err != nil {
			logrus.Error(err)
			panic(err)
		}
		value := rand.Float64()
		stockChange := entity.StockChange{
			UUID:      newUUID.String(),
			NewValue:  value,
			OldValue:  value,
			StockUUID: newUUID.String(),
		}
		err = uc.UpdateStockChange(stockChange)
		if err != nil {
			logrus.Error(err)
			panic(err)
		}
	}

}
