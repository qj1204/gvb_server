package email

import (
	"gopkg.in/gomail.v2"
	"gvb_server/global"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	e := global.Config.Email
	return sendEmail(e.User, e.Password, e.Host, e.Port, name, e.DefaultFromEmail, string(a.Subject), body)
}

func NewCode() Api {
	return Api{Subject: Code}
}

func NewNote() Api {
	return Api{Subject: Note}
}

func NewAlarm() Api {
	return Api{Subject: Alarm}
}

func sendEmail(userName, authCode, host string, port int, mailTo, sendName, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	return d.DialAndSend(m)
}
