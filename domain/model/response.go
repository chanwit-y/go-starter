package model

type Response[T any] struct {
	Data   T
	Status bool `json:"status"`
}
