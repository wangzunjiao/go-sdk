package responseCustoms

import (
	"github.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"github.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response"
)

type AlipayCustomsQueryResponse struct {
	response.AlipayResponse
	DeclarationRequestsNotFound []string                  `json:"declarationRequestsNotFound"`
	DeclarationRecords          []model.DeclarationRecord `json:"declarationRecords"`
}
