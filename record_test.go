package alidns

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestClient_AddDomainRecord2(t *testing.T) {

	cli := New(akid, akey)

	RR := "tangx12223"
	Type := "A"
	Value := "123.231.12.11"
	ValueUpdate := "111.222.33.44"

	// add
	resp, errResp, err := cli.AddDomainRecord("rockontrol.com", RR, Type, Value, nil)

	if err != nil {
		log.Fatal(errResp.Code, errResp.Message)
	}
	fmt.Println(resp.RecordId)
	rrID := resp.RecordId

	// update
	time.Sleep(1 * time.Second)
	resp, errResp, err = cli.UpdateDomainRecord(RR, rrID, Type, ValueUpdate, nil)
	if err != nil {
		log.Fatal(errResp.Message)
	}
	fmt.Println(resp.RecordId)

	// describe
	time.Sleep(1 * time.Second)
	descResp, errResp, err := cli.DescribeDomainRecordInfo(rrID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(descResp)

	// delete
	time.Sleep(1 * time.Second)
	resp, errResp, err = cli.DeleteDomainRecord(rrID)
	if err != nil {
		log.Fatal(errResp.Message)
	}
	fmt.Println(resp.RecordId)
}
