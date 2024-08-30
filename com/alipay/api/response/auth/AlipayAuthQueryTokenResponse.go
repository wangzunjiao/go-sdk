package responseAuth

import "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response"

type AlipayAuthQueryTokenResponse struct {
	response.AlipayResponse
	AccessToken            string `json:"accessToken"`
	AccessTokenExpiryTime  string `json:"accessTokenExpiryTime"`
	RefreshToken           string `json:"refreshToken"`
	RefreshTokenExpiryTime string `json:"refreshTokenExpiryTime"`
	TokenStatusType        string `json:"tokenStatusType"`
}
