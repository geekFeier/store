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

func payedUserList(request *restful.Request, response *restful.Response) {
	productName := request.PathParameter("product")

	var users []PayedUser
	//err := engine.Join("INNER", "user_product", "user_product.login = user.login").Where("status = ?", "payed").And("product_name = ?", productName).Find(&users)
	err := engine.Join("INNER", "user_product", "user_product.login = user.login").Where("status = ?", "payed").Find(&users)
	_ = productName

	if err != nil {
		response.WriteEntity(&Res{1, "get payed user list failed"})
	}
	response.WriteEntity(&users)
}
