package server

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/fanux/store/store/server/module"
	"net/http"
)

// return RUL
func pay(request *restful.Request, response *restful.Response) {
	//TODO check sign
	login := request.PathParameter("login")
	productName := request.PathParameter("product")
	referrer := request.PathParameter(Referrer)

	fmt.Println("notify called", login, productName, referrer)

	//save up
	up := &module.UserProduct{
		Login:       login,
		ProductName: productName,
	}
	has, err := up.Get(login, productName)
	if err != nil {
		fmt.Println("get user pro error", err)
	}
	up.Referrer = referrer
	up.PayReferrer = module.GetProductDevide(productName)
	up.ProductPrice = module.GetProductPrice(productName)
	if !has {
		fmt.Println("can't find up")
	}
	up.Status = "payed"
	up.Referrer = referrer
	_, err = up.Update()
	if err != nil {
		fmt.Println("save up error", err)
	}

	upa := &module.UserPayeeAccount{
		Login: referrer,
	}
	has, err = upa.Get(referrer)
	if err != nil {
		fmt.Println("referrer pay error : ", err)
	} else {
		upa.Amount += up.PayReferrer
		if !has {
			upa.Save()
		} else {
			upa.Update()
		}
	}

	response.AddHeader("Content-Type", "application/x-gzip")
	http.Redirect(response, request.Request, module.GetProductURL(productName), http.StatusMovedPermanently)
}
