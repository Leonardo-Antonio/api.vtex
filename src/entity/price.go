package entity

type Price struct {
	SkuId     string  `json:"sku_id,omitempty"`
	CostPrice float64 `json:"costPrice,omitempty"`
	BasePrice float64 `json:"basePrice,omitempty"`
	ListPrice float64 `json:"listPrice,omitempty"`
}
