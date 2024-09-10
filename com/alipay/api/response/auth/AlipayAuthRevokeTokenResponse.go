package responseAuth

import "github.com/wangzunjiao/go-sdk/com/alipay/api/response"

type AlipayAuthRevokeTokenResponse struct {
	response.AlipayResponse
	ExtendInfo string `json:"extendInfo"`
}
