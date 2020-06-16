package main

import (
	"fmt"

	dynu "github.com/danielerez/dynu-client/client"
)

const (
	DynuApiKey   = "api-secret-key"
	DynuDomainId = "dns-domain-id"
)

func main() {
	d := dynu.NewDynuClient(DynuApiKey)
	record, err := d.CreateRecordA(DynuDomainId, "api", "test-cluster", "192.168.126.100")
	if err != nil {
		fmt.Printf("Couldn't create 'api' record:\n%v", err)
		return
	}
	fmt.Printf("%+v\n", *record)

	record, err = d.CreateRecordA(DynuDomainId, "apps", "test-cluster", "192.168.126.101")
	if err != nil {
		fmt.Printf("Couldn't create 'apps' record:\n%v", err)
		return
	}
	fmt.Printf("%+v\n", *record)
}
