package responsePay

import (
	"github.com/wangzunjiao/go-sdk/com/alipay/api/model"
	"github.com/wangzunjiao/go-sdk/com/alipay/api/response"
)

type AlipayPayConsultResponse struct {
	response.AlipayResponse
	PaymentOptions     []model.PaymentOption     `json:"paymentOptions"`
	PaymentMethodInfos []model.PaymentMethodInfo `json:"paymentMethodInfos"`
	ExtendInfo         string                    `json:"extendInfo"`
}
