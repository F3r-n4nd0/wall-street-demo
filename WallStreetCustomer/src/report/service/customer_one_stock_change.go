package service

import (
	"log"

	"github.com/F3r-n4nd0/WallStreetCustomer/src/report"
	"github.com/F3r-n4nd0/WallStreetCustomer/src/report/entity"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type defaultWallStreetService struct {
	channel    *amqp.Channel
	routingKey string
}



func NewCustomerOneStockChangeService(channel *amqp.Channel, routingKey string) report.WallStreetService {
	return &defaultWallStreetService{
		channel:    channel,
		routingKey: routingKey,
	}
}

func (s *defaultWallStreetService) ConnectWallStreet() error {
	s.connectChannel()
	return nil
}


func (s *defaultWallStreetService) connectChannel() {
	msgs, err := s.channel.Consume(
		s.routingKey,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}
	go func() {
		for d := range msgs {
			message, err := entity.UnMarshal(d.Body)
			if err != nil {
				logrus.Error(err)
				continue
			}
			logrus.Info("Stock Change %s", message.UUID)
		}
	}()


}