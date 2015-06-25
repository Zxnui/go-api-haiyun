package controllers

import (
	"github.com/go-martini/martini"
	"go-api-haiyun/conf"
	"go-api-haiyun/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func UploadAds(pushdata models.Pushdate) models.Returndata {
	urlstring, err := conf.Cfg.GetValue("url", upurl)
	if err != nil {
		log.Fatalf("配置文件解析失败:%s", err)
	}

	//将push_data数据解析为json格式
	json, err := models.GetJson(pushdata)
	if err != nil {
		log.Fatalf("json编码失败:%s", err)
	}

	//将json数据通过api地址发送
	client := &http.Client{}
	request, err := http.NewRequest("POST", urlstring, json)
	if err != nil {
		log.Fatalf("网络连接失败:%s", err)
	}

	//根据文档要求注明数据类型
	request.Header.Set("Content-Type", "application/json")
	res, err := client.Do(request)
	if err != nil {
		log.Fatalf("api地址连接失败:%s", err)
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		r, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("响应数据读取失败:%s", err)
			return nil
		}

	} else {
		return nil
	}
}
