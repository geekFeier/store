package serve

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

var privateStr = string(``)

var privateText = []byte(privateStr)

//const is
const (
	//ContentTypeForm = "application/x-www-form-urlencoded;charset=utf-8"
	ContentTypeForm = "text/html;charset=utf-8"
)

//Alipay is
type Alipay struct {
	AppID      string `json:"app_id"`
	Method     string `json:"method"`
	Charset    string `json:"charset"`
	ReturnURL  string `json:"return_url"`
	SignType   string `json:"sign_type"`
	Sign       string `json:"sign"`
	Timestamp  string `json:"timestamp"`
	Version    string `json:"version"`
	NotifyURL  string `json:"notify_url"`
	BizContent string `json:"biz_content"`
}

//AlipayReq is
type AlipayReq struct {
	OutTradeNo  string  `json:"out_trade_no"`
	ProductCode string  `json:"product_code"`
	TotalAmount float64 `json:"total_amount"`
	Subject     string  `json:"subject"`
}

//GetBizContent is
func GetBizContent(req *AlipayReq) string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("error pay req : %v", req)
		return ""
	}

	return fmt.Sprintf("%s", b)
}

//SortPay is
func SortPay(pay *Alipay, req *AlipayReq) string {
	pay.BizContent = GetBizContent(req)

	s := fmt.Sprintf("app_id=%s&biz_content=%s&charset=utf-8&method=%s&notity_url=%s&return_url=%s&sign_type=%s&timestamp=%s&version=1.0",
		strings.TrimSpace(pay.AppID),
		strings.TrimSpace(pay.BizContent),
		strings.TrimSpace(pay.Method),
		strings.TrimSpace(pay.NotifyURL),
		strings.TrimSpace(pay.ReturnURL),
		strings.TrimSpace(pay.SignType),
		strings.TrimSpace(pay.Timestamp),
	)
	fmt.Println("req: ", s)
	return s
}

//URLEscape is
func URLEscape(pay *Alipay, req *AlipayReq) string {
	pay.BizContent = GetBizContent(req)

	s := fmt.Sprintf("app_id=%s&biz_content=%s&charset=utf-8&method=%s&notity_url=%s&return_url=%s&sign_type=%s&timestamp=%s&version=1.0",
		url.QueryEscape(strings.TrimSpace(pay.AppID)),
		url.QueryEscape(strings.TrimSpace(pay.BizContent)),
		url.QueryEscape(strings.TrimSpace(pay.Method)),
		url.QueryEscape(strings.TrimSpace(pay.NotifyURL)),
		url.QueryEscape(strings.TrimSpace(pay.ReturnURL)),
		url.QueryEscape(strings.TrimSpace(pay.SignType)),
		url.QueryEscape(strings.TrimSpace(pay.Timestamp)),
	)
	fmt.Println("req: ", s)
	return s

}

//Sign is
func Sign(body string) string {
	/*
		rsaPrivateKey, err := ioutil.ReadFile("rsa_private_key.pem")
		if err != nil {
			fmt.Println("read private key failed : ", err)
			return ""
		}
	*/
	rsaStr := os.Getenv("RSA_PRIVATE_KEY")
	rsaPrivateKey, err := base64.StdEncoding.DecodeString(rsaStr)
	if err != nil {
		fmt.Println("read private key failed : ", err)
		return ""
	}

	block, _ := pem.Decode(rsaPrivateKey)

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("load private key failed", err)
		return ""
	}

	rng := rand.Reader
	para := []byte(body)
	hashed := sha256.Sum256(para)

	signature, err := rsa.SignPKCS1v15(rng, key, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return ""
	}
	encoded := base64.StdEncoding.EncodeToString(signature)
	fmt.Printf("signature %s", encoded)
	return encoded
}

// PayURL 生成支付宝即时到帐提交表单html代码
func PayURL(totalAmount float64, outTradeNo, subject, returnURL, notifyURL string) string {
	//实例化参数
	pay := &Alipay{
		AppID:     "2018121662557851",
		Method:    "alipay.trade.page.pay",
		ReturnURL: returnURL,
		NotifyURL: notifyURL,
	}
	req := &AlipayReq{
		OutTradeNo:  outTradeNo,
		ProductCode: "FAST_INSTANT_TRADE_PAY",
		TotalAmount: totalAmount,
		Subject:     subject,
	}
	pay.SignType = "RSA2"
	pay.Timestamp = time.Now().Format("2006-01-02 15:04:05")

	st := SortPay(pay, req)
	//生成签名
	sign := Sign(st)

	//追加参数
	pay.Sign = sign

	url := fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s&sign=%s", URLEscape(pay, req), url.QueryEscape(sign))
	fmt.Println("encode: ", url)

	return url

	//生成自动提交form
	/*
		return `
			<form id="alipaysubmit" name="alipaysubmit" action="https://openapi.alipay.com/gateway.do" method="post" style='display:none;'>
				<input type="hidden" name="app_id" value="` + pay.AppID + `">
				<input type="hidden" name="biz_content" value="` + strings.TrimSpace(pay.BizContent) + `">
				<input type="hidden" name="charset" value="utf-8">
				<input type="hidden" name="method" value="` + pay.Method + `">
				<input type="hidden" name="sign_type" value="` + pay.SignType + `">
				<input type="hidden" name="timestamp" value="` + strings.TrimSpace(pay.Timestamp) + `">
				<input type="hidden" name="version" value="1.0">
				<input type="hidden" name="sign" value="` + pay.Sign + `">
			</form>
			<script>
				document.forms['alipaysubmit'].submit();
			</script>
		`
	*/
}
