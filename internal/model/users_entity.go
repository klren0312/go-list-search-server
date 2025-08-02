package model

/**
 * 用户
 */
type Users struct {
	Id         int64  `xorm:"pk autoincr 'id'"`
	UserId     string `xorm:"varchar(50) not null 'user_id'"`
	Username   string `xorm:"varchar(50) not null 'username'"`
	Password   string `xorm:"varchar(50) not null 'password'"`
	Type       int    `xorm:"int not null 'type'"`       // 0: 超管 1: 经销商 2: 业务员
	ResellerId string `xorm:"varchar(50) 'reseller_id'"` // 业务员的话 会有经销商userid标记
}

func (u *Users) TableName() string {
	return "users"
}
