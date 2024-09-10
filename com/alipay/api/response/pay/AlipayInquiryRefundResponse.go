package responsePay

import (
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response"
)

type AlipayInquiryRefundResponse struct {
	response.AlipayResponse
	RefundId              string                      `json:"refundId"`
	RefundRequestId       string                      `json:"refundRequestId"`
	RefundAmount          model.Amount                `json:"refundAmount"`
	RefundStatus          model.TransactionStatusType `json:"refundStatus"`
	RefundTime            string                      `json:"refundTime"`
	GrossSettlementAmount model.Amount                `json:"grossSettlementAmount"`
	SettlementQuote       model.Quote                 `json:"settlementQuote"`
	AcquirerInfo          model.AcquirerInfo          `json:"acquirerInfo"`
}
