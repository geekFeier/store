package module

//PayedUser is
type PayedUser struct {
	User `xorm:"extends"`
	Name string
}

//TableName is
func (PayedUser) TableName() string {
	return "user"
}

func InUserList(u PayedUser, list []PayedUser) bool {
	for _, ul := range list {
		if ul.User.Login == u.User.Login {
			return true
		}
	}

	return false
}

func PayedUserLoginList(users []PayedUser) ([]PayedUser, error) {
	//err := engine.Join("INNER", "user_product", "user_product.login = user.login").Where("status = ?", "payed").And("product_name = ?", productName).Find(&users)
	//err := engine.Join("INNER", "user_product", "user_product.login = user.login").Where("status = ?", "payed").Find(&users)
	err := engine.Join("INNER", "user_product", "user_product.login = user.login").Find(&users)
	return users, err
}
