package main

import (
	"fmt"
	"github.com/google/uuid"
	defaultAlipayClient "github.com/wangzunjiao/go-sdk/com/alipay/api"
	"github.com/wangzunjiao/go-sdk/com/alipay/api/model"
	"github.com/wangzunjiao/go-sdk/com/alipay/api/request/customs"
	responseCustoms "github.com/wangzunjiao/go-sdk/com/alipay/api/response/customs"
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
	Declare("202408221940108001001887E0207467163", client)
	//InquiryDeclaration([]string{""}, client)

}

func Declare(paymentId string, client *defaultAlipayClient.DefaultAlipayClient) {
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

func InquiryDeclaration(declareRequestId []string, client *defaultAlipayClient.DefaultAlipayClient) {
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
