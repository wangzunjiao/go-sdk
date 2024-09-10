package main

import (
	defaultAlipayClient "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/request/customs"
	responseCustoms "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response/customs"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	const alipayGatewayUrl = ""
	const alipayClientId = ""
	const alipayMerchantPrivateKey = ""
	const alipayAlipayPublicKey = ""

	client := defaultAlipayClient.NewDefaultAlipayClient(
		alipayGatewayUrl,
		alipayClientId,
		alipayMerchantPrivateKey,
		alipayAlipayPublicKey)
	declare("202408221940108001001887E0207467163", client)
	//inquiryDeclaration([]string{""}, client)

}

func declare(paymentId string, client *defaultAlipayClient.DefaultAlipayClient) {
	request, customsDeclareRequest := customs.NewAlipayCustomsDeclareRequest()
	customsDeclareRequest.PaymentId = paymentId
	customsDeclareRequest.DeclarationRequestId = uuid.NewString()
	customsDeclareRequest.DeclarationAmount = &model.Amount{Value: "CNY", Currency: "1000"}
	customsDeclareRequest.MerchantCustomsInfo = &model.MerchantCustomsInfo{
		MerchantCustomsName: "merchantCustomsName",
		MerchantCustomsCode: "merchantCustomsCode",
	}
	customsDeclareRequest.SplitOrder = false
	customsDeclareRequest.Customs = &model.CustomsInfo{
		Region:      "CN",
		CustomsCode: "ZHENGZHOU",
	}
	customsDeclareRequest.BuyerCertificate = &model.Certificate{
		CertificateNo:   "certificateNo",
		CertificateType: model.CertificateType_ID_CARD,
		HolderName: &model.UserName{
			FirstName: "firstName",
			LastName:  "lastName",
			FullName:  "fullName",
		},
	}
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responseCustoms.AlipayCustomsDeclareResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)

}

func inquiryDeclaration(declareRequestId []string, client *defaultAlipayClient.DefaultAlipayClient) {
	request, customsQueryRequest := customs.NewAlipayCustomsQueryRequest()
	customsQueryRequest.DeclarationRequestIds = declareRequestId
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responseCustoms.AlipayCustomsQueryResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)
}
