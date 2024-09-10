package responseSubscription

import "github.com/wangzunjiao/go-sdk/com/alipay/api/response"

type AlipaySubscriptionCreateResponse struct {
	response.AlipayResponse
	SchemeUrl     string `json:"schemeUrl"`
	ApplinkUrl    string `json:"applinkUrl"`
	NormalUrl     string `json:"normalUrl"`
	AppIdentifier string `json:"appIdentifier"`
}
