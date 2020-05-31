package module

import "fmt"

type ProjectVersion struct {
	ProjectName string //项目名称
	ProductName string //商品名称
	Version     string //版本
}

//List is
func (p *ProjectVersion) ListFromProjectName(projectName, class string) ([]ProjectVersion, error) {
	var versions []ProjectVersion

	err := engine.
		Where("project_name= ?", projectName).
		Join("left", "project", "project.name = project_version.project_name").
		Where("project.class = ?", class).
		Find(&versions)
	if err != nil {
		return versions, fmt.Errorf("list project versions failed %v", err)
	}
	return versions, nil
}
