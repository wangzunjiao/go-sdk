package auth

import (
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/request"
	responseAuth "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response/auth"
)

type AlipayAuthQueryTokenRequest struct {
	AccessToken string `json:"accessToken"`
}

func NewAlipayAuthQueryTokenRequest() (*request.AlipayRequest, *AlipayAuthQueryTokenRequest) {
	alipayAuthQueryTokenRequest := &AlipayAuthQueryTokenRequest{}
	alipayRequest := request.NewAlipayRequest(alipayAuthQueryTokenRequest, model.AUTH_QUERY_PATH, &responseAuth.AlipayAuthQueryTokenResponse{})
	return alipayRequest, alipayAuthQueryTokenRequest
}
