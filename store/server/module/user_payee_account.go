package module

//UserPayeeAccount use for alipay
type UserPayeeAccount struct {
	Login        string
	PayeeAccount string
	Amount       float64 //user earned money
	Passwd       string  //
}

//Get is
func (upa *UserPayeeAccount) Get(login string) (bool, error) {
	return engine.Where("login = ?", login).Get(upa)
}

//Save is
func (upa *UserPayeeAccount) Save() (int64, error) {
	return engine.Insert(upa)
}

//Update is
func (upa *UserPayeeAccount) Update() (int64, error) {
	return engine.Where("login = ?", upa.Login).Update(upa)
}
