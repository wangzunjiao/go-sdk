package responseAuth

import (
	"github.com/wangzunjiao/go-sdk/com/alipay/api/model"
	"github.com/wangzunjiao/go-sdk/com/alipay/api/response"
)

type AlipayAuthConsultResponse struct {
	response.AlipayResponse
	AuthUrl       string             `json:"authUrl"`
	ExtendInfo    string             `json:"extendInfo"`
	NormalUrl     string             `json:"normalUrl"`
	SchemeUrl     string             `json:"schemeUrl"`
	ApplinkUrl    string             `json:"applinkUrl"`
	AppIdentifier string             `json:"appIdentifier"`
	AuthCodeForm  model.AuthCodeForm `json:"authCodeForm"`
	GrantType     string             `json:"grant_type"`
}
