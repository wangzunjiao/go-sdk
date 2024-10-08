package pay

import (
	"github.com/wangzunjiao/go-sdk/com/alipay/api/model"
	"github.com/wangzunjiao/go-sdk/com/alipay/api/request"
)

type AlipayPayQueryRequest struct {
	PaymentRequestId  string `json:"paymentRequestId,omitempty"`
	PaymentId         string `json:"paymentId,omitempty"`
	MerchantAccountId string `json:"MerchantAccountId,omitempty"`
}

func (alipayPayQueryRequest *AlipayPayQueryRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayPayQueryRequest, model.INQUIRY_PAYMENT_PATH, &responsePay.AlipayPayQueryResponse{})
}

func NewAlipayPayQueryRequest() (*request.AlipayRequest, *AlipayPayQueryRequest) {
	alipayPayQueryRequest := &AlipayPayQueryRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPayQueryRequest, model.INQUIRY_PAYMENT_PATH, &responsePay.AlipayPayQueryResponse{})
	return alipayRequest, alipayPayQueryRequest
}
