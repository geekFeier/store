package server

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/fanux/store/store/server/module"
	"io"
	"net/http"
)

// if check cookie failed, redirect to login page
func checkCookie(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.AddHeader("Cache-Control", "no-cache,no-store,must-revalidate")
	resp.AddHeader("Pragma", "no-cache")
	resp.AddHeader("Expires", "0")
	cookie, err := req.Request.Cookie("user")
	if err != nil || cookie == nil || len(cookie.Value) == 0 {
		fmt.Println("login please : ", err, req.Request.URL.String(), req.QueryParameter(Referrer))
		state := fmt.Sprintf("%s", req.Request.URL.String())
		http.Redirect(resp, req.Request, module.GetLoginURL(state), http.StatusMovedPermanently)
		return
	}
	chain.ProcessFilter(req, resp)
}

func (u UserResource) nop(request *restful.Request, response *restful.Response) {
	io.WriteString(response.ResponseWriter, "this would be a normal response")
}

//notify RUL
func notify(request *restful.Request, response *restful.Response) {
}
