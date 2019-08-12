package service

import (
	entity2 "github.com/F3r-n4nd0/WallStreetFirm/src/businesslogic/entity"
	"github.com/F3r-n4nd0/WallStreetFirm/src/report"
	"github.com/F3r-n4nd0/WallStreetFirm/src/report/entity"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type customerOneStockChangeService struct {
	Channel    *amqp.Channel
	RoutingKey string
}

func NewCustomerOneStockChangeService(channel *amqp.Channel, routingKey string) report.StockChangeService {
	return &customerOneStockChangeService{
		Channel:    channel,
		RoutingKey: routingKey,
	}
}

func (s *customerOneStockChangeService) NotifyChange(stockChange entity2.StockChange) error {

	message := entity.CustomerOneMessage{
		UUID:  stockChange.UUID,
		Value: stockChange.NewValue,
	}
	logrus.Info("Publish message %s", stockChange.UUID)
	data, err := message.Marshal()
	if err != nil {
		return err
	}
	return s.sendMessage(data)

}

func (s *customerOneStockChangeService) sendMessage(body []byte) error {
	err := s.Channel.Publish(
		"",
		s.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}
	return nil

}
