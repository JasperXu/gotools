package alipay

func CreateAlipayTrade(PayID, TotalFee, Subject, Body string) string {
	params := make(map[string]string)
	params["_input_charset"] = "utf-8"
	params["body"] = Body
	params["notify_url"] = "WebNotifyUrl"
	params["out_trade_no"] = PayID
	params["partner"] = AlipayPartner
	params["payment_type"] = "1"
	params["return_url"] = WebReturnUrl
	params["seller_email"] = WebSellerEmail
	params["service"] = "create_direct_pay_by_user"
	params["subject"] = Subject
	params["total_fee"] = TotalFee
	params["sign"] = computeSignature(params)
	params["sign_type"] = "MD5"

	RequestHtmlText := "<form id='alipaysubmit' name='alipaysubmit' action='" + GATEWAY + "?_input_charset=utf-8' method='get'>"
	for k, v := range params {
		RequestHtmlText += "<input type='hidden' name='" + k + "' value='" + v + "'/>"
	}
	RequestHtmlText += "<input type='submit' value='Submit' style='display:none;'></form>"
	RequestHtmlText += "<script>document.forms['alipaysubmit'].submit();</script>"
	return RequestHtmlText
}
