package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/F3r-n4nd0/WallStreetFirm/src/businesslogic/usecase"
	"github.com/F3r-n4nd0/WallStreetFirm/src/report/service"
	usecase2 "github.com/F3r-n4nd0/WallStreetFirm/src/report/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {

	webAddr := os.Getenv("WEB_ADDRESS")
	rabbitMqHost := os.Getenv("RABBIT_MQ_HOST")
	rabbitMqPort := os.Getenv("RABBIT_MQ_PORT")

	rabbitMqUser := os.Getenv("RABBIT_MQ_USER")
	rabbitMqPassword := os.Getenv("RABBIT_MQ_PASSWORD")
	rabbitMqKeyQueueTradeForeCasting := os.Getenv("RABBIT_MQ_TFC")

	rabbitMqChannel := configureRabbitMq(rabbitMqHost, rabbitMqPort, rabbitMqUser, rabbitMqPassword)

	configureLog()
	engine := createGinHTTPDelivery()

	customerOneNotification := service.NewCustomerOneStockChangeService(rabbitMqChannel, rabbitMqKeyQueueTradeForeCasting)
	customerOneUseCase := usecase2.NewReportUseCase(customerOneNotification, 2)
	dummyBusinessLogic := usecase.NewDummyBusinessLogicUseCase(customerOneUseCase, 2)

	go dummyBusinessLogic.StartDummyTest()

	engine.Run(webAddr)

}

func configureRabbitMq(rabbitMqHost string, rabbitMqPort string, user string, password string) *amqp.Channel {
	connectionURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, rabbitMqHost, rabbitMqPort)
	conn, err := amqp.Dial(connectionURL)
	if err != nil {
		log.Panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}
	return ch
}

func createGinHTTPDelivery() *gin.Engine {

	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	config.AllowAllOrigins = true
	server.Use(cors.New(config))
	return server

}

func configureLog() {
	file, err := os.OpenFile("./wallStreetFirm.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
}
