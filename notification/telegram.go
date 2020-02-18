package notification

import (
	"github.com/BaronMsk/ssl-checker/config"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"fmt"
)

type CertificateInfoStruct struct {
	DNS []string
	Serial string
	NotAfter string
}

func NewNotification(config *config.ConfigurationStruct, infoStruct *CertificateInfoStruct)  {
	bot, err := tgbotapi.NewBotAPI(config.Notification.Telegram.Token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	msg := fmt.Sprintf("Certificate for domain %s is expired %s", infoStruct.DNS, infoStruct.NotAfter)
	u := tgbotapi.NewMessage(config.Notification.Telegram.ChatId, msg)
	bot.Send(u)

}
