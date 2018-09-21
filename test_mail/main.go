package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()

	m.SetHeader("From", "xxxxx@qq.com")
	m.SetHeader("To", "yyyyy@qq.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 587, "xxxxx@ qq.com", "okbnsnqptzjzfigd")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
