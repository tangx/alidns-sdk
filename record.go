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

	// 请求
	err := cli.Do("AddDomainRecord", body, &respInfo)
	if err != nil {
		return 0, err
	}

	//fmt.Println(respInfo)
	rrID, _ := strconv.Atoi(respInfo.RecordID)
	return rrID, nil
}

// DeleteDomainRecord 调用DeleteDomainRecord根据传入参数删除解析记录
// https://help.aliyun.com/document_detail/29773.html?spm=a2c4g.11186623.6.642.701b2846ThgD9h
func (cli *Client) DeleteDomainRecord(RecordId int) (bool, error) {
	// 设置必要参数
	body := map[string]string{
		"RecordId": strconv.Itoa(RecordId),
	}

	var respInfo struct {
		RecordID  string
		RequestID string
	}

	err := cli.Do("DeleteDomainRecord", body, &respInfo)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateDomainRecord 调用UpdateDomainRecord根据传入参数修改解析记录。
// https://help.aliyun.com/document_detail/29774.html?spm=a2c4g.11186623.6.643.1f743192JMf3pj
func (cli *Client) UpdateDomainRecord(RR string, RecordId int, Type string, Value string, data map[string]string) (bool, error) {

	body := map[string]string{
		"RR":       RR,
		"RecordId": strconv.Itoa(RecordId),
		"Type":     strings.ToUpper(Type),
		"Value":    Value,
	}

	for k, v := range data {
		body[k] = v
	}
	var respInfo struct {
		RecordID  string
		RequestID string
	}

	err := cli.Do("UpdateDomainRecord", body, &respInfo)
	if err != nil {
		return false, err
	}

	return true, nil
}

// DescribeDomainRecords 调用DescribeDomainRecords根据传入参数获取指定主域名的所有解析记录列表。
// https://help.aliyun.com/document_detail/29776.html?spm=a2c4g.11186623.6.638.f7553b59curplN
func (cli *Client) DescribeDomainRecords(domain string, data map[string]string) (interface{}, error) {
	body := map[string]string{
		"DomainName": domain,
	}

	for k, v := range data {
		body[k] = v
	}

	var respInfo struct {
		DomainRecords struct {
			Record []struct {
				DomainName string `json:"DomainName"`
				Line       string `json:"Line"`
				Locked     bool   `json:"Locked"`
				Priority   int    `json:"Priority"`
				RR         string `json:"RR"`
				RecordID   string `json:"RecordId"`
				Status     string `json:"Status"`
				TTL        int    `json:"TTL"`
				Type       string `json:"Type"`
				Value      string `json:"Value"`
			} `json:"Record"`
		} `json:"DomainRecords"`
		PageNumber int    `json:"PageNumber"`
		PageSize   int    `json:"PageSize"`
		RequestID  string `json:"RequestId"`
		TotalCount int    `json:"TotalCount"`
	}

	err := cli.Do("DescribeDomainRecords", body, &respInfo)
	if err != nil {
		return nil, err
	}
	return respInfo, nil
}
