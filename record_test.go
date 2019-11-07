package alidns

import (
	"fmt"
	"testing"
)

func TestClient_AddDomainRecord(t *testing.T) {

	cli := New(akid, akey)
	data := map[string]string{
		"RR":    "tangx22223",
		"Type":  "A",
		"Value": "123.223.112.121",
	}

	rrId, err := cli.AddDomainRecord("rockontrol.com", data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("RequestId: %d \n", rrId)
}
