package alidns

import (
	"strconv"
)

// AddDomainRecord Add Domain Record
func (cli *Client) AddDomainRecord(domain string, data map[string]string) (int, error) {

	data["DomainName"] = domain

	var respInfo struct {
		RecordID  string
		RequestID string
	}
	err := cli.Do("AddDomainRecord", data, &respInfo)
	if err != nil {
		return 0, err
	}

	//fmt.Println(respInfo)
	rrID, err := strconv.Atoi(respInfo.RecordID)
	return rrID, nil
}
