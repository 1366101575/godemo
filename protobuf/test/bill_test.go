package protobuf

import (
	"fmt"
	"godemo/protobuf/proto_src/bill"
	"testing"
)

func TestBill(t *testing.T) {
	billRequest := bill.ApiRequestEvent{
		AppId: 11,
	}

	billRequest.TestOneof = &bill.ApiRequestEvent_Name{Name: "test"}

	fmt.Printf("\n\n%+v\n\n", billRequest.GetTestOneof())

	billRequest.TestOneof = &bill.ApiRequestEvent_Id{Id: 999}

	fmt.Printf("\n\n%+v\n\n", billRequest.GetTestOneof())
}
