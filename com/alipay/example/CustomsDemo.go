package main

import (
	defaultAlipayClient "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/request/customs"
	responseCustoms "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response/customs"
	"fmt"
)

func main() {

	client := defaultAlipayClient.NewDefaultAlipayClient(
		alipay_getwayurl,
		alipay_clientid,
		alipay_merchantPrivateKey,
		alipay_alipayPublicKey)
	//declares("202408221940108001001887E0207467163", client)
	inquiryDeclaration([]string{""}, client)

}

func declares(paymentId string, client *defaultAlipayClient.DefaultAlipayClient) {
	request, customsDeclareRequest := customs.NewAlipayCustomsDeclareRequest()
	customsDeclareRequest.PaymentId = paymentId
	customsDeclareRequest.DeclarationRequestId = "12124324324235235"
	customsDeclareRequest.DeclarationAmount = &model.Amount{"CNY", "1000"}
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
