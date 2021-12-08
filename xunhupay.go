package xunhupay

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type HuPiClient struct {
	appId     *string //appId
	appSecret *string //appSecret
}

// NewHuPi 初始化客户端
func NewHuPi(appId, appSecret *string) *HuPiClient {
	return &HuPiClient{
		appId:     appId,
		appSecret: appSecret,
	}
}

// Execute 执行请求操作
func (client *HuPiClient) Execute(host string, params map[string]string) (string, error) {
	data := url.Values{}
	simple := strconv.FormatInt(time.Now().Unix(), 10)
	params["appid"] = *client.appId
	params["time"] = simple
	params["nonce_str"] = simple
	for k, v := range params {
		data.Add(k, v)
	}
	data.Add("hash", client.Sign(params))
	resp, err := http.PostForm(host, data)
	if err != nil {
		return "error", err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "error", err
	}
	return string(all), err
}

// Sign 签名方法
func (client *HuPiClient) Sign(params map[string]string) string {
	var data string
	keys := make([]string, 0, 0)
	params["appid"] = *client.appId
	for key, _ := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	//拼接
	for _, k := range keys {
		data = fmt.Sprintf("%s%s=%s&", data, k, params[k])
	}
	data = strings.Trim(data, "&")
	data = fmt.Sprintf("%s%s", data, *client.appSecret)
	m := md5.New()
	m.Write([]byte(data))
	sign := fmt.Sprintf("%x", m.Sum(nil))
	return sign
}
