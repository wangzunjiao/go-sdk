package subscription

import (
	"github.com/wangzunjiao/go-sdk/com/alipay/api/model"
	"github.com/wangzunjiao/go-sdk/com/alipay/api/request"
	responseSubscription "github.com/wangzunjiao/go-sdk/com/alipay/api/response/subscription"
)

type AlipaySubscriptionCancelRequest struct {
	SubscriptionId        string                 `json:"subscriptionId,omitempty"`
	SubscriptionRequestId string                 `json:"subscriptionRequestId,omitempty"`
	CancellationType      model.CancellationType `json:"cancellationType,omitempty"`
}

func (alipaySubscriptionCancelRequest *AlipaySubscriptionCancelRequest) NewRequest() *request.AlipayRequest {
	return request.NewAlipayRequest(&alipaySubscriptionCancelRequest, model.SUBSCRIPTION_CANCEL_PATH, &responseSubscription.AlipaySubscriptionCancelResponse{})
}

func NewAlipaySubscriptionCancelRequest() (*request.AlipayRequest, *AlipaySubscriptionCancelRequest) {
	alipaySubscriptionCancelRequest := &AlipaySubscriptionCancelRequest{}
	alipayRequest := request.NewAlipayRequest(alipaySubscriptionCancelRequest, model.SUBSCRIPTION_CANCEL_PATH, &responseSubscription.AlipaySubscriptionCancelResponse{})
	return alipayRequest, alipaySubscriptionCancelRequest
}
