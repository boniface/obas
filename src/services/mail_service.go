package services

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

func SendMail() {
	from := mail.NewEmail("Do_Not_Reply", "dont_reply@example.com")
	subject := "Thanks For Registering"
	to := mail.NewEmail("Example User", "test@example.com")
	plainTextContent := "Thanks For Registering"
	htmlContent := "<strong>Your Credentials </strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
