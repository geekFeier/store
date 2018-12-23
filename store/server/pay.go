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
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"time"
)

var privateStr = string(`
MIIEowIBAAKCAQEA3wA8ROUMryWVwbicdwaasILBurRx4LEZSpMK/llBTR8Njj/C+7E1Gwx4G9Ovr2U7BRtihO35d8SEd0bevXleYBN0YcOQ23r3DdymSTeL6ZXBZGGygJCyd25+Nf28Mq1geUkhZSgWMH+A+gPyKeA/UH8tQ4veIp+eLY/LCh8AbtbK20j/UA870ueBFT/wSNJKEv2UShVa6bL2OweikK9zlgVWGJtCXZVhTbQcPjDdetuTLJ+/0kdQydEI8yt6spKheetqJ4QM7P9ufxFIeWqKcbeEvdw6KKJTzKtRDyAV2lz2RTCDlJbylJgoPEtn1Pw4zFhRlS6UPvcY70aj2ZZWCwIDAQABAoIBAQDB1U2NwN4+m1fJc/MkjmwFAxLre48EdEt8g8VpgiF9rIE25PtRlR2I7lS0M1MhDMe5T61ZyBQwY3OUzdgsL4O11RMzKy8NZ5u4w9MSDyMhHRdlbnoewCcwIq23tl5QWl1h7wwFBkwLSbjNGIL0nUPLb0/jx36E7+MozTT6DiptK8zg46RkrOS+oJKSE9eIbw7TMAp7pJR0ZuAaFF98LdW4jOqA/xClQMIeHRMR788nPbM1lxv5CUVg91U7ZmfMy5uJt0bRyRb9a0YEh2/HmtG48HF4fD7+B+mzdO5Qjr3alPB0rHlhma7+lDHDZgek0X4ge0CWVfHp/xuJYa5z0G8RAoGBAPP09rYvyW9eOjbxWghjSiDIZNKdUjXX6fd+KUElZL9ZZanbLKd48N++7JBYlHmisATfzU0yNkBJOZFIz3JWHWn4tDYdUu96aG7VW66OZDaOWxwkg+Ib6GLD/7m75qsjJgcZqUnbiAQlCoiAHOezRoNFBxuKkpLZQsFOS3d1XdCpAoGBAOoCcB6127ULU9KJ78tbyduUi75Ru1abVet1cY+C7ukKGE8lvLidbLBoexkyYwBbNgLnfqQvLlvg/3wA3IT0wLbc5x42VQuj4PMkwLIM854ks+sNAO3jEm3qZoK0WRxxRD/xNKAwoSrljQ5iJUGVxVpYf0YvieFNLkC3XoSrTX2TAoGAHbR6RfzTnkfu9rm0qMjOQeekvzCAziWYS5aFF3WiKtqL0n7plQrY6aWp0Hm4uobgv/cwXHH8wR8pb1NaTyXFNx4dc2lmq9pP1Q8NtteHxEzZMzgPBv028q/C9661i6kf/EVXo7KjgT2xZqWS33Oo5lsFXoklB9SEmF0cO1ODNOECgYAOMmeq9U2HBlDnjQbHR6JPeAuiWEMNVg84Yb+p/T+RU9N1ucxhRuu5KB2PwcbP1rjIJFNCkro1SMk3NLYsOs0WSwCajKFMO30CQ1CNfMuq5H8l23wa8pPDp1zgwSBG4XhKiM+wiEK+335XSQ0JlZSLjqLqFyg+SvhrZLSPR/VDkQKBgCMt+g+Hr6HTRifYhJC3DeepVT/8GnP1Zuc6EeguaLijHn8fVwsP0LFy8Zj0iABlxw9mP3MT8h48o5YvX7Vmr7aHOstnBYXJFlhEF9NLSGMlOTeCsRlAXjAhp8NUucMc2tOihG7wqPy1cduRJ/eNiYatMUAeqPAAjC5H4QjWhxVr
	`)

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

	s := fmt.Sprintf("app_id=%s&biz_content=%s&charset=utf-8&method=%s&sign_type=%s&timestamp=%s&version=1.0",
		pay.AppID,
		strings.TrimSpace(pay.BizContent),
		pay.Method,
		pay.SignType,
		strings.TrimSpace(pay.Timestamp),
	)
	fmt.Println("req: ", s)
	return s
}

//URLEscape is
func URLEscape(pay *Alipay, req *AlipayReq) string {
	pay.BizContent = GetBizContent(req)

	s := fmt.Sprintf("app_id=%s&biz_content=%s&charset=utf-8&method=%s&sign_type=%s&timestamp=%s&version=1.0",
		url.QueryEscape(strings.TrimSpace(pay.AppID)),
		url.QueryEscape(strings.TrimSpace(pay.BizContent)),
		url.QueryEscape(strings.TrimSpace(pay.Method)),
		url.QueryEscape(strings.TrimSpace(pay.SignType)),
		url.QueryEscape(strings.TrimSpace(pay.Timestamp)),
	)
	fmt.Println("req: ", s)
	return s
}

//Sign is
func Sign(body string) string {
	rsaPrivateKey, err := ioutil.ReadFile("rsa_private_key.pem")
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

// Form 生成支付宝即时到帐提交表单html代码
func Form() string {
	//实例化参数
	pay := &Alipay{
		AppID:  "2018121662557851",
		Method: "alipay.trade.page.pay",
	}
	req := &AlipayReq{
		OutTradeNo:  "1klaskdjfaa",
		ProductCode: "FAST_INSTANT_TRADE_PAY",
		TotalAmount: 50,
		Subject:     "kubernetesv1.13.1",
	}
	pay.SignType = "RSA2"
	pay.Timestamp = time.Now().Format("2006-01-02 15:04:05")

	st := SortPay(pay, req)
	//生成签名
	sign := Sign(st)

	//追加参数
	pay.Sign = sign

	fmt.Println("encode: ", fmt.Sprintf("%s&sign=%s", URLEscape(pay, req), url.QueryEscape(sign)))

	//生成自动提交form
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
}
