package auth

import (
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/request"
	responseAuth "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response/auth"
)

type AlipayAuthRevokeTokenRequest struct {
	AccessToken string `json:"accessToken,omitempty"`
	ExtendInfo  string `json:"extendInfo,omitempty"`
}

func (alipayAuthRevokeTokenRequest *AlipayAuthRevokeTokenRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayAuthRevokeTokenRequest, model.AUTH_REVOKE_PATH, &responseAuth.AlipayAuthRevokeTokenResponse{})
}

func NewAlipayAuthRevokeTokenRequest() (*request.AlipayRequest, *AlipayAuthRevokeTokenRequest) {
	alipayAuthRevokeTokenRequest := &AlipayAuthRevokeTokenRequest{}
	alipayRequest := request.NewAlipayRequest(alipayAuthRevokeTokenRequest, model.AUTH_REVOKE_PATH, &responseAuth.AlipayAuthRevokeTokenResponse{})
	return alipayRequest, alipayAuthRevokeTokenRequest
}
