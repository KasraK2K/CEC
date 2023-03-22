package helper

import (
	"context"
	"time"

	"github.com/mailgun/mailgun-go/v4"

	"app/pkg/config"
)

func SendEmail(recipients []string, body, subject string, bcc ...string) (string, string, error) {
	mg := mailgun.NewMailgun(
		config.AppConfig.MAILGUN_DOMAIN,
		config.AppConfig.MAILGUN_PRIVATE_API_KEY,
	)
	mg.SetAPIBase(config.AppConfig.MAILGUN_API_BASE)

	message := mg.NewMessage(config.AppConfig.MAILGUN_SENDER, subject, "", recipients...)
	for _, item := range bcc {
		message.AddBCC(item)
	}
	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		return "", "", err
	}

	return resp, id, nil
}
