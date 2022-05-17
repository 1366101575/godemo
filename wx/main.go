package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ./notify -hook="https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=66"
func main() {
	var hook string

	flag.StringVar(&hook, "hook", "", "企业微信群机器人hook地址")

	flag.Parse()

	fmt.Println(hook)
	err := SendWx(hook, "先打个卡压压惊？")
	if err != nil {
		fmt.Println("send wx err:", err.Error())
	}
}

func SendWx(WxWebHook, content string) error {
	data := fmt.Sprintf("{\"msgtype\": \"text\",\"text\": {\"content\": \"%s\",\"mentioned_mobile_list\":[\"@all\"]}}", content)
	resp, err := http.Post(WxWebHook, "Content-Type: application/json", strings.NewReader(data))
	if err != nil {
		return err
	}
	//关闭连接
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("http close err:", err.Error())
		}
	}()
	//读取报文中所有内容
	body, err := ioutil.ReadAll(resp.Body)
	//输出内容
	fmt.Println(string(body))
	return nil
}
