package serve

import (
	restful "github.com/emicklei/go-restful"
	"github.com/fanux/store/store/server/module"
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
