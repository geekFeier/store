package serve

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/fanux/store/store/server/module"
	pay2 "github.com/fanux/store/store/server/pay"
	"io"
	"net/http"
	"time"
)

//from server.go
func vipChargeNotify(request *restful.Request, response *restful.Response) {
	login := request.PathParameter("login")
	time := time.Now().Unix()

	vip := &module.VIP{Login: login}
	vip.Type = "nomal"
	vip.Date = time
	vip.Update()
	io.WriteString(response.ResponseWriter, "充值年费会员成功，一年内享有免费下载所有软件包权利")
}
func vipCharge(request *restful.Request, response *restful.Response) {
	cookie, err := request.Request.Cookie("user")
	if err != nil {
		fmt.Println("Can't get cookie")
		return
	}
	login := cookie.Value
	var price float64
	price = 69
	vip := &module.VIP{}
	vip.Login = login
	vip.Price = price

	if isVip(login) {
		response.WriteEntity(&Res{1, "您已经是会员了，还要再充一次，钱多吗？"})
		return
	}

	vip.Save()

	productName := "sealyunvip"
	viprand := fmt.Sprintf("%s-%d", login, time.Now().Unix())
	fmt.Printf("vip %s charge", login)
	returnURL := fmt.Sprintf("/loginless/vip/notify/%s", login)
	notifyURL := fmt.Sprintf("/loginless/vip/notify/%s", login)
	payURL := pay2.PayURL(price, viprand, productName, GetFullURL(returnURL), GetFullURL(notifyURL))
	http.Redirect(response, request.Request, payURL, http.StatusMovedPermanently)
}

// from vip.go
func overTime(t int64, years, months, days int) bool {
	now := time.Now().Unix()
	after := time.Unix(t, 0).AddDate(years, months, days).Unix()

	fmt.Printf("Now %d a years later %d, vip time %d", now, after, t)
	if after <= now {
		fmt.Println("vip not past due")
		return true
	}
	fmt.Println("vip past due")
	return false
}

func isVip(login string) bool {
	vip := &module.VIP{Login: login}
	ok, err := vip.Get(login)
	if err != nil || !ok {
		fmt.Printf("get vip %s failed", login)
		return false
	}
	if overTime(vip.Date, 1, 0, 0) {
		return false
	}

	return true
}
