package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	// valid
	user = "tufei7818@126.com"
	password = "*******" // check cross-chain account
	host = "smtp.126.com"
	port = 25
	receipt = "fukun@onchain.com"
)

func main() {
	auth := smtp.PlainAuth("", user, password, host)
	//auth := LoginAuth(user, password)

	to := []string{receipt}
	msg := getTestMsg(to)

	serverAddr := fmt.Sprintf("%s:%d", host, port)
	if err := smtp.SendMail(serverAddr, auth, user, to, msg); err != nil {
		fmt.Printf("send mail error: %v", err)
	} else {
		fmt.Println("send email success")
	}
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}

func getTestMsg(to []string) []byte {
	nickname := "test"
	subject := "test mail"
	contentType := "Content-Type: text/plain; charset=UTF-8"
	body := "This is the email body."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	return msg
}
