package main

import (
	"log"

	eventconsumer "github.com/Phaseant/VK_Bot/internal/consumer/eventConsumer"
	"github.com/Phaseant/VK_Bot/internal/events/telegramEvents"
	"github.com/Phaseant/VK_Bot/internal/telegram"
	"github.com/spf13/viper"
)

const (
	HOST      = "api.telegram.org"
	batchSize = 100
)

func main() {
	tg := telegram.New(HOST, getApiToken())

	eventsProcessor := telegramEvents.New(tg)
	eventconsumer.New(eventsProcessor, eventsProcessor, batchSize).Start()

}

func getApiToken() string {
	viper.AddConfigPath("./configs") //get configs from configs folder
	viper.SetConfigName("config")
	viper.ReadInConfig()
	token := viper.GetString("TELEGRAM_API_TOKEN")

	if token == "" {
		log.Fatal("token is not provided")
	}

	return token
}
