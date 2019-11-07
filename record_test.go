package alidns

import (
	"fmt"
	"testing"
	"time"
)

func TestClient_AddDomainRecord(t *testing.T) {

	cli := New(akid, akey)

	RR := "tangx1222323"
	Type := "A"
	Value := "123.231.12.11"

	// add
	rrId, err := cli.AddDomainRecord("rockontrol.com", RR, Type, Value, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("RequestId: %d \n", rrId)

	time.Sleep(3 * time.Second)

	// update
	ok, err := cli.UpdateDomainRecord(RR, rrId, Type, "13.131.12.11", nil)
	if err != nil {
		fmt.Println(err)
	}
	if ok {
		fmt.Println("update success")
	} else {
		fmt.Println("update faild")
	}

	time.Sleep(3 * time.Second)

	// delete
	ok, err = cli.DeleteDomainRecord(rrId)
	if err != nil {
		fmt.Println(err)
	}
	if ok {
		fmt.Println("delete success")
	} else {
		fmt.Println("delete faild")
	}
}

func TestClient_DescribeDomainRecords(t *testing.T) {
	cli := New(akid, akey)

	data, _ := cli.DescribeDomainRecords("rockontrol.com", nil)
	fmt.Println(data)
}
