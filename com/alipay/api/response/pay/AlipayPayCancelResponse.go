package responsePay

import "github.com/wangzunjiao/go-sdk/com/alipay/api/response"

type AlipayPayCancelResponse struct {
	response.AlipayResponse
	PaymentId        string `json:"paymentId"`
	PaymentRequestId string `json:"paymentRequestId"`
	CancelTime       string `json:"cancelTime"`
}
