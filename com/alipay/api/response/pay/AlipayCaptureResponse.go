package responsePay

import (
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response"
)

type AlipayCaptureResponse struct {
	response.AlipayResponse
	CaptureRequestId    string       `json:"captureRequestId"`
	CaptureId           string       `json:"captureId"`
	PaymentId           string       `json:"paymentId"`
	CaptureAmount       model.Amount `json:"captureAmount"`
	CaptureTime         string       `json:"captureTime"`
	AcquirerReferenceNo string       `json:"acquirerReferenceNo"`
}
