package alidns

import (
	"strings"
)

// BaseRecordResponse 定义返回结构体
type BaseRecordResponse struct {
	RecordId  string
	RequestID string
	Status    string
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

// AddDomainRecord 调用AddDomainRecord根据传入参数添加解析记录。
// https://help.aliyun.com/document_detail/29772.html?spm=a2c4g.11186623.6.641.1ee13b590wUMUS
func (cli *Client) AddDomainRecord(domain string, RR string, Type string, Value string, optional map[string]string) (respInfo BaseRecordResponse, errResp ErrorResponse, err error) {
	// 设置必要参数
	body := map[string]string{
		"DomainName": domain,
		"RR":         RR,
		"Type":       strings.ToUpper(Type),
		"Value":      Value,
	}

	errResp, err = cli.Do("AddDomainRecord", body, optional, &respInfo)
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

// DescribeDomainRecordInfo 调用DescribeDomainRecordInfo获取解析记录的详细信息。
// https://help.aliyun.com/document_detail/29777.html?spm=a2c4g.11186623.6.639.31795eb4kuGJJO
func (cli *Client) DescribeDomainRecordInfo(RecordId string) (respInfo RecordInfo, errResp ErrorResponse, err error) {
	body := map[string]string{
		"RecordId": RecordId,
	}

	errResp, err = cli.Do("DescribeDomainRecordInfo", body, nil, &respInfo)

	return respInfo, errResp, err
}

// SetDomainRecordStatus 调用SetDomainRecordStatus根据传入参数设置解析记录状态。
// https://help.aliyun.com/document_detail/29775.html?spm=a2c4g.11186623.6.645.334628465tIHev
func (cli *Client) SetDomainRecordStatus(RecordId string, status bool) (respInfo BaseRecordResponse, errResp ErrorResponse, err error) {

	var Status string
	if status {
		Status = "Disable"
	} else {
		Status = "Enable"
	}
	body := map[string]string{
		"RecordId": RecordId,
		"Status":   Status,
	}
	errResp, err = cli.Do("SetDomainRecordStatus", body, nil, &respInfo)
	return respInfo, errResp, err
}
