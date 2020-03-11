package serve

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

//GetEngine is
func GetEngine() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:Fanux#123@/store?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic("start engine failed")
	}
}

//VIP is
type VIP struct {
	Login string
	Price float64
	Date  int64
	Type  string // normal super
}

//Product is
type Product struct {
	ProductName string

	ProductURL   string // http://sealyun.oss-cn-beijing.aliyuncs.com/c89602f7cb2a/kube1.13.1.tar.gz
	ProductPrice float64

	ProductDivide float64
}

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

//UserPayeeAccount use for alipay
type UserPayeeAccount struct {
	Login        string
	PayeeAccount string
	Amount       float64 //user earned money
	Passwd       string  //
}

//CreateTables is
func CreateTables() {
	err := engine.CreateTables(new(User))
	err = engine.CreateTables(new(VIP))
	err = engine.CreateTables(new(Product))
	err = engine.CreateTables(new(UserProduct))
	err = engine.CreateTables(new(UserPayeeAccount))

	err = engine.Sync(new(User))
	err = engine.Sync(new(VIP))
	err = engine.Sync(new(Product))
	err = engine.Sync(new(UserProduct))
	err = engine.Sync(new(UserPayeeAccount))

	if err != nil {
		fmt.Println("new table failed", err)
	}
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

//Save is
func (user *User) Save() (int64, error) {
	return engine.Insert(user)
}

//Get is
func (user *User) Get(login string) (bool, error) {
	return engine.Where("login = ?", login).Get(user)
}

//Save is
func (vip *VIP) Save() (int64, error) {
	if ok,_:= vip.Get(vip.Login);ok {
		return 0,nil
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

//Save is
func (p *Product) Save() (int64, error) {
	return engine.Insert(p)
}

//Update is
func (p *Product) Update() (int64, error) {
	return engine.Where("product_name = ?", p.ProductName).Update(p)
}

//Get is
func (p *Product) Get(name string) (bool, error) {
	return engine.Where("product_name= ?", name).Get(p)
}

//GetProductURL is
func GetProductURL(name string) string {
	p := &Product{ProductName: name}
	p.Get(name)
	return p.ProductURL
}

//GetProductPrice is
func GetProductPrice(name string) float64 {
	p := &Product{ProductName: name}
	p.Get(name)
	return p.ProductPrice
}

//GetProductDevide is
func GetProductDevide(name string) float64 {
	p := &Product{ProductName: name}
	has, err := p.Get(name)
	if !has || err != nil {
		fmt.Println("get devide failed")
		return 0
	}
	return p.ProductPrice * p.ProductDivide
}

func saveProduct(p *Product) {
	has, err := p.Get(p.ProductName)
	if err != nil {
		fmt.Println("get product info failed")
	}
	if has {
		p.Update()
	} else {
		//TODO use update not save
		_, err := p.Save()
		if err != nil {
			fmt.Println("save product failed")
		}
	}
}

func init() {
	GetEngine()
	CreateTables()

	/*
	p := &Product{
		ProductName:   "kubernetes1.13.1",
		ProductURL:    "http://sealyun.oss-cn-beijing.aliyuncs.com/c89602f7cb2a/kube1.13.1.tar.gz",
		ProductPrice:  50,
		ProductDivide: 0.6,
	}
	saveProduct(p)

	p = &Product{
		ProductName:   "kubernetes1.13.2",
		ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/tk2hhcegu9Z1-13-2/kube1.13.2.tar.gz",
		ProductPrice:  50,
		ProductDivide: 0.6,
	}
	saveProduct(p)

	p = &Product{
		ProductName:   "kubernetes1.13.3",
		ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/1172b02fc1-13-3/kube1.13.3.tar.gz",
		ProductPrice:  50,
		ProductDivide: 0.6,
	}
	saveProduct(p)

	p = &Product{
		ProductName:   "kubernetes1.13.4",
		ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/3358d153a6381-13-4/kube1.13.4.tar.gz",
		ProductPrice:  50,
		ProductDivide: 0.6,
	}
	saveProduct(p)

	p = &Product{
		ProductName:   "kubernetes1.14.0",
		ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/5d36974951c81-14-0/kube1.14.0.tar.gz",
		ProductPrice:  50,
		ProductDivide: 0.6,
	}
	saveProduct(p)

	p = &Product{
		ProductName:   "kubernetes1.14.1",
		ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/6300ecea5ab9621-14-1/kube1.14.1.tar.gz",
		ProductPrice:  0.01,
		ProductDivide: 0.6,
	}
	saveProduct(p)

	p = &Product{
		ProductName:   "kubernetes1.14.2",
		ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/kasjfjkadf1-14-2/kube1.14.2.tar.gz",
		ProductPrice:  50,
		ProductDivide: 0.6,
	}
	saveProduct(p)

	p = &Product{
		ProductName:   "kubernetes1.14.3",
		ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/kasjfjkadf1-14-2/kube1.14.3.tar.gz",
		ProductPrice:  50,
		ProductDivide: 0.6,
	}
	saveProduct(p)

	p = &Product{
		ProductName:   "kubernetes1.15.0",
		ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/free/kube1.15.0.tar.gz",
		ProductPrice:  0.01,
		ProductDivide: 0.6,
	}
	saveProduct(p)

	p = &Product{
		ProductName:   "kubernetes1.14.4",
		ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/9763df396eeb1359fa9d0a34de64e8a7-1.14.4/kube1.14.4.tar.gz",
		ProductPrice:  50,
		ProductDivide: 0.6,
	}
	saveProduct(p)

	p = &Product{
		ProductName:   "kubernetes1.15.1",
		ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/502777a359bf20b89fad1c8565f9fbac-1.15.1/kube1.15.1.tar.gz",
		ProductPrice:  50,
		ProductDivide: 0.6,
	}
	saveProduct(p)
	*/
}
