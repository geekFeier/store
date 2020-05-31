package module

import (
	"reflect"
	"testing"
)

func TestProjectClass_List(t *testing.T) {
	type fields struct {
		Name     string
		Alias    string
		Label    string
		Describe string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []ProjectClass
		wantErr bool
	}{
		{name: "all", fields: fields{
			Name:     "cloud_kernel",
			Alias:    "云内核",
			Label:    "key=kernel",
			Describe: "xxx",
		}, want: []ProjectClass{
			{
				Name:     "cloud_kernel",
				Alias:    "云内核",
				Label:    "key=kernel",
				Describe: "xxx",
			},
			{
				Name:     "stateful_set",
				Alias:    "有状态应用",
				Label:    "key=app",
				Describe: "xxx",
			},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProjectClass{
				Name:     tt.fields.Name,
				Alias:    tt.fields.Alias,
				Label:    tt.fields.Label,
				Describe: tt.fields.Describe,
			}
			got, err := p.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}
