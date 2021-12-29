package entity

type ImageSku struct {
	SkuId string
	Url   string
	Name  string
	Text  string
	Label string
}

type ImagesSkus []*ImageSku
