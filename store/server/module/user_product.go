package module

//UserProduct is
type UserProduct struct {
	ID           string
	Login        string
	ProductName  string
	ProductPrice float64

	Referrer    string
	PayReferrer float64

	Status     string // [see,payed,unknow]
	ClickCount int
}

//Save is
func (up *UserProduct) Save() (int64, error) {
	return engine.Insert(up)
}

//Update is
func (up *UserProduct) Update() (int64, error) {
	return engine.Where("login = ?", up.Login).And("product_name = ?", up.ProductName).Update(up)
}

//Get is
func (up *UserProduct) Get(login, product string) (bool, error) {
	return engine.Where("login = ?", login).And("product_name = ?", product).Get(up)
}
