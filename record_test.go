package alidns

import (
	"log"
	"net/url"
	"testing"
)

func TestClient_AddDomainRecord(t *testing.T) {

	cli := New(akid, akey)

	//rr := Record{
	//	RR:    "tangxo322232322",
	//	Type:  "a",
	//	Value: "222.213.111.222",
	//}

	data := map[string]string{
		"RR":    "tangx22223",
		"Type":  "A",
		"Value": "123.23.12.111",
	}
	rrid, err := cli.AddDomainRecord("rockontrol.com", data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rrid)
}

func TestClient_DescribeDomainRecords(t *testing.T) {

	cli := New(akid, akey)

	uri := url.Values{
		"DomainName": {"rockontrol.com"},
		"Format":     {"JSON"},
	}
	var respInfo struct {
		DomainRecords struct {
			Record []Record `json:"Record"`
		} `json:"DomainRecords"`
		PageNumber int    `json:"PageNumber"`
		PageSize   int    `json:"PageSize"`
		RequestID  string `json:"RequestId"`
		TotalCount int    `json:"TotalCount"`
	}
	_, err := cli.requestGET("DescribeDomainRecords", uri, respInfo)
	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_Do2(t *testing.T) {

	cli := New(akid, akey)

	body := map[string]string{
		"DomainName": "rockontrol.com",
		"Format":     "JSON",
	}

	_, err := cli.Do("DescribeDomainRecords", body)
	if err != nil {
		log.Fatal(err)
	}

	body = map[string]string{
		"RR":         "tangxo322232322",
		"Type":       "a",
		"Value":      "222.213.111.222",
		"DomainName": "rockontrol.com",
	}

	_, err = cli.Do("AddDomainRecord", body)
	if err != nil {
		log.Fatal(err)
	}

}
