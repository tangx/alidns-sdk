package alidns

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
)

const (
	//URI 云解析API请求的资源地址
	ALIDNS      = "alidns.aliyuncs.com"
	apiVersion  = "2015-01-09"
	sigVersion  = "1.0"
	reqProtocol = "https"
)

//基础的API响应结构
type BaseResponse struct {
	Code     int
	Message  string
	CodeDesc string
}

type RespError struct {
	Code      string `json:"Code,omitempty"`
	HostID    string `json:"HostId,omitempty"`
	Message   string `json:"Message,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
}

type RespInfo struct{}

// 请求 alidns API
func (cli *Client) request(method, action string, param url.Values, body io.Reader, respInfo interface{}) ([]byte, error) {

	if param == nil {
		param = url.Values{}
	}

	// 设置时区:
	//    https://blog.csdn.net/qq_26981997/article/details/53454606
	loc, _ := time.LoadLocation("") //参数就是解压文件的“目录”+“/”+“文件名”。
	//fmt.Println(time.Now().In(loc))
	//timeNow := time.Now().In(loc)
	//timeNow.Format("2006-01-02T15:04:05Z")
	timestamp := time.Now().In(loc).Format("2006-01-02T15:04:05Z")

	// 阿里云服务器时间使用的是 UTC 时区。 中国时区+8
	// 会一直报错: Specified time stamp or date value is expired
	param.Set("Timestamp", timestamp)

	// common body
	param.Set("Format", "JSON")
	param.Set("SignatureMethod", "HMAC-SHA1")
	param.Set("SignatureVersion", sigVersion)
	param.Set("Version", apiVersion)
	//param.Set("Timestamp", time.Now().Format("2016-01-02T15:04:05Z"))
	param.Set("SignatureNonce", uuid.New().String())

	param.Set("AccessKeyId", cli.SecretID)

	// ActionBody 请求是传入
	//param.Set("DomainName", "example.com")
	param.Set("Action", action)

	// 获取签名
	// 注意: 阿里云对用户 key 签名有特殊说明
	//    https://help.aliyun.com/document_detail/29747.html?spm=a2c4g.11186623.6.619.57ad2846HCScB1
	signature := Signature(method, param, cli.SecretKey+"&")

	// 请求体中增加签名参数
	//param.Set("Signature", url.QueryEscape(signature))
	param.Set("Signature", signature)

	//fmt.Println(signature)

	// 需要解决问题，signature 中的code转码之后多了一个 25 ，
	// 例如 / -> %2F (%252F) ,
	// 	   = -> %3D (%253D)
	// 问题原因:  param.Set("Signature", url.QueryEscape(signature))

	// 构建 url 请求地址
	reqURL := reqProtocol + "://" + ALIDNS + "/?" + param.Encode()
	req, err := http.NewRequest(method, reqURL, body)
	if err != nil {
		return nil, fmt.Errorf("构建请求错误: %s", err)
	}

	// 发起请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("执行请求错误: %s", err)
	}
	// 关闭请求
	defer resp.Body.Close()

	// 获取 body 内容
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return respBody, nil
}

// 使用 GET 方法请求 API
func (cli *Client) requestGET(action string, param url.Values, respInfo interface{}) ([]byte, error) {
	return cli.request("GET", action, param, nil, respInfo)
}

// Do to start requset
func (cli *Client) Do(action string, body map[string]string) ([]byte, error) {
	param := url.Values{}
	for k, v := range body {
		param.Set(k, v)
	}

	respInfo := RespInfo{}
	respBody, err := cli.requestGET(action, param, respInfo)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
