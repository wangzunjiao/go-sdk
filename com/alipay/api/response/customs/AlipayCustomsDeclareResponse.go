package responseCustoms

import (
	"github.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"github.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response"
)

type AlipayCustomsDeclareResponse struct {
	response.AlipayResponse
	CustomsPaymentId              string                    `json:"customsPaymentId"`
	CustomsOrderId                string                    `json:"customsOrderId"`
	IdentityCheckResult           model.IdentityCheckResult `json:"identityCheckResult"`
	ClearingChannel               model.ClearingChannel     `json:"clearingChannel"`
	ClearingTransactionId         string                    `json:"clearingTransactionId"`
	CustomsProviderRegistrationId string                    `json:"customsProviderRegistrationId"`
}
