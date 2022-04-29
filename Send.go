package main

import (
	"main/config"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


func Send(content string)(bool,error){
	//打卡地址
	Url := "https://leetcode-cn.com/graphql"

	payload := url.Values{}
	payload.Set("inChina", "是")

	req, _ := http.NewRequest("POST", Url, strings.NewReader(payload.Encode()))
	req.Header.Set("Cookie",config.Config.Send.Cookie)
	req.Header.Add("Content-Type", "application/json")

	response, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var message Model.Message
	_ = json.Unmarshal(body, &message)

	return true,nil

	//如果成功,把打卡时间修改为当天
}
