package module

import "fmt"

//Product is
type Product struct {
	ProductName string

	ProductURL   string // http://sealyun.oss-cn-beijing.aliyuncs.com/c89602f7cb2a/kube1.13.1.tar.gz
	ProductPrice float64

	ProductDivide float64
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

//List is
func (p *Product) List() ([]Product, error) {
	var products []Product

	err := engine.Find(&products)
	if err != nil {
		return products, fmt.Errorf("List product failed %s", err)
	}
	return products, nil
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
