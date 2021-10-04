package entity

type Specification struct {
	SkuId              string `json:"sku_id,omitempty"`
	Name               string `json:"name,omitempty"`
	FieldId            string `json:"field_id,omitempty"`
	FieldName          string `json:"field_name,omitempty"`
	FieldType          string `json:"field_type,omitempty"`
	FieldValueId       string `json:"field_value_id,omitempty"`
	FieldValueName     string `json:"field_value_name,omitempty"`
	SpecificationCode  string `json:"specification_code,omitempty"`
	SpecificationValue string `json:"specification_value,omitempty"`
	SkuReferenceCode   string `json:"sku_reference_code,omitempty"`
}

type ProductsSpecification struct {
	Search            []string         `json:"sku_ids,omitempty"`
	DataSpecification []*Specification `json:"data,omitempty"`
}
