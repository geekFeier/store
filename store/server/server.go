package serve

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
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
var (
	Domain   = "store.lameleg.com"
	BackPort = "8080"
)

//GetFullURL is
func GetFullURL(path string) string {
	return fmt.Sprintf("http://%s:%s/%s", Domain, BackPort, path)
}

//RegisterTo is
func (u UserResource) RegisterTo(container *restful.Container) {
	loginless := new(restful.WebService)
	loginless.
		Path("").
		Consumes("*/*").
		Produces("*/*")
	loginless.Route(loginless.GET("/callback").To(u.callback))

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

	container.Add(ws)
	container.Add(loginless)
}

// if check cookie failed, redirect to login page
func checkCookie(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.AddHeader("Cache-Control", "no-cache,no-store,must-revalidate")
	resp.AddHeader("Pragma", "no-cache")
	resp.AddHeader("Expires", "0")
	cookie, err := req.Request.Cookie("user")
	if err != nil || cookie == nil {
		fmt.Println("login please : ", err, req.Request.URL.String())
		http.Redirect(resp, req.Request, GetLoginURL(req.Request.URL.String()), http.StatusMovedPermanently)
		return
	}
	chain.ProcessFilter(req, resp)
}

func product(request *restful.Request, response *restful.Response) {
	referrer := request.QueryParameter("referrer")
	productName := request.PathParameter("product")

	cookie, err := request.Request.Cookie("user")
	if err != nil {
		fmt.Println("Can't get cookie")
		return
	}

	login := cookie.Value
	fmt.Printf("user %s buy product %s", login, productName)

	//read user product in database by key (login , productname)
	up := &UserProduct{}
	var has bool
	has, err = up.Get(login, productName)
	fmt.Printf("user %s product %s status %s", login, productName, up.Status)
	if err != nil {
		fmt.Println("get userpro failed", err)
		return
	}
	prorand := fmt.Sprintf("%s-%s-%d", login, productName, time.Now().Unix())
	up.ID = prorand
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

	if up.Status == "payed" {
		response.AddHeader("Content-Type", "application/x-gzip")
		http.Redirect(response, request.Request, GetProductURL(productName), http.StatusMovedPermanently)
		return
	}
	if referrer == "" && up.Referrer == "" {
		referrer = "fanux"
	}

	price := GetProductPrice(productName)
	//returnURL := fmt.Sprintf("pro/%s", productName)
	returnURL := fmt.Sprintf("/pro/pay/notify/%s/%s/%s", login, productName, referrer)
	notifyURL := fmt.Sprintf("/pro/pay/notify/%s/%s/%s", login, productName, referrer)
	payURL := PayURL(price, prorand, productName, GetFullURL(returnURL), GetFullURL(notifyURL))
	http.Redirect(response, request.Request, payURL, http.StatusMovedPermanently)
}

func (u UserResource) nop(request *restful.Request, response *restful.Response) {
	io.WriteString(response.ResponseWriter, "this would be a normal response")
}

//notify RUL
func notify(request *restful.Request, response *restful.Response) {
}

// return RUL
func pay(request *restful.Request, response *restful.Response) {
	//TODO check sign
	login := request.PathParameter("login")
	productName := request.PathParameter("product")
	referrer := request.PathParameter("referrer")

	fmt.Println("notify called", login, productName, referrer)

	//save up
	up := &UserProduct{
		Login:       login,
		ProductName: productName,
	}
	has, err := up.Get(login, productName)
	if err != nil {
		fmt.Println("get user pro error", err)
	}
	up.Referrer = referrer
	if !has {
		fmt.Println("can't find up")
	}
	up.Status = "payed"
	up.Referrer = referrer
	_, err = up.Update()
	if err != nil {
		fmt.Println("save up error", err)
	}

	response.AddHeader("Content-Type", "application/x-gzip")
	http.Redirect(response, request.Request, GetProductURL(productName), http.StatusMovedPermanently)
}

func (u UserResource) callback(request *restful.Request, response *restful.Response) {
	code := request.QueryParameter("code")
	accessToken, err := GetGithubAccessToken(clientID, clientSecret, code)
	if err != nil {
		io.WriteString(response.ResponseWriter, "fetch token failed"+accessToken)
	}
	user, err := GetUserInfo(accessToken)
	if err != nil {
		fmt.Println(err)
	}

	has, err := user.Get(user.Login)
	if !has || err != nil {
		//has,err := user.Get(user.Login)
		_, err = user.Save()
		if err != nil {
			fmt.Println("save suer faieled: ", err)
		}
	}

	// Set cookie
	cookie := http.Cookie{Name: "user", Value: user.Login, Path: "/", MaxAge: 86400}
	http.SetCookie(response, &cookie)

	request.Request.AddCookie(&cookie)

	state := request.QueryParameter("state")
	fmt.Println("redirect url is : ", state)
	//redirect back to user request
	var url string
	if state == "" {
		//TODO return to home page
		url = fmt.Sprintf("http://%s", Domain)
	} else {
		url = fmt.Sprintf("http://%s:%s%s?user=%s", Domain, BackPort, state, user.Login)
	}
	//io.WriteString(response.ResponseWriter, "this would be a normal response")
	//http.Redirect(response, request.Request, url+state, http.StatusMovedPermanently)
	http.Redirect(response, request.Request, url, http.StatusMovedPermanently)
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
		AllowedMethods: []string{"GET", "POST"},
		CookiesAllowed: false,
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
