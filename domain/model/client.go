package model

// type Client struct {
// 	Code string `json:"code"`
// 	Name string `json:"name" validate:"required"`
// 	Num  int    `json:"num" default:"1" validate:"required,number,min=0,max=10"`
// }

type Client struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Num  int    `json:"num"`
}
