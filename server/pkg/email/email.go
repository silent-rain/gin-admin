package email

import (
	"strings"

	"gopkg.in/gomail.v2"
)

// Options 邮箱参数
type Options struct {
	Host     string `json:"host"`     // 服务地址
	Port     int    `json:"port"`     // 服务端口
	Username string `json:"username"` // 发件人
	Alias    string `json:"alias"`    // 发件人别名
	Password string `json:"password"` // 发件人密码或授权码
	To       string `json:"to"`       // 收件人 多个用;分割
	Cc       string `json:"cc"`       // 抄送
	Subject  string `json:"subject"`  // 邮件主题
	Body     string `json:"body"`     // 邮件内容
}

// Send 发送邮件
func Send(o *Options) error {
	m := gomail.NewMessage()

	// 设置发件人
	m.SetHeader("From", m.FormatAddress(o.Username, o.Alias))

	// 设置发送给多个用户
	mailArrTo := strings.Split(o.To, ";")
	m.SetHeader("To", mailArrTo...)

	// 设置邮件主题
	m.SetHeader("Subject", o.Subject)

	// 设置邮件正文
	m.SetBody("text/html", o.Body)

	d := gomail.NewDialer(o.Host, o.Port, o.Username, o.Password)

	return d.DialAndSend(m)
}
