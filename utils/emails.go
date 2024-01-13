package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func SendRequestEmail(toEmail string, code string) {
	var body bytes.Buffer
	template, err := template.ParseFiles("utils/template/email_template.html")
	template.Execute(&body, struct {
		Code string
	}{
		Code: code,
	})
	if err != nil {
		fmt.Println("error parsing files")
	}

	auth := smtp.PlainAuth(
		"",
		"dontreply.marica.id@gmail.com",
		"uxzs ktln dprk szxx",
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: ini adalah token untuk reset password anda " + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"dontreply.marica.id@gmail.com",
		[]string{toEmail},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}

}
