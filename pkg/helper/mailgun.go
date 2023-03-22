package helper

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func SendEmail() {
	// Your available domain names can be found here:
	// (https://app.mailgun.com/app/domains)
	var yourDomain string = "your-domain-name" // e.g. mg.yourcompany.com

	// You can find the Private API Key in your Account Menu, under "Settings":
	// (https://app.mailgun.com/app/account/security)
	var privateAPIKey string = "your-private-key"

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	//When you have an EU-domain, you must specify the endpoint:
	//mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	sender := "sender@example.com"
	subject := "Fancy subject!"
	recipient := "recipient@example.com"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, "", recipient)
	body := `
		<html>
		<body>
			<h1>Sending HTML emails with Mailgun</h1>
			<p style="color:blue; font-size:30px;">Hello world</p>
			<p style="font-size:30px;">More examples can be found <a href="https://documentation.mailgun.com/en/latest/api-sending.html#examples">here</a></p>
		</body>
		</html>
	`
	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
