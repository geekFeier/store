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
const (
	Referrer = "referrer"
)

//const
var (
	Domain   = "store.lameleg.com"
	BackPort = "8080"
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
	return fmt.Sprintf("http://%s:%s/%s", Domain, BackPort, path)
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

func userWithdraw(request *restful.Request, response *restful.Response) {
	// TODO if passwd or payee accoun is null redirect to PUT user payee
	cookie, err := request.Request.Cookie("user")
	if err != nil {
		fmt.Println("Can't get cookie")
		return
	}

	upa := &UserPayeeAccount{Login: cookie.Value}
	has, err := upa.Get(cookie.Value)
	if err != nil {
		fmt.Println("Can't get user payee account")
		return
	}

	if !has {
		fmt.Println("uer payee account not found")
	}

	if upa.PayeeAccount == "" || upa.Amount == 0 || upa.Passwd == "" {
		response.WriteEntity(upa)
		return
	}

	//TODO do withdraw
}
func userInfo(request *restful.Request, response *restful.Response) {
	fmt.Println("user info called")
	cookie, err := request.Request.Cookie("user")
	if err != nil {
		fmt.Println("Can't get cookie")
		return
	}
	/*
		u := request.PathParameter("user")
		if u != cookie.Value {
			fmt.Printf("user %s not Equal cookie %s", u, cookie.Value)
		}
	*/

	user := &User{}
	has, err := user.Get(cookie.Value)
	if err != nil {
		io.WriteString(response.ResponseWriter, "get user info failed")
		return
	}
	if has {
		response.WriteEntity(user)
		return
	}
	io.WriteString(response.ResponseWriter, "get user not found")
	return
}
func userPayeeInfo(request *restful.Request, response *restful.Response) {
	cookie, err := request.Request.Cookie("user")
	if err != nil {
		fmt.Println("Can't get cookie")
		return
	}
	/*
		u := request.PathParameter("user")
		if u != cookie.Value {
			fmt.Printf("user %s not Equal cookie %s", u, cookie.Value)
		}
	*/

	upa := &UserPayeeAccount{
		Login: cookie.Value,
	}
	has, err := upa.Get(cookie.Value)
	if err != nil {
		io.WriteString(response.ResponseWriter, "get user payee account info failed")
		return
	}
	if has {
		upa.Passwd = "want passwd? no way!"
		response.WriteEntity(upa)
		return
	}
	io.WriteString(response.ResponseWriter, "user payee account not found")
	return
}
func updateUserPayeeInfo(request *restful.Request, response *restful.Response) {
	cookie, err := request.Request.Cookie("user")
	if err != nil {
		response.WriteEntity(&Res{1, "Can't get cookie"})
		return
	}

	upa := &UserPayeeAccount{}
	err = request.ReadEntity(upa)
	if err != nil {
		response.WriteEntity(&Res{1, "get user payee account failed"})
		return
	}
	upa.Login = cookie.Value

	/*
		if upa.Login != cookie.Value {
			io.WriteString(response.ResponseWriter, "cookie not equal to upa.Login")
			return
		}
	*/

	upaDB := &UserPayeeAccount{}
	has, err := upaDB.Get(upa.Login)
	if err != nil {
		response.WriteEntity(&Res{1, "get user payee account info failed"})
		return
	}
	if has {
		//update
		if upa.PayeeAccount != "" {
			upaDB.PayeeAccount = upa.PayeeAccount
		}
		if upa.Amount != 0 {
			upaDB.Amount = upa.Amount
		}
		if upa.Passwd != "" {
			upaDB.Passwd = upa.Passwd
		}
		_, err = upaDB.Update()
		if err != nil {
			response.WriteEntity(&Res{1, "update user payee account info failed"})
			return
		}
	} else {
		//create
		_, err = upa.Save()
		if err != nil {
			response.WriteEntity(&Res{1, "save user payee account info failed"})
			return
		}
	}
	response.WriteEntity(&Res{0, "save user payee account successed"})
}

// if check cookie failed, redirect to login page
func checkCookie(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.AddHeader("Cache-Control", "no-cache,no-store,must-revalidate")
	resp.AddHeader("Pragma", "no-cache")
	resp.AddHeader("Expires", "0")
	cookie, err := req.Request.Cookie("user")
	if err != nil || cookie == nil {
		fmt.Println("login please : ", err, req.Request.URL.String(), req.QueryParameter(Referrer))
		state := fmt.Sprintf("%s", req.Request.URL.String())
		http.Redirect(resp, req.Request, GetLoginURL(state), http.StatusMovedPermanently)
		return
	}
	chain.ProcessFilter(req, resp)
}

func vipChargeNotify(request *restful.Request, response *restful.Response) {
	login := request.PathParameter("login")
	time := time.Now().Unix()

	vip := &VIP{Login: login}
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
	vip := &VIP{}
	vip.Login = login
	vip.Price = price

	if isVip(login) {
		response.WriteString("您已经是会员了，还要再充一次，钱多吗？")
		return
	}

	vip.Save()

	productName := "sealyunvip"
	viprand := fmt.Sprintf("%s-%d", login, time.Now().Unix())
	fmt.Printf("vip %s charge", login)
	returnURL := fmt.Sprintf("/loginless/vip/notify/%s", login)
	notifyURL := fmt.Sprintf("/loginless/vip/notify/%s", login)
	payURL := PayURL(price, viprand, productName, GetFullURL(returnURL), GetFullURL(notifyURL))
	http.Redirect(response, request.Request, payURL, http.StatusMovedPermanently)
}

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
	referrer := request.PathParameter(Referrer)

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
	up.PayReferrer = GetProductDevide(productName)
	up.ProductPrice = GetProductPrice(productName)
	if !has {
		fmt.Println("can't find up")
	}
	up.Status = "payed"
	up.Referrer = referrer
	_, err = up.Update()
	if err != nil {
		fmt.Println("save up error", err)
	}

	upa := &UserPayeeAccount{
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
	http.Redirect(response, request.Request, GetProductURL(productName), http.StatusMovedPermanently)
}

func (u UserResource) callback(request *restful.Request, response *restful.Response) {
	code := request.QueryParameter("code")
	accessToken, err := GetGithubAccessToken(clientID, clientSecret, code)
	if err != nil {
		io.WriteString(response.ResponseWriter, "fetch token failed"+accessToken)
		return
	}
	user, err := GetUserInfo(accessToken)
	if err != nil {
		fmt.Println(err)
		io.WriteString(response.ResponseWriter, fmt.Sprintf("get user info failed %s", err))
		return
	}

	has, err := user.Get(user.Login)
	if err != nil {
		fmt.Println("get user failed: ", user.Login)
		io.WriteString(response.ResponseWriter, fmt.Sprintf("query user info failed %s", err))
		return
	}
	if !has {
		//has,err := user.Get(user.Login)
		_, err = user.Save()
		if err != nil {
			fmt.Println("save suer faieled: ", err)
			io.WriteString(response.ResponseWriter, fmt.Sprintf("save user info failed %s", err))
			return
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
		fmt.Println("login redirect rul is: ", url)
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
