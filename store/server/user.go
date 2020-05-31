package server

import (
	"fmt"
	restful "github.com/emicklei/go-restful"
	"github.com/fanux/store/store/server/module"
	"io"
)

func payedUserList(request *restful.Request, response *restful.Response) {
	productName := request.PathParameter("product")

	var users []module.PayedUser
	var usersUni []module.PayedUser
	users, err := module.PayedUserLoginList(users)
	_ = productName

	c := 0
	for _, u := range users {
		if module.InUserList(u, usersUni) {
			continue
		}
		c++
		usersUni = append(usersUni, u)
		if c > 250 {
			break
		}
	}

	if err != nil {
		response.WriteEntity(&Res{1, "get payed user list failed"})
	}
	response.WriteEntity(&usersUni)
}

// from server.go

func userWithdraw(request *restful.Request, response *restful.Response) {
	// TODO if passwd or payee accoun is null redirect to PUT user payee
	cookie, err := request.Request.Cookie("user")
	if err != nil {
		fmt.Println("Can't get cookie")
		return
	}

	upa := &module.UserPayeeAccount{Login: cookie.Value}
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

	user := &module.User{}
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

	upa := &module.UserPayeeAccount{
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

	upa := &module.UserPayeeAccount{}
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

	upaDB := &module.UserPayeeAccount{}
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
