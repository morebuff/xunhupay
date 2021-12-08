package xunhupay

import "fmt"

// Pay 支付示例
func Pay() {

	appId := "xxxxxxxxxxxxxxxxxxx"                        //Appid
	appSecret := "YYYYYYYYYYYYYYYYYYYYYYY"                //密钥
	var host = "https://api.xunhupay.com/payment/do.html" //跳转支付页接口URL

	client := NewHuPi(&appId, &appSecret) //初始化调用

	//支付参数，appid、time、nonce_str和hash这四个参数不用传，调用的时候执行方法内部已经处理
	params := map[string]string{
		"version":        "1.1",
		"trade_order_id": "123456789",
		"total_fee":      "0.1",
		"title":          "测试标题",
		"notify_url":     "http://xxxxxxx.com",
		"return_url":     "http://xxxx.com",
		"wap_name":       "XXX店铺",
		"callback_url":   "",
	}

	execute, err := client.Execute(host, params) //执行支付操作
	if err != nil {
		panic(err)
	}
	fmt.Println(execute) //打印支付结果
}

// Query 查询示例
func Query() {

	appId := "xxxxxxxxxxxxxxxxxxx"                           //Appid
	appSecret := "YYYYYYYYYYYYYYYYYYYYYYY"                   //密钥
	var host = "https://api.xunhupay.com/payment/query.html" //查询接口URL

	client := NewHuPi(&appId, &appSecret) //初始化调用

	//查询参数，appid、time、nonce_str和hash这四个参数不用传，调用的时候执行方法内部已经处理
	params := map[string]string{
		"out_trade_order": "52c0194467c459082e61e56fccd3ece7",
		//"open_order_id":"123456789",
	}

	execute, err := client.Execute(host, params) //执行查询操作

	if err != nil {
		panic(err)
	}
	fmt.Println(execute) //打印查询结果
}
