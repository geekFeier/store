package module

import (
	"fmt"
	"github.com/fanux/store/store/server/vars"
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

	//产品相关
	err = engine.Sync(new(Project))
	err = engine.Sync(new(ProjectVersion))
	err = engine.Sync(new(ProjectClass))
	if err != nil {
		fmt.Println("new table failed", err)
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
