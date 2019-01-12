package serve

import (
	"fmt"

	restful "github.com/emicklei/go-restful"
)

func productList(request *restful.Request, response *restful.Response) {
	p := &Product{}
	ps, err := p.List()
	if err != nil {
		response.WriteEntity(&Res{1, "list products failed"})
	}
	response.WriteEntity(&ps)
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
