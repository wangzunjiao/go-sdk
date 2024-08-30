package responseAuth

import "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response"

type AlipayAuthRevokeTokenResponse struct {
	response.AlipayResponse
	ExtendInfo string `json:"extendInfo"`
}
