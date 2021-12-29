package entity

type Category struct {
	Id                   string `json:"Id,omitempty"`
	Name                 string `json:"Name,omitempty"`
	Keywords             string `json:"Keywords,omitempty"`
	Title                string `json:"Title,omitempty"`
	Description          string `json:"Description,omitempty"`
	FatherCategoryId     uint   `json:"FatherCategoryId,omitempty"`
	GlobalCategoryId     uint   `json:"GlobalCategoryId,omitempty"`
	ShowInStoreFront     bool   `json:"ShowInStoreFront,omitempty"`
	IsActive             bool   `json:"IsActive,omitempty"`
	ActiveStoreFrontLink bool   `json:"ActiveStoreFrontLink,omitempty"`
	ShowBrandFilter      bool   `json:"ShowBrandFilter,omitempty"`
}
