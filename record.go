package alidns

import (
	"strings"
)

// BaseRecordResponse 定义返回结构体
type BaseRecordResponse struct {
	RecordId  string
	RequestID string
}

// RecordInfo 结构体
type RecordInfo struct {
	DomainName string `json:"DomainName"`
	Line       string `json:"Line"`
	Locked     bool   `json:"Locked"`
	Priority   int    `json:"Priority"`
	RR         string `json:"RR"`
	RecordId   string `json:"RecordId"`
	Status     string `json:"Status"`
	TTL        int    `json:"TTL"`
	Type       string `json:"Type"`
	Value      string `json:"Value"`
	RequestId  string `json:"RequestId"`
}

// DescribeRecordResponse 返回结构体
var DescribeRecordResponse struct {
	DomainRecords struct {
		Record []RecordInfo `json:"Record"`
	} `json:"DomainRecords"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	RequestID  string `json:"RequestId"`
	TotalCount int    `json:"TotalCount"`
}

// AddDomainRecord Add Domain Record
func (cli *Client) AddDomainRecord(domain string, RR string, Type string, Value string, optional map[string]string) (BaseRecordResponse, ErrorResponse, error) {
	// 设置必要参数
	body := map[string]string{
		"DomainName": domain,
		"RR":         RR,
		"Type":       strings.ToUpper(Type),
		"Value":      Value,
	}

	var respInfo BaseRecordResponse
	// 请求
	errResp, err := cli.Do("AddDomainRecord", body, optional, &respInfo)
	return respInfo, errResp, err

}

// DeleteDomainRecord 调用DeleteDomainRecord根据传入参数删除解析记录
// https://help.aliyun.com/document_detail/29773.html?spm=a2c4g.11186623.6.642.701b2846ThgD9h
func (cli *Client) DeleteDomainRecord(RecordId string) (respInfo BaseRecordResponse, errResp ErrorResponse, err error) {
	// 设置必要参数
	body := map[string]string{
		"RecordId": RecordId,
	}

	// var respInfo BaseRecordResponse
	errResp, err = cli.Do("DeleteDomainRecord", body, nil, &respInfo)

	return respInfo, errResp, nil
}

// UpdateDomainRecord 调用UpdateDomainRecord根据传入参数修改解析记录。
// https://help.aliyun.com/document_detail/29774.html?spm=a2c4g.11186623.6.643.1f743192JMf3pj
func (cli *Client) UpdateDomainRecord(RR string, RecordId string, Type string, Value string, optional map[string]string) (respInfo BaseRecordResponse, errResp ErrorResponse, err error) {

	body := map[string]string{
		"RR":       RR,
		"RecordId": RecordId,
		"Type":     strings.ToUpper(Type),
		"Value":    Value,
	}

	errResp, err = cli.Do("UpdateDomainRecord", body, optional, &respInfo)
	return respInfo, errResp, err
}

// DescribeDomainRecords 调用DescribeDomainRecords根据传入参数获取指定主域名的所有解析记录列表。
// https://help.aliyun.com/document_detail/29776.html?spm=a2c4g.11186623.6.638.f7553b59curplN
func (cli *Client) DescribeDomainRecords(domain string, optional map[string]string) (respInfo BaseRecordResponse, errResp ErrorResponse, err error) {
	body := map[string]string{
		"DomainName": domain,
	}

	errResp, err = cli.Do("DescribeDomainRecords", body, optional, &respInfo)
	return respInfo, errResp, err
}

// DescribeDomainRecordInfo
// https://help.aliyun.com/document_detail/29777.html?spm=a2c4g.11186623.6.639.31795eb4kuGJJO
func (cli *Client) DescribeDomainRecordInfo(RecordId string) (respInfo RecordInfo, errResp ErrorResponse, err error) {
	body := map[string]string{
		"RecordId": RecordId,
	}

	errResp, err = cli.Do("DescribeDomainRecordInfo", body, &respInfo)

	return respInfo, errResp, err
}
