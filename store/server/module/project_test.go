package module

import (
	"reflect"
	"testing"
)

func TestProject_ListFromClass(t *testing.T) {
	type fields struct {
		Name     string
		Describe string
		Owner    string
		VipPrice float64
		Docs     string
		Homepage string
		Class    string
	}
	type args struct {
		class string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Project
		wantErr bool
	}{
		{name: "fromClass", fields: fields{}, args: args{
			class: "cloud_kernel",
		}, want: []Project{
			{
				Name:     "kubernetes",
				Describe: "kubernetes一键安装",
				Owner:    "fanux",
				VipPrice: 69,
				Docs:     "xxx",
				Homepage: "sealyun.com",
				Class:    "cloud_kernel",
			},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Project{
				Name:     tt.fields.Name,
				Describe: tt.fields.Describe,
				Owner:    tt.fields.Owner,
				VipPrice: tt.fields.VipPrice,
				Docs:     tt.fields.Docs,
				Homepage: tt.fields.Homepage,
				Class:    tt.fields.Class,
			}
			got, err := p.ListFromClass(tt.args.class)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListFromClass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListFromClass() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProject_GetFromProject(t *testing.T) {
	type fields struct {
		Name     string
		Describe string
		Owner    string
		VipPrice float64
		Docs     string
		Homepage string
		Class    string
	}
	type args struct {
		class       string
		projectName string
		version     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []ProjectInfo
		wantErr bool
	}{
		{name: "fromProject", fields: fields{}, args: args{
			class:       "cloud_kernel",
			projectName: "kubernetes",
			version:     "1.14.4",
		}, want: []ProjectInfo{
			{
				Project: Project{
					Name:     "kubernetes",
					Describe: "kubernetes一键安装",
					Owner:    "fanux",
					VipPrice: 69,
					Docs:     "xxx",
					Homepage: "sealyun.com",
					Class:    "cloud_kernel",
				},
				Product: Product{
					ProductName:   "kubernetes1.14.4",
					ProductURL:    "https://sealyun.oss-cn-beijing.aliyuncs.com/9763df396eeb1359fa9d0a34de64e8a7-1.14.4/kube1.14.4.tar.gz",
					ProductPrice:  50,
					ProductDivide: 0.6,
				},
				ProjectVersion: ProjectVersion{
					ProjectName: "kubernetes",
					ProductName: "kubernetes1.14.4",
					Version:     "1.14.4",
				},
			},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Project{
				Name:     tt.fields.Name,
				Describe: tt.fields.Describe,
				Owner:    tt.fields.Owner,
				VipPrice: tt.fields.VipPrice,
				Docs:     tt.fields.Docs,
				Homepage: tt.fields.Homepage,
				Class:    tt.fields.Class,
			}
			got, err := p.GetFromProject(tt.args.class, tt.args.projectName, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFromProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFromProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}
