package alidns

import (
	"strconv"
	"strings"
)

// AddDomainRecord Add Domain Record
func (cli *Client) AddDomainRecord(domain string, RR string, Type string, Value string, data map[string]string) (int, error) {

	// 设置必要参数
	body := map[string]string{
		"DomainName": domain,
		"RR":         RR,
		"Type":       strings.ToUpper(Type),
		"Value":      Value,
	}

	// 补充可选参数
	for k, v := range data {
		body[k] = v
	}

	// 定义返回结构体
	var respInfo struct {
		RecordID  string
		RequestID string
	}
	err := cli.Do("AddDomainRecord", body, &respInfo)
	if err != nil {
		return 0, err
	}

	//fmt.Println(respInfo)
	rrID, err := strconv.Atoi(respInfo.RecordID)
	return rrID, nil
}
