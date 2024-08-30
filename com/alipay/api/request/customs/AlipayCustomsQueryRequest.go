package customs

import (
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/request"
	responseCustoms "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response/customs"
)

type AlipayCustomsQueryRequest struct {
	DeclarationRequestIds []string `json:"declarationRequestIds,omitempty"`
}

func (alipayCustomsQueryRequest *AlipayCustomsQueryRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCustomsQueryRequest, model.INQUIRY_DECLARE_PATH, &responseCustoms.AlipayCustomsQueryResponse{})
}

func NewAlipayCustomsQueryRequest() (*request.AlipayRequest, *AlipayCustomsQueryRequest) {
	alipayCustomsQueryRequest := &AlipayCustomsQueryRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCustomsQueryRequest, model.INQUIRY_DECLARE_PATH, &responseCustoms.AlipayCustomsQueryResponse{})
	return alipayRequest, alipayCustomsQueryRequest
}
