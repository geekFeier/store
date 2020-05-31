package module

import "fmt"

type Project struct {
	Name     string  //名字
	Describe string  //描述
	Owner    string  //作者
	VipPrice float64 //vip价格
	Docs     string  //文档
	Homepage string  //主页
	Class    string  //类别
}

//List is
func (p *Project) ListFromClass(class string) ([]Project, error) {
	var projects []Project
	err := engine.Where("class= ?", class).Find(&projects)
	if err != nil {
		return projects, fmt.Errorf("list project failed %v", err)
	}
	return projects, nil
}

type ProjectInfo struct {
	Project        `xorm:"extends"`
	ProjectVersion `xorm:"extends"`
	Product        `xorm:"extends"`
}

func (ProjectInfo) TableName() string {
	return "project"
}

//engine.SQL("" ).Find()
func (p *Project) ListVersionsFromProject(class, projectName string) ([]ProjectInfo, error) {
	var projects []ProjectInfo
	err := engine.
		Join("left", "project_version", "project_version.project_name=project.name").
		Join("left", "product", "project_version.product_name=product.product_name").
		Where("project.class=?", class).
		Where("project.name=?", projectName).
		OrderBy("project_version.version desc").
		//Where("project_version.version=?", version).
		Find(&projects)
	if err != nil {
		return projects, fmt.Errorf("list project failed %v", err)
	}
	return projects, nil
}
