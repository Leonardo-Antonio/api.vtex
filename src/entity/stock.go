package entity

type Stock struct {
	SkuId             string `json:"sku_id,omitempty"`
	WarehousesId      string `json:"warehouses,omitempty"`
	Quantity          int    `json:"quantity,omitempty"`
	UnlimitedQuantity bool   `json:"unlimitedQuantity,omitempty"`
}
