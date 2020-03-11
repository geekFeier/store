package serve

import restful "github.com/emicklei/go-restful"

//PayedUser is
type PayedUser struct {
	User `xorm:"extends"`
	Name string
}

//TableName is
func (PayedUser) TableName() string {
	return "user"
}

func inUserList(u PayedUser, list []PayedUser) bool {
	for _, ul := range list {
		if ul.User.Login == u.User.Login {
			return true
		}
	}

	return false
}

func payedUserList(request *restful.Request, response *restful.Response) {
	productName := request.PathParameter("product")

	var users []PayedUser
	var usersUni []PayedUser
	//err := engine.Join("INNER", "user_product", "user_product.login = user.login").Where("status = ?", "payed").And("product_name = ?", productName).Find(&users)
	//err := engine.Join("INNER", "user_product", "user_product.login = user.login").Where("status = ?", "payed").Find(&users)
	err := engine.Join("INNER", "user_product", "user_product.login = user.login").Find(&users)
	_ = productName

	c := 0
	for _, u := range users {
		if inUserList(u, usersUni) {
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
