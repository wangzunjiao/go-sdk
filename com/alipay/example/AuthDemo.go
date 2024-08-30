package main

import (
	defaultAlipayClient "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/model"
	"code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/request/auth"
	responseAuth "code.alipay.com/wangzunjiao.wzj/go-sdk/com/alipay/api/response/auth"
	"fmt"
)

func main() {

	client := defaultAlipayClient.NewDefaultAlipayClient(
		alipay_getwayurl,
		alipay_clientid,
		alipay_merchantPrivateKey,
		alipay_alipayPublicKey)

	authConsult(client)
	//applyToken("281001139639787089651362", client)
	//queryToken("28288803001291161724296551000BgIrDiWzU0171000529", client)
	//revokeToken("28288803001291161724296551000BgIrDiWzU0171000529", client)
}

func authConsult(client *defaultAlipayClient.DefaultAlipayClient) {
	request, authConsultRequest := auth.NewAlipayAuthConsultRequest()
	authConsultRequest.AuthRedirectUrl = "https://www.alipay.com"
	authConsultRequest.AuthState = "123456"
	authConsultRequest.CustomerBelongsTo = model.ALIPAY_CN
	authConsultRequest.OsType = model.ANDROID
	authConsultRequest.OsVersion = "1.0.0"
	authConsultRequest.Scopes = []model.ScopeType{model.ScopeTypeAgreementPay}
	authConsultRequest.TerminalType = model.APP

	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responseAuth.AlipayAuthConsultResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)

}

func applyToken(authCode string, client *defaultAlipayClient.DefaultAlipayClient) {

	request, authApplyTokenRequest := auth.NewAlipayAuthApplyTokenRequest()
	authApplyTokenRequest.GrantType = model.GrantTypeAUTHORIZATION_CODE
	authApplyTokenRequest.CustomerBelongsTo = model.ALIPAY_CN
	authApplyTokenRequest.AuthCode = authCode

	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responseAuth.AlipayAuthApplyTokenResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)

}

func queryToken(accessToken string, client *defaultAlipayClient.DefaultAlipayClient) {
	request, authQueryTokenRequest := auth.NewAlipayAuthQueryTokenRequest()
	authQueryTokenRequest.AccessToken = accessToken
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responseAuth.AlipayAuthQueryTokenResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)
}
func revokeToken(accessToken string, client *defaultAlipayClient.DefaultAlipayClient) {
	request, authRevokeTokenRequest := auth.NewAlipayAuthRevokeTokenRequest()
	authRevokeTokenRequest.AccessToken = accessToken
	execute, err := client.Execute(request)
	if err != nil {
		print(err.Error())
		return
	}
	response := execute.(*responseAuth.AlipayAuthRevokeTokenResponse)
	fmt.Println("result: ", response.AlipayResponse.Result)
	fmt.Println("response: ", response)
}
