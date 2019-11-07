package alidns

import (
	"fmt"
	"net/url"
	"testing"
	"time"
)

func Test_Sig(t *testing.T) {

	var p = url.Values{}

	p.Add("akid", "test")

	fmt.Println(p.Encode())

	//uri:=`Format=XML&AccessKeyId=testid&Action=DescribeDomainRecords&SignatureMethod=HMAC-SHA1&DomainName=example.com&SignatureNonce=f59ed6a9-83fc-473b-9cc6-99c95df3856e&SignatureVersion=1.0&Version=2015-01-09&Timestamp=2016-03-24T16:41:54Z`
	uri := "/"
	fmt.Println(url.QueryEscape(uri))

	//time2, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	//
	//fmt.Println(time2)

	timestamp := time.Now().Format("2006-01-02T15:04:05Z")

	fmt.Println(timestamp)
}

func Test_url(t *testing.T) {
	uri := url.Values{}

	uri.Add("AccessKeyId", "testid")
	uri.Add("Action", "DescribeDomainRecords")
	uri.Add("SignatureMethod", "HMAC-SHA1")
	uri.Add("DomainName", "example.com")
	uri.Add("SignatureVersion", "1.0")
	uri.Add("Version", "2015-01-09")
	uri.Add("SignatureNonce", "f59ed6a9-83fc-473b-9cc6-99c95df3856e")
	uri.Add("Format", "JSON")

	fmt.Println(uri.Encode())

	fmt.Println(url.QueryEscape(uri.Encode()))
}

// Signature 签名机制
// https://help.aliyun.com/document_detail/29747.html?spm=a2c4g.11186623.6.619.648d5279NR2Fyp
func Test_Signature(t *testing.T) {

	uri := url.Values{}

	uri.Add("Format", "XML")
	uri.Add("AccessKeyId", "testid")
	uri.Add("Action", "DescribeDomainRecords")
	uri.Add("SignatureMethod", "HMAC-SHA1")
	uri.Add("DomainName", "example.com")
	uri.Add("SignatureNonce", "f59ed6a9-83fc-473b-9cc6-99c95df3856e")
	uri.Add("SignatureVersion", "1.0")
	uri.Add("Version", "2015-01-09")
	uri.Add("Timestamp", "2016-03-24T16:41:54Z")

	cannoURI := "GET" + "&" + url.QueryEscape("/") + "&" + url.QueryEscape(uri.Encode())
	code := ShaHmac1(cannoURI, "testsecret&")

	fmt.Println(code)
	if code == "uRpHwaSEt3J+6KQD//svCh/x+pI=" {
		fmt.Println("success")
	}

	code2 := Signature("GET", uri, "testsecret&")
	fmt.Println(code2)
	if code2 == "uRpHwaSEt3J+6KQD//svCh/x+pI=" {
		fmt.Println("success")
	}

}

func TestNew(t *testing.T) {
	//nowtime := time.Now()
	//
	//fmt.Println(nowtime.Format("2006-01-02T15:04:05Z"))
	//fmt.Println(nowtime.Format("2006-01-02T15:04:05Z-0700"))
	//fmt.Println(nowtime.Format("2006-01-02T15:04:05Z+0800"))

	loc, _ := time.LoadLocation("") //参数就是解压文件的“目录”+“/”+“文件名”。
	//fmt.Println(time.Now().In(loc))

	timeNow := time.Now().In(loc).Format("2006-01-02T15:04:05Z")
	fmt.Println(timeNow)
	fmt.Println(url.QueryEscape(timeNow))

	fmt.Println(url.QueryEscape("2016-03-24T16:41:54Z"))

}
