package tools

import (
	"gopkg.in/gomail.v2"
	"os"
)

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"` // 发件人邮箱
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default-from-email"` // 默认的发件人名字
	UseSSL           bool   `json:"use_ssl" yaml:"use_ssl"`                       // 是否使用ssl
	UserTls          bool   `json:"user_tls" yaml:"user_tls"`
}

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func NewEmail() Email {
	return Email{}
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}
func NewNote() Api {
	return Api{
		Subject: Note,
	}
}
func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}
// send 邮件发送  发给谁，主题，正文
func send(name, subject, body string) error {
	e := global.GlobalC.Email
	return sendMail(
		e.User,             // 发送人
		e.Password,         // 发送人的认证码
		e.Host,             // 对应邮箱厂商的host
		e.Port,             // 对应邮箱厂商的port
		name,               // 收件人
		e.DefaultFromEmail, // 默认的发件人名字
		subject,            // 发送标题
		body,               // 发送主题
	)
}

func sendMail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName)) // 谁发的,发送标题
	m.SetHeader("To", mailTo)                                // 发送给谁
	m.SetHeader("Subject", subject)                          // 主题
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	err := d.DialAndSend(m)
	return err
}
