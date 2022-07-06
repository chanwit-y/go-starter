package client

import (
	"fmt"
	"go-stater/module/master/client/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	repo := mock.NewMockRepository(t)
	service := NewService(repo)

	result := service.FindAll()
	fmt.Println(result)

	assert.Equal(t, true, result.Status)
}
