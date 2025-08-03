package model

type BrandDiscount struct {
	Id       int64   `xorm:"pk autoincr 'id'"`
	Brand    string  `xorm:"varchar(255) not null 'brand'"`
	Discount float64 `xorm:"float not null 'discount'"`
}

func (u *BrandDiscount) TableName() string {
	return "brand_discount"
}
