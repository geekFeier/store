package module

import "fmt"

type ProjectClass struct {
	Name     string //名称
	Alias    string //别名
	Label    string //标签
	Describe string //描述
}

//List is
func (p *ProjectClass) List() ([]ProjectClass, error) {
	var classes []ProjectClass

	err := engine.Find(&classes)
	if err != nil {
		return classes, fmt.Errorf("list project class failed %v", err)
	}
	return classes, nil
}
