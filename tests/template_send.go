package main

import (
	"github.com/hichatxyz/sendcloud"
)

func main() {
	s := sendcloud.SendCloud{
		ApiUser:  "你的 api_user",
		ApiKey:   "你的 api_key",
		From:     "service@example.com",
		Fromname: "service",
	}

	tpl_name := "邮件模板调用名称"

	tos := []string{"XXX@example.com", "XXX@example.com"}
	subs := map[string]interface{}{
		//邮件模板调用名称
		"%name%":     []string{"XXX", "XXX"},
		"%authcode%": []string{"XXX", "XXXX"},
	}

	s.TemplateSend(tos, subs, tpl_name)

}
