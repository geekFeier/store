package module

import (
	"reflect"
	"testing"
)

func TestProjectVersion_ListFromProjectName(t *testing.T) {
	type fields struct {
		ProjectName string
		ProductName string
		Version     string
	}
	type args struct {
		projectName string
		class       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []ProjectVersion
		wantErr bool
	}{
		{name: "fromProjectName", fields: fields{}, args: args{
			projectName: "kubernetes",
			class:       "cloud_kernel",
		}, want: []ProjectVersion{
			{
				ProjectName: "kubernetes",
				ProductName: "kubernetes1.14.4",
				Version:     "1.14.4",
			},
			{
				ProjectName: "kubernetes",
				ProductName: "kubernetes1.14.5",
				Version:     "1.14.5",
			},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProjectVersion{
				ProjectName: tt.fields.ProjectName,
				ProductName: tt.fields.ProductName,
				Version:     tt.fields.Version,
			}
			got, err := p.ListFromProjectName(tt.args.projectName, tt.args.class)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListFromProjectName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListFromProjectName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
