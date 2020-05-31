package server

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/fanux/store/store/server/module"
	"github.com/fanux/store/store/server/vars"
	"io"
	"net/http"
)

func (u UserResource) callback(request *restful.Request, response *restful.Response) {
	code := request.QueryParameter("code")
	accessToken, err := module.GetGithubAccessToken(module.ClientID, module.ClientSecret, code)
	if err != nil {
		io.WriteString(response.ResponseWriter, "fetch token failed"+accessToken)
		return
	}
	user, err := module.GetUserInfo(accessToken)
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
		url = fmt.Sprintf("http://%s", vars.Domain)
		fmt.Println("login redirect rul is: ", url)
	} else {
		url = fmt.Sprintf("http://%s:%s%s?user=%s", vars.Domain, vars.BackPort, state, user.Login)
	}
	//io.WriteString(response.ResponseWriter, "this would be a normal response")
	//http.Redirect(response, request.Request, url+state, http.StatusMovedPermanently)
	http.Redirect(response, request.Request, url, http.StatusMovedPermanently)
}
