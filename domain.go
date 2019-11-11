package alidns

// DescribeRecordResponse 返回结构体
type DescribeRecordResponse struct {
	DomainRecords struct {
		Record []RecordInfo `json:"Record"`
	} `json:"DomainRecords"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	RequestID  string `json:"RequestId"`
	TotalCount int    `json:"TotalCount"`
}

// DescribeDomainRecords 调用DescribeDomainRecords根据传入参数获取指定主域名的所有解析记录列表。
// https://help.aliyun.com/document_detail/29776.html?spm=a2c4g.11186623.6.638.f7553b59curplN
func (cli *Client) DescribeDomainRecords(domain string, optional map[string]string) (respInfo DescribeRecordResponse, errResp ErrorResponse, err error) {
	body := map[string]string{
		"DomainName": domain,
	}

	errResp, err = cli.Do("DescribeDomainRecords", body, optional, &respInfo)
	return respInfo, errResp, err
}
