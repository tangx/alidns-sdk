package alidns

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
)

const ()

// Signature 签名机制
// https://help.aliyun.com/document_detail/29747.html?spm=a2c4g.11186623.6.619.648d5279NR2Fyp
func Signature(reqMethod string, param url.Values, secret string) string {

	source := reqMethod + "&" + url.QueryEscape("/") + "&" + url.QueryEscape(param.Encode())
	return ShaHmac1(source, secret)

}

// ShaHmac1 计算 HMAC-SHA1 校验码
func ShaHmac1(source, secret string) string {
	key := []byte(secret)
	sha1hmac := hmac.New(sha1.New, key)
	sha1hmac.Write([]byte(source))
	signedBytes := sha1hmac.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}
