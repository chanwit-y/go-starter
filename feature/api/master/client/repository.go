package client

import (
	"go-stater/domain/schema"
	"go-stater/pkg/database"

	"gorm.io/gorm"
)

type (
	Repository interface {
		FindAll() []schema.Client
		Create(client schema.Client) bool
		// Update(client schema.Client) schema.Client
		// Delete(id int64) bool
	}
	repository struct{ db *gorm.DB }
)

func NewRepository() Repository {
	return &repository{db: database.GetDB()}
}

func (r *repository) FindAll() []schema.Client {
	var clients []schema.Client
	r.db.Find(&clients)
	return clients
}

func (r *repository) Create(client schema.Client) bool {
	r.db.Create(&client)
	return true
}
