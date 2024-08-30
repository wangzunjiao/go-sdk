package pay

import (
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/request"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response/pay"
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
