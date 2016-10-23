# Gotrans
-----------
##### A Midtrans.com client for golang

Supporting both SNAP and regular API.

### How to use?

```
package main

import (
	"fmt"

	"github.com/JesusIslam/gotrans"
)

const (
	MerchantID   = "your_merchant_id"
	ServerKey    = "your_server_key"
	ClientKey    = "your_client_key"
	IsProduction = false
	IsSanitized  = false
	Is3DSecure   = false
)

func main() {
	client := gotrans.New(ServerKey, ClientKey, IsProduction, IsSanitized, Is3DSecure, 0)

	trans := &gotrans.Transaction{
		TransactionDetails: &gotrans.TransactionDetail{
			OrderId:     "12345678913",
			GrossAmount: 100000,
		},
		CustomerDetail: &gotrans.CustomerDetail{
			Email:     "your@email.com",
			FirstName: "Your",
			LastName:  "Name",
			Phone:     "your_phone_number",
		},
		ItemDetails: []*gotrans.ItemDetail{
			&gotrans.ItemDetail{
				ID:       "999999",
				Price:    100000.0,
				Quantity: 1.0,
				Name:     "Something",
			},
		},
		BankTransfer: &gotrans.BankTransfer{
			Bank:     "bca",
			VANumber: "111111",
			FreeText: &gotrans.BCAFreeText{
				Inquiry: []*gotrans.InquiryPayment{
					&gotrans.InquiryPayment{
						ID: "Some text",
						EN: "Some text",
					},
				},
				Payment: []*gotrans.InquiryPayment{
					&gotrans.InquiryPayment{
						ID: "Some text",
						EN: "Some text",
					},
				},
			},
		},
	}

	res, err := client.ChargeBankTransfer(trans)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

```

### Notes
- Still minimum documentations
- Credit card transaction hasn't been tested yet
- Convenience store transaction hasn't been tested yet
- Ewallet transaction hasn't been tested yet
- Internet banking transaction hasn't been tested yet
- Capture transaction hasn't been tested yet
- Approve transaction hasn't been tested yet
- Cancel transaction hasn't been tested yet
- Expire transaction hasn't been tested yet
- SNAP transaction hasn't been tested yet