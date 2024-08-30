package responseSubscription

import "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response"

type AlipaySubscriptionCreateResponse struct {
	response.AlipayResponse
	SchemeUrl     string `json:"schemeUrl"`
	ApplinkUrl    string `json:"applinkUrl"`
	NormalUrl     string `json:"normalUrl"`
	AppIdentifier string `json:"appIdentifier"`
}
