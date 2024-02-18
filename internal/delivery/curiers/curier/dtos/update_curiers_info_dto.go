package dtos

type UpdateCuriersInfoDto struct {
	FirstName *string `json:"FirstName"`
	LastName *string `json:"LastName"`
	IsActive *bool `json:"IsActive"`
}
