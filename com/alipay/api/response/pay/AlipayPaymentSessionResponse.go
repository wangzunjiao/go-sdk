package responsePay

import "github.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response"

type AlipayPaymentSessionResponse struct {
	response.AlipayResponse
	PaymentSessionData       string `json:"paymentSessionData"`
	PaymentSessionExpiryTime string `json:"paymentSessionExpiryTime"`
	PaymentSessionId         string `json:"paymentSessionId"`
	NormalUrl                string `json:"normalUrl"`
}
