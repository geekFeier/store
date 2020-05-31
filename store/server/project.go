package server

import (
	"github.com/emicklei/go-restful"
	"github.com/fanux/store/store/server/module"
)

func classList(request *restful.Request, response *restful.Response) {
	p := &module.ProjectClass{}
	ps, err := p.List()
	if err != nil {
		response.WriteEntity(&Res{1, "list project class failed"})
	}
	response.WriteEntity(&ps)
}
