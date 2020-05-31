package server

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/fanux/store/store/server/vars"
	"log"
	"net/http"
)

// Cross-origin resource sharing (CORS) is a mechanism that allows JavaScript on a web page
// to make XMLHttpRequests to another domain, not the domain the JavaScript originated from.
//
// http://en.wikipedia.org/wiki/Cross-origin_resource_sharing
// http://enable-cors.org/server.html
//
// GET http://localhost:8080/users
//
// GET http://localhost:8080/users/1
//
// PUT http://localhost:8080/users/1
//
// DELETE http://localhost:8080/users/1
//
// OPTIONS http://localhost:8080/users/1  with Header "Origin" set to some domain and

//UserResource  s
type UserResource struct{}

//const
const (
	Referrer = "referrer"
)

//Res is
type Res struct {
	Code   int
	Reason string
}

func (r Res) String() string {
	return fmt.Sprintf("{Code:%d,Reason:%s}", r.Code, r.Reason)
}

//GetFullURL is
func GetFullURL(path string) string {
	return fmt.Sprintf("http://%s:%s/%s", vars.Domain, vars.BackPort, path)
}

//RegisterTo is
func (u UserResource) RegisterTo(container *restful.Container) {
	loginless := new(restful.WebService)
	loginless.
		Path("").
		Consumes("*/*").
		Produces(restful.MIME_JSON)
	loginless.Route(loginless.GET("/callback").To(u.callback))
	loginless.Route(loginless.GET("/loginless/pro/{product}/payed").To(payedUserList))
	loginless.Route(loginless.GET("/loginless/pro").To(productList))
	loginless.Route(loginless.GET("/loginless/user/payee").To(userPayeeInfo))
	loginless.Route(loginless.GET("/loginless/info/user").To(userInfo))
	loginless.Route(loginless.GET("/loginless/vip/notify/{login}").To(vipChargeNotify))
	loginless.Route(loginless.POST("/loginless/vip/notify/{login}").To(vipChargeNotify))
	// /class/{class_name}/project/{project_name}/version/{version}
	loginless.Route(loginless.GET("/loginless/class").To(classList))
	loginless.Route(loginless.GET("/loginless/class/{class_name}/project").To(projectList))
	loginless.Route(loginless.GET("/loginless/class/{class_name}/project/{project_name}/version").To(versionList))

	ws := new(restful.WebService)
	ws.
		Path("/pro").
		Consumes("*/*").
		Produces("*/*")
	ws.Filter(checkCookie)
	ws.Route(ws.GET("/pay/notify/{login}/{product}/{referrer}").To(pay))
	ws.Route(ws.POST("/pay/notify/{login}/{product}/{referrer}").To(notify))
	ws.Route(ws.POST("").To(u.nop))
	ws.Route(ws.PUT("/{user-id}").To(u.nop))
	ws.Route(ws.DELETE("/{user-id}").To(u.nop))
	ws.Route(ws.GET("/{product}").To(product))

	user := new(restful.WebService)
	user.Path("/user").
		Consumes("*/*").
		Produces(restful.MIME_JSON)
	user.Filter(checkCookie)
	user.Route(user.GET("/info").To(userInfo))
	user.Route(user.GET("/info/payee").To(userPayeeInfo))
	user.Route(user.PUT("/info/payee").To(updateUserPayeeInfo))
	user.Route(user.POST("/info/withdraw").To(userWithdraw))
	user.Route(user.GET("/vip/charge").To(vipCharge))

	container.Add(ws)
	container.Add(user)
	container.Add(loginless)
}

//Run is
func Run() {
	wsContainer := restful.NewContainer()
	u := UserResource{}
	u.RegisterTo(wsContainer)

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT"},
		CookiesAllowed: true,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	log.Print("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

/*
	var client = alipay.New("2018121662557851", aliPublicKey, privateKey, true)

	var p = alipay.AliPayTradeWapPay{}
	p.NotifyURL = "http://xxx"
	p.ReturnURL = "http://xxx"
	p.Subject = "k8s"
	p.OutTradeNo = "asfsfdaf"
	p.TotalAmount = "10.00"
	p.ProductCode = "1231231"

	var url, err = client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	var payURL = url.String()
	fmt.Println(payURL)
*/
