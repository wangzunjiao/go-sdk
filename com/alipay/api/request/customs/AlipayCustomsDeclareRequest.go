package customs

import (
	"github.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"github.com/wangzunjiao.wzj/go-sdk/com/alipay/api/request"
	responseCustoms "github.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response/customs"
)

type AlipayCustomsDeclareRequest struct {
	DeclarationRequestId string                     `json:"declarationRequestId,omitempty"`
	PaymentId            string                     `json:"paymentId,omitempty"`
	DeclarationAmount    *model.Amount              `json:"declarationAmount,omitempty"`
	Customs              *model.CustomsInfo         `json:"customs,omitempty"`
	MerchantCustomsInfo  *model.MerchantCustomsInfo `json:"merchantCustomsInfo,omitempty"`
	SplitOrder           bool                       `json:"splitOrder,omitempty"`
	SubOrderId           string                     `json:"subOrderId,omitempty"`
	BuyerCertificate     *model.Certificate         `json:"buyerCertificate,omitempty"`
}

func (alipayCustomsDeclareRequest *AlipayCustomsDeclareRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipayCustomsDeclareRequest, model.DECLARE_PATH, &responseCustoms.AlipayCustomsDeclareResponse{})
}

func NewAlipayCustomsDeclareRequest() (*request.AlipayRequest, *AlipayCustomsDeclareRequest) {
	alipayCustomsDeclareRequest := &AlipayCustomsDeclareRequest{}
	alipayRequest := request.NewAlipayRequest(alipayCustomsDeclareRequest, model.DECLARE_PATH, &responseCustoms.AlipayCustomsDeclareResponse{})
	return alipayRequest, alipayCustomsDeclareRequest
}
