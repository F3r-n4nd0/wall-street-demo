package usecase

import (
	"time"

	entity2 "github.com/F3r-n4nd0/WallStreetFirm/src/businesslogic/entity"
	"github.com/F3r-n4nd0/WallStreetFirm/src/report"
)

type usecase struct {
	reportService  report.StockChangeService
	contextTimeout time.Duration
}

func NewReportUseCase(service report.StockChangeService, timeout time.Duration) report.UseCase {
	return &usecase{
		reportService:  service,
		contextTimeout: timeout,
	}
}

func (uc *usecase) Fetch(stockChange entity2.StockChange) error {
	err := uc.reportService.NotifyChange(stockChange)
	if err != nil {
		return err
	}
	return nil
}
