package model

// ShopCart 购物车表
type ShopCart struct {
	ItemId        string `json:"itemId"`
	ItemImgUrl    string `json:"itemImgUrl"`
	ItemName      string `json:"itemName"`
	SpecId        string `json:"specId"`
	SpecName      string `json:"specName"`
	PriceDiscount string `json:"priceDiscount"`
	PriceNormal   string `json:"priceNormal"`
	Status        int    `json:"status" gorm:"not null;comment:'状态: 1 - 正常、0 - 禁用'"`
}
