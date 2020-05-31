package server

import (
	"fmt"
	"github.com/fanux/store/store/server/module"
	pay2 "github.com/fanux/store/store/server/pay"
	"net/http"
	"time"

	restful "github.com/emicklei/go-restful"
)

func productList(request *restful.Request, response *restful.Response) {
	p := &module.Product{}
	ps, err := p.List()
	if err != nil {
		response.WriteEntity(&Res{1, "list products failed"})
	}
	for i := range ps {
		ps[i].ProductURL = "www.sealyun.com"
	}
	response.WriteEntity(&ps)
}

// from server.go
func product(request *restful.Request, response *restful.Response) {
	referrer := request.QueryParameter(Referrer)
	productName := request.PathParameter("product")

	cookie, err := request.Request.Cookie("user")
	if err != nil {
		fmt.Println("Can't get cookie")
		return
	}

	login := cookie.Value
	fmt.Printf("user %s buy product %s", login, productName)

	//read user product in database by key (login , productname)
	up := &module.UserProduct{}
	var has bool
	has, err = up.Get(login, productName)
	fmt.Printf("user %s product %s status %s", login, productName, up.Status)
	if err != nil {
		fmt.Println("get userpro failed", err)
		return
	}
	prorand := fmt.Sprintf("%s-%s-%d", login, productName, time.Now().Unix())
	up.ID = prorand
	up.ClickCount++
	if !has {
		up.Login = login
		up.ProductName = productName
		up.Status = "see"
		_, err = up.Save()
		if err != nil {
			fmt.Println("save user product failed", err)
			return
		}
	}

	if up.Status == "payed" || isVip(login) {
		response.AddHeader("Content-Type", "application/x-gzip")
		http.Redirect(response, request.Request, module.GetProductURL(productName), http.StatusMovedPermanently)
		return
	}
	if referrer == "" && up.Referrer == "" {
		referrer = "fanux"
	}

	price := module.GetProductPrice(productName)
	//returnURL := fmt.Sprintf("pro/%s", productName)
	returnURL := fmt.Sprintf("/pro/pay/notify/%s/%s/%s", login, productName, referrer)
	notifyURL := fmt.Sprintf("/pro/pay/notify/%s/%s/%s", login, productName, referrer)
	payURL := pay2.PayURL(price, prorand, productName, GetFullURL(returnURL), GetFullURL(notifyURL))
	http.Redirect(response, request.Request, payURL, http.StatusMovedPermanently)
}
