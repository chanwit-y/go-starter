//go:generate mockery --outpkg mock --all --output ./mock
package client

import (
	"go-stater/domain/model"
	"go-stater/domain/schema"

	"github.com/samber/lo"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{repository}
}

func (s *Service) FindAll() model.Response[[]model.Client] {
	data := s.repository.FindAll()
	return model.Response[[]model.Client]{
		Data: lo.Map(data,
			func(client schema.Client, _ int) model.Client {
				return model.Client{
					Code: client.Code,
					Name: client.Name,
				}
			}),
		Status: true,
	}
}
