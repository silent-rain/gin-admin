package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	options := &Options{
		Host:     "smtp.163.com",
		Port:     465,
		Username: "xxx@163.com",
		Alias:    "gin-admin",
		Password: "", //密码或授权码
		To:       "",
		Cc:       "",
		Subject:  "subject",
		Body:     "body",
	}
	if err := Send(options); err != nil {
		t.Error("Mail Send error", err)
		return
	}
	t.Log("success")
}
