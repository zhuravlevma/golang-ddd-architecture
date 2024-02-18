package dtos

type UpdateOrderStatusDto struct {
	Delivered *bool `json:"Delivered"`
	Returned *bool `json:"Returned"`
}
