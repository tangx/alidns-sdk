# alidns-sdk

## demo

```go

	cli := New(akid, akey)

	RR := "tangx12223"
	Type := "A"
	Value := "123.231.12.11"

	// add domain record
	resp, errResp, err := cli.AddDomainRecord("example.com", RR, Type, Value, nil)

    // 如果 err != nil , 则 errResp 值为 aliyun 错误返回值， 否则为 0值。
	if err != nil {
		log.Fatal(errResp.Code, errResp.Message)
	}
	fmt.Println(resp.RecordId)

```

## SKD 验签
+ [x] signature

## Record

+ [x] 增加域名记录
+ [x] 根据ID删除域名记录
+ [x] 更新域名记录信息
+ [x] 显示域名所有记录

## Domain
+ + [x] 显示域名的所有解析记录
