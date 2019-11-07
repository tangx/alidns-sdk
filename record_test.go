package alidns

import (
	"fmt"
	"testing"
)

func TestClient_AddDomainRecord(t *testing.T) {

	cli := New(akid, akey)

	RR := "tangx123"
	Type := "A"
	Value := "123.231.12.11"

	rrId, err := cli.AddDomainRecord("rockontrol.com", RR, Type, Value, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("RequestId: %d \n", rrId)
}
