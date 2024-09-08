package dto

type ExampleCreate struct {
	Name   string `json:"name" valid:"required"`
	Detail string `json:"detail"`
}

type ExamplePut struct {
	Name   string `json:"name" valid:"required"`
	Detail string `json:"detail"`
}