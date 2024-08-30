package responsePay

import (
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response"
)

type AlipayPayConsultResponse struct {
	response.AlipayResponse
	PaymentOptions     []model.PaymentOption     `json:"paymentOptions"`
	PaymentMethodInfos []model.PaymentMethodInfo `json:"paymentMethodInfos"`
	ExtendInfo         string                    `json:"extendInfo"`
}
