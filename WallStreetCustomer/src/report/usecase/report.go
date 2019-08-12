package usecase

import (
	"time"

	"github.com/F3r-n4nd0/WallStreetCustomer/src/report"
)

type usecase struct {
	wallStreetService  report.WallStreetService
	contextTimeout time.Duration
}



func NewReportUseCase(service report.WallStreetService, timeout time.Duration) report.UseCase {
	return &usecase{
		wallStreetService:  service,
		contextTimeout: timeout,
	}
}

func (uc *usecase) StartReceiveInformationWallStreet() error {
	err := uc.wallStreetService.ConnectWallStreet()
	if err != nil {
		return err
	}
	return nil
}