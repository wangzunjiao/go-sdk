package main

import (
	defaultAlipayClient "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/request/pay"
	responsePay "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response/pay"
	"fmt"
)

const alipay_merchantPrivateKey = ""
const alipay_alipayPublicKey = ""
const alipay_clientid = ""
const alipay_getwayurl = ""

func main() {

	client := defaultAlipayClient.NewDefaultAlipayClient(
		alipay_getwayurl,
		alipay_clientid,
		alipay_merchantPrivateKey,
		alipay_alipayPublicKey)

	AliPay(client)

	//refund("202408151940108001001886C0209996792", client)

	//queryRefund("ad53716a-81-4c4c-b151-216c5225e", client)
	//consult(client)
	//cancel("ad53716a-81-4c4c-b51-20916c5225e", client) // ad53716a-81-4c4c-b51-20916c5225e 202408221940108001001888E0210714455

	//createSession(client)
}

func AliPay(body *defaultAlipayClient.DefaultAlipayClient) {

	payRequest, request := pay.NewAlipayPayRequest()

	request.PaymentRequestId = "ad716a-81-4c4c-b51-20916c5225e"
	order := &model.Order{}
	order.OrderDescription = "example order"
	order.ReferenceOrderId = "28947397358748"
	order.OrderAmount = model.NewAmount("100", "HKD")
	merchant := &model.Merchant{}
	merchant.ReferenceMerchantId = "1238rye8yr8erwer"
	merchant.MerchantMCC = "7011"
	merchant.MerchantName = "example merchant"
	merchant.Store = &model.Store{StoreMCC: "7011", ReferenceStoreId: "289285674", StoreName: "store 1111"}
	order.Merchant = merchant
	order.Env = &model.Env{OsType: model.ANDROID, TerminalType: model.WEB}
	request.Order = order

	request.PaymentAmount = model.NewAmount("100", "HKD")

	request.PaymentNotifyUrl = "https://www.alipay.com"
	request.PaymentRedirectUrl = "https://www.alipay.com"

	request.PaymentMethod = &model.PaymentMethod{PaymentMethodType: model.ALIPAY_HK, PaymentMethodId: "1234567890"}

	request.ProductCode = model.CASHIER_PAYMENT

	execute, err := body.Execute(payRequest)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPayResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)
}

func queryPay(paymentRequestId string, body *defaultAlipayClient.DefaultAlipayClient) {
	queryRequest := pay.AlipayPayQueryRequest{}
	queryRequest.PaymentRequestId = paymentRequestId
	request := queryRequest.NewRequest()
	//1. 这里接收结果
	execute, err := body.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPayQueryResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)
}

func refund(paymentId string, client *defaultAlipayClient.DefaultAlipayClient) {
	refundRequest := pay.AlipayRefundRequest{}
	refundRequest.RefundRequestId = "ad53716a-81-4c4c-b151-216c5225e"
	refundRequest.PaymentId = paymentId
	refundRequest.RefundAmount = model.NewAmount("100", "HKD")
	refundRequest.RefundReason = "example refund"
	request := refundRequest.NewRequest()
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayRefundResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)

}

func queryRefund(refundRequestId string, client *defaultAlipayClient.DefaultAlipayClient) {
	queryRefundRequest := pay.AlipayInquiryRefundRequest{}
	queryRefundRequest.RefundRequestId = refundRequestId
	request := queryRefundRequest.NewRequest()
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayInquiryRefundResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)
}

func consult(client *defaultAlipayClient.DefaultAlipayClient) {
	consultRequest := &pay.AlipayPayConsultRequest{}
	request := consultRequest.NewRequest()
	consultRequest.SettlementStrategy = &model.SettlementStrategy{
		SettlementCurrency: "USD",
	}
	consultRequest.ProductCode = model.CASHIER_PAYMENT
	consultRequest.UserRegion = "SG"
	consultRequest.AllowedPaymentMethodRegions = []string{"SG", "HK", "CN"}
	consultRequest.Env = &model.Env{
		OsType:       model.IOS,
		TerminalType: model.APP,
	}
	consultRequest.PaymentAmount = model.NewAmount("1000", "USD")
	consultRequest.PaymentFactor = &model.PaymentFactor{
		PresentmentMode: model.BUNDLE,
	}

	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPayConsultResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)

}

func cancel(paymentRequestId string, client *defaultAlipayClient.DefaultAlipayClient) {
	request, cancelRequest := pay.NewAlipayPayCancelRequest()
	cancelRequest.PaymentRequestId = paymentRequestId
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPayCancelResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)

}

func createSession(client *defaultAlipayClient.DefaultAlipayClient) {
	request, createSessionRequest := pay.NewAlipayPaymentSessionRequest()
	createSessionRequest.PaymentRequestId = "ad53716a-81-4c4c-b151-216c5225e"
	createSessionRequest.Order = &model.Order{
		OrderDescription: "example order",
		ReferenceOrderId: "289473927358748",
		OrderAmount:      model.NewAmount("100", "HKD"),
		Buyer: &model.Buyer{
			ReferenceBuyerId: "111112132143434",
		},
	}
	createSessionRequest.PaymentAmount = model.NewAmount("100", "HKD")
	createSessionRequest.ProductCode = model.CASHIER_PAYMENT
	createSessionRequest.PaymentMethod = &model.PaymentMethod{
		PaymentMethodType: model.SHOPEEPAY_SG,
	}
	createSessionRequest.PaymentNotifyUrl = "https://www.alipay.com"
	createSessionRequest.PaymentRedirectUrl = "https://www.alipay.com"
	createSessionRequest.Env = &model.Env{
		OsType:       model.IOS,
		TerminalType: model.APP,
	}

	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responsePay.AlipayPaymentSessionResponse)
	fmt.Println(response.AlipayResponse.Result.ResultCode)
	fmt.Println(response)
}
