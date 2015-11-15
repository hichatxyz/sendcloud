package sendcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SendCloud struct {
	ApiUser  string
	ApiKey   string
	From     string
	Fromname string
}

func (s *SendCloud) TemplateSend(tos []string, subs map[string]interface{}, tpl_name string) {
	substitution_vars, err := json.Marshal(map[string]interface{}{
		"to":  tos,
		"sub": subs})
	fmt.Println(string(substitution_vars), err)

	resp, err := http.PostForm("http://sendcloud.sohu.com/webapi/mail.send_template.json", url.Values{
		"api_user":             {s.ApiUser},
		"api_key":              {s.ApiKey},
		"from":                 {s.From},
		"substitution_vars":    {string(substitution_vars)},
		"template_invoke_name": {tpl_name},
		"fromname":             {s.Fromname},
	})

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

}
