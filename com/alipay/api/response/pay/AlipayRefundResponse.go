package responsePay

import (
	"github.com/wangzunjiao/go-sdk/com/alipay/api/model"
	"github.com/wangzunjiao/go-sdk/com/alipay/api/response"
)

type AlipayRefundResponse struct {
	response.AlipayResponse
	RefundRequestId                string             `json:"refundRequestId"`
	RefundId                       string             `json:"refundId"`
	PaymentId                      string             `json:"paymentId"`
	RefundAmount                   model.Amount       `json:"refundAmount"`
	RefundTime                     string             `json:"refundTime"`
	RefundNonGuaranteeCouponAmount model.Amount       `json:"refundNonGuaranteeCouponAmount"`
	GrossSettlementAmount          model.Amount       `json:"grossSettlementAmount"`
	SettlementQuote                model.Quote        `json:"settlementQuote"`
	AcquirerInfo                   model.AcquirerInfo `json:"acquirerInfo"`
	AcquirerReferenceNo            string             `json:"acquirerReferenceNo"`
}
