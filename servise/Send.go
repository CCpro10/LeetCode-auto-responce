package servise

import (
	"bytes"
	"log"

	"encoding/json"
	"main/config"

	"io/ioutil"
	"net/http"
)

type message map[string]interface{}

func Send(content string) (bool, error) {
	//打卡地址
	Url := "https://leetcode-cn.com/graphql"

	data := "{\"operationName\":\"qaPublishAnswer\",\"variables\":{\"data\":{\"content\":\"" + content + "\",\"summary\":\"" + content + "\",\"mentionedUserSlugs\":[],\"thumbnail\":\"\",\"questionId\":\"qGB2zg\",\"postAnonymously\":false,\"replyTo\":\"yrH5N7\"}},\"query\":\"mutation qaPublishAnswer($data: QAPublishAnswerInput!) {\\n  qaPublishAnswer(data: $data) {\\n    answer {\\n      uuid\\n      parent {\\n        uuid\\n        __typename\\n      }\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n\"}"
	jsonStr := []byte(data)

	req, _ := http.NewRequest("POST", Url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Cookie", config.Config.Send.Cookie)
	req.Header.Set("origin", "https://leetcode-cn.com")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-timezone", "Asia/Shanghai")

	req.Header.Set("x-csrftoken", config.Config.Send.XCsrftoken)
	req.Header.Set("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36")
	req.Header.Set("x-definition-name", "qaPublishAnswer")
	req.Header.Set("x-operation-name", "qaPublishAnswer")

	response, e := http.DefaultClient.Do(req)
	if e != nil {
		log.Println(e)
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var mes message
	_ = json.Unmarshal(body, &mes)
	log.Println(mes)
	return true, nil

}
