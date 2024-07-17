package email

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/cbroglie/mustache"
	"gopkg.in/gomail.v2"
)

var templatePath = "internal/templates/transaction_success.html"

func SendEmail(to, playerID, productName string, productPrice float32) error {
	template, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("could not read template: %v", err)
	}

	price := fmt.Sprintf("IDR %.2f", productPrice)

	expiredAt := time.Now().AddDate(0, 0, 1).Format("02-01-2006")

	data := map[string]string{
		"player_id": playerID,
		"product":   productName,
		"price":     price,
		"expiredAt": expiredAt,
	}

	body, err := mustache.Render(string(template), data)
	if err != nil {
		return fmt.Errorf("could not render template: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "boedi@microservices.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Payment Pending - Go Store")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, "27de36f4ba6439", "1752eeafc75a03")

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %v", err)
	}

	log.Println("Email sent successfully")
	return nil
}
