package serve

import (
	"github.com/fanux/store/store/server/module"

	restful "github.com/emicklei/go-restful"
)

func productList(request *restful.Request, response *restful.Response) {
	p := &module.Product{}
	ps, err := p.List()
	if err != nil {
		response.WriteEntity(&Res{1, "list products failed"})
	}
	for i := range ps {
		ps[i].ProductURL = "www.sealyun.com"
	}
	response.WriteEntity(&ps)
}
