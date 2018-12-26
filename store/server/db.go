package serve

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

//GetEngine is
func GetEngine() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "/data/stroe.db")
	if err != nil {
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
	Login       string
	ProductName string

	Referrer    string
	PayReferrer float64

	Status string
}

//CreateTables is
func CreateTables() {
	err := engine.Sync2(new(User))
	err := engine.Sync2(new(UserProduct))
}
