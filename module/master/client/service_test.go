package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	repo := NewRepositoryMock()
	service := NewService(repo)

	result := service.FindAll()

	assert.Equal(t, true, result.Status)
}
