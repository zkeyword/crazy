package shop

import "time"

type ShopProduct struct {
	ID          uint      `gorm:"primary_key"`
	Name        string    `gorm:"size:255"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"type:decimal(10,2)"`
	Stock       uint      `gorm:"type:uint"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type ShopCategory struct {
	ID          uint      `gorm:"primary_key"`
	Name        string    `gorm:"size:255"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type ShopCart struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `gorm:"type:uint"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type ShopCartItem struct {
	ID        uint `gorm:"primary_key"`
	CartID    uint `gorm:"type:uint"`
	ProductID uint `gorm:"type:uint"`
	Quantity  uint `gorm:"type:uint"`
}

type ShopOrder struct {
	ID         uint      `gorm:"primary_key"`
	UserID     uint      `gorm:"type:uint"`
	CouponID   uint      `gorm:"type:uint"`
	TotalPrice float64   `gorm:"type:decimal(10,2)"`
	Status     string    `gorm:"size:255"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type ShopOrderItem struct {
	ID        uint    `gorm:"primary_key"`
	OrderID   uint    `gorm:"type:uint"`
	ProductID uint    `gorm:"type:uint"`
	Quantity  uint    `gorm:"type:uint"`
	Price     float64 `gorm:"type:decimal(10,2)"`
}

type ShopAddress struct {
	ID         uint      `gorm:"primary_key"`
	UserID     uint      `gorm:"type:uint"`
	Address    string    `gorm:"type:text"`
	PostalCode string    `gorm:"size:20"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type ShopCoupon struct {
	ID        uint      `gorm:"primary_key"`
	Code      string    `gorm:"size:255"`
	Discount  float64   `gorm:"type:decimal(10,2)"`
	ValidFrom time.Time `gorm:"type:datetime"`
	ValidTo   time.Time `gorm:"type:datetime"`
	Used      bool      `gorm:"type:boolean"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type ShopUserCoupon struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `gorm:"type:uint"`
	CouponID  uint      `gorm:"type:uint"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
