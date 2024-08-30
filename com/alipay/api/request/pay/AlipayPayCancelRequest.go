package pay

import (
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/request"
	responsePay "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response/pay"
)

type AlipayPayCancelRequest struct {
	PaymentId         string `json:"paymentId,omitempty"`
	PaymentRequestId  string `json:"paymentRequestId,omitempty"`
	MerchantAccountId string `json:"merchantAccountId,omitempty"`
}

func NewAlipayPayCancelRequest() (*request.AlipayRequest, *AlipayPayCancelRequest) {
	alipayPayCancelRequest := &AlipayPayCancelRequest{}
	alipayRequest := request.NewAlipayRequest(alipayPayCancelRequest, model.CANCEL_PATH, &responsePay.AlipayPayCancelResponse{})
	return alipayRequest, alipayPayCancelRequest
}
