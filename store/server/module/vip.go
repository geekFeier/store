package module

//VIP is
type VIP struct {
	Login string
	Price float64
	Date  int64
	Type  string // normal super
}

//Save is
func (vip *VIP) Save() (int64, error) {
	if ok, _ := vip.Get(vip.Login); ok {
		return 0, nil
	}
	return engine.Insert(vip)
}

//Get is
func (vip *VIP) Get(login string) (bool, error) {
	return engine.Where("login = ?", login).Get(vip)
}

//Update is
func (vip *VIP) Update() (int64, error) {
	return engine.Where("login = ?", vip.Login).Update(vip)
}
