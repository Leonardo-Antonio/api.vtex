package entity

type Image struct {
	Url    string `json:"url,omitempty"`
	SkuEan string `json:"sku_ean,omitempty"` // body
	Name   string `json:"name,omitempty"`    // body
	Text   string `json:"text,omitempty"`    // body
	Label  string `json:"label,omitempty"`
	SkuId  string `json:"sku_id,omitempty"` // body
}
