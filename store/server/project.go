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

func projectList(request *restful.Request, response *restful.Response) {
	className := request.PathParameter("class_name")
	p := &module.Project{}
	ps, err := p.ListFromClass(className)
	if err != nil {
		response.WriteEntity(&Res{1, "list project from class name failed"})
	}
	response.WriteEntity(&ps)
}

func versionList(request *restful.Request, response *restful.Response) {
	className := request.PathParameter("class_name")
	projectName := request.PathParameter("project_name")
	p := &module.ProjectInfo{}
	ps, err := p.ListVersionsFromProject(className, projectName)
	if err != nil {
		response.WriteEntity(&Res{1, "list project version from class and class name failed"})
	}
	response.WriteEntity(&ps)
}
