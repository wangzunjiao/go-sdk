package responseAuth

import (
	"github.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"github.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response"
)

type AlipayAuthApplyTokenResponse struct {
	response.AlipayResponse
	AccessToken            string                `json:"accessToken"`
	AccessTokenExpiryTime  string                `json:"accessTokenExpiryTime"`
	RefreshToken           string                `json:"refreshToken"`
	RefreshTokenExpiryTime string                `json:"refreshTokenExpiryTime"`
	ExtendInfo             string                `json:"extendInfo"`
	UserLoginId            string                `json:"userLoginId"`
	PspCustomerInfo        model.PspCustomerInfo `json:"pspCustomerInfo"`
	GrantType              string                `json:"grant_type"`
}
