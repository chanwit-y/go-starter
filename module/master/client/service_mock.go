package client

import (
	"go-stater/domain/schema"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func NewRepositoryMock() *repositoryMock {
	return &repositoryMock{}
}

func (c *repositoryMock) FindAll() []schema.Client {
	mockData := []schema.Client{
		{
			Code: "1",
			Name: "Unit Test",
		},
	}
	return mockData
}

func (c *repositoryMock) Create(client schema.Client) bool {
	return true
}
