//package alidns
//
//import (
//	"errors"
//	"fmt"
//	"net/url"
//	"strconv"
//	"strings"
//)
//
//type Record struct {
//	Action       string `json:"Action,omitempty"`
//	DomainName   string `json:"DomainName,omitempty"`
//	Line         string `json:"Line,omitempty"`
//	Locked       bool   `json:"Locked,omitempty"`
//	Priority     int    `json:"Priority,omitempty"`
//	RR           string `json:"RR,omitempty"`
//	RecordID     string `json:"RecordId,omitempty"`
//	Status       bool   `json:"Status,omitempty"`
//	TTL          int    `json:"TTL,omitempty"`
//	Type         string `json:"Type,omitempty"`
//	UserClientIP string `json:"UserClientIp,omitempty"`
//	Value        string `json:"Value,omitempty"`
//	Weight       int    `json:"Weight,omitempty"`
//}
//
//type Resp struct {
//	Code      string `json:"Code,omitempty"`
//	HostID    string `json:"HostId,omitempty"`
//	Message   string `json:"Message,omitempty"`
//	RecordID  string `json:"RecordId,omitempty"`
//	RequestID string `json:"RequestId,omitempty"`
//}
//
//func (cli *Client) AddDomainRecord(domain string, record Record) (int, error) {
//
//	param := url.Values{
//		"DomainName": {domain},
//		"RR":         {record.RR},
//		"Type":       {strings.ToUpper(record.Type)},
//		"Value":      {record.Value},
//	}
//
//	var respInfo struct {
//		Resp
//	}
//
//	_, err := cli.requestGET("AddDomainRecord", param, &respInfo)
//	if err != nil {
//		return 0, err
//	}
//
//	if respInfo.RecordID == "" {
//		return 0, errors.New(respInfo.Message)
//	}
//
//	rrID, _ := strconv.Atoi(respInfo.Message)
//
//	return rrID, nil
//}
//
//func (cli *Client) DescribeDomainRecords(domain string) error {
//	param := url.Values{
//		"DomainName": {domain},
//		"Line":       {"default"},
//		"PageNumber": {"1"},
//		"PageSize":   {"20"},
//	}
//
//	var respInfo struct {
//		DomainRecords struct {
//			Record []Record `json:"Record"`
//		} `json:"DomainRecords"`
//		PageNumber int    `json:"PageNumber"`
//		PageSize   int    `json:"PageSize"`
//		RequestID  string `json:"RequestId"`
//		TotalCount int    `json:"TotalCount"`
//	}
//
//	_, err := cli.requestGET("DescribeDomainRecords", param, &respInfo)
//	if err != nil {
//		return err
//	}
//
//	fmt.Println(respInfo.PageSize)
//	return nil
//
//}

package alidns

import (
	"fmt"
)

type Record struct{}

func (cli *Client) AddDomainRecord(domain string, data map[string]string) (int, error) {

	action := "AddDomainRecord"

	data["DomainName"] = domain

	body, err := cli.Do(action, data)
	if err != nil {

		return 0, err
	}
	fmt.Println(string(body))

	return 0, nil
}
