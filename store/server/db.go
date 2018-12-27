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

//Product is
type Product struct {
	ProductName string

	ProductURL   string // http://sealyun.oss-cn-beijing.aliyuncs.com/c89602f7cb2a/kube1.13.1.tar.gz
	ProductPrice float64
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
}

//CreateTables is
func CreateTables() {
	err := engine.CreateTables(new(User))
	err = engine.CreateTables(new(Product))
	err = engine.CreateTables(new(UserProduct))
	err = engine.CreateTables(new(UserPayeeAccount))

	if err != nil {
		fmt.Println("new table failed", err)
	}
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
func (up *UserProduct) Save() (int64, error) {
	return engine.Insert(up)
}

//Get is
func (up *UserProduct) Get(login, product string) (bool, error) {
	return engine.Where("login = ?", login).And("product_name = ?", product).Get(up)
}

//Save is
func (p *Product) Save() (int64, error) {
	return engine.Insert(p)
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

func init() {
	GetEngine()
	CreateTables()

	p := &Product{
		ProductName:  "kubernetes1.13.1",
		ProductURL:   "http://sealyun.oss-cn-beijing.aliyuncs.com/c89602f7cb2a/kube1.13.1.tar.gz",
		ProductPrice: 0.01,
	}
	_, err := p.Save()
	if err != nil {
		fmt.Println("save product failed")
	}
}
