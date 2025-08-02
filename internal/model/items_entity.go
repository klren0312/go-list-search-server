package model

type Items struct {
	Id           int64   `xorm:"pk autoincr 'id'"`
	Name         string  `xorm:"varchar(255) not null 'name'"`
	ItemNo       string  `xorm:"varchar(255) not null 'item_no'"`
	Standard     string  `xorm:"varchar(255) not null 'standard'"`
	Brand        string  `xorm:"varchar(255) not null 'brand'"`
	DeliveryTime string  `xorm:"varchar(255) not null 'delivery_time'"`
	Price        float64 `xorm:"float not null 'price'"`
	Discount     float64 `xorm:"float not null 'discount'"`
}

func (u *Items) TableName() string {
	return "items"
}
