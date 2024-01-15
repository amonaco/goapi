package notification

import (
	"bytes"
	"html/template"
	"log"

	"github.com/amonaco/goapi/libs/config"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	fromName  = "Notifications"
	fromEmail = "notifications@logement3d.com"
)

func sendEmail(subject string, name string, email string, body string) {
	conf := config.Get()
	from := mail.NewEmail(fromName, fromEmail)
	to := mail.NewEmail(name, email)
	message := mail.NewSingleEmail(from, subject, to, body, body)

	client := sendgrid.NewSendClient(conf.Sendgrid)
	res, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("%+v\n", res)
	}
}

func userCreate(fields map[string]string) {

	const subject = "Setup Your Account"

	t, err := template.New("user_create").
		Parse(userCreateTemplate)
	if err != nil {
		log.Fatal(err)
	}

	data := bytes.Buffer{}
	err = t.Execute(&data, fields)
	if err != nil {
		log.Fatal(err)
	}

	sendEmail(subject, fields["Name"], fields["Email"], data.String())
}

func userForgotPassword(fields map[string]string) {

	const subject = "Reset Your Password"

	t, err := template.New("user_forgot_password").
		Parse(forgotPasswordTemplate)
	if err != nil {
		log.Fatal(err)
	}

	data := bytes.Buffer{}
	err = t.Execute(&data, fields)
	if err != nil {
		log.Fatal(err)
	}

	sendEmail(subject, fields["Name"], fields["Email"], data.String())
}
