package entity

type Brand struct {
	Id                     string `json:"Id,omitempty" validmor:"required"`
	Name                   string `json:"Name,omitempty" validmor:"required"`
	Keywords               string `json:"Keywords,omitempty" validmor:"required"`
	Text                   string `json:"Text,omitempty" validmor:"required"`
	SiteTitle              string `json:"SiteTitle,omitempty" validmor:"required"`
	AdWordsRemarketingCode string `json:"AdWordsRemarketingCode,omitempty"`
	LomadeeCampaignCode    string `json:"LomadeeCampaignCode,omitempty"`
	Score                  int32  `json:"Score,omitempty" validmor:"required"`
	MenuHome               bool   `json:"MenuHome,omitempty" validmor:"required"`
	Active                 bool   `json:"Active,omitempty" validmor:"required"`
	LinkId                 string `json:"LinkId,omitempty"`
}
