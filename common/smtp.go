// Copyright 2020 The shop Authors

// Package common implements common.
package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"net/smtp"
	"shop/logger"
	"strconv"
	"strings"
	"time"
)

const (
	// EmailName ...
	EmailName = "学院"
	// EmailCodeSubject ...
	EmailCodeSubject = "验证码"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

type SmtpServer struct {
	Host string
	Port string
}

func EmailSendCode(nickname, to, code string) error {
	body := `
        <html>
        <body>
        <h3>
        ` + nickname + `您好：
        </h3>
        非常感谢您使用` + EmailName + `，您的验证码为：` + code + `此验证码有效期30分钟，请妥善保存。<br/>
        如果这不是您本人的操作，请忽略本邮件。<br/>
        </body>
        </html>
        `

	return SendToMail(to, EmailCodeSubject, body)
}

func (s *SmtpServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s<%s>\r\n", mail.Sender, mail.Sender)
	if len(mail.To) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	message += "Content-Type: text/html; charset=UTF-8"
	message += "\r\n\r\n" + mail.Body

	return message
}

func SendToMail(to, subject, body string) error {
	emailUser := beego.AppConfig.String("emailuser")
	emailPassword := beego.AppConfig.String("emailpassword")
	emailHost := beego.AppConfig.String("emailHost")
	emailPort := beego.AppConfig.String("emailPort")

	mail := Mail{}
	mail.Sender = emailUser
	mail.To = strings.Split(to, ";")
	mail.Subject = subject
	mail.Body = body

	messageBody := mail.BuildMessage()
	smtpServer := SmtpServer{
		Host: emailHost,
		Port: emailPort,
	}
	auth := smtp.PlainAuth("", mail.Sender, emailPassword, smtpServer.Host)
	err := smtp.SendMail(smtpServer.ServerName(), auth, mail.Sender, mail.To, []byte(messageBody))
	if err != nil {
		logger.Logger.Errorf("smtp send mail failed! err:[%v].", err)
		return err
	}
	return nil
}

func MakeCode() (code string) {
	code = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(899999) + 100000)
	return
}
