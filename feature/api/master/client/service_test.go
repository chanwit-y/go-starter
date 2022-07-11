//go:generate mockery --outpkg mock --all --output ./mock
package client

import (
	"fmt"
	"go-stater/domain/model"
	"go-stater/feature/api/master/client/mock"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		repository Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_FindAll(t *testing.T) {
	type fields struct {
		repository Repository
	}
	tests := []struct {
		name   string
		fields fields
		want   model.Response[[]model.Client]
	}{
		// TODO: Add test cases.
		{
			name:   "a",
			fields: fields{repository: mock.NewMockRepository(t)},
			want: model.Response[[]model.Client]{
				Data: []model.Client{{
					Code: "1",
					Name: "Unit Test",
				}},
				Status: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
			}
			fmt.Println(s.FindAll())
			if got := s.FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
