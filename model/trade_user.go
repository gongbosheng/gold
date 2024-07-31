package model

type TradeUser struct {
	UserID    int    `gorm:"primary_key;auto_increment" json:"user_id"`
	Username  string `gorm:"username" json:"username"`
	Email     string `gorm:"email" json:"email"`
	Phone     string `gorm:"phone" json:"phone"`
	UserRole  string `gorm:"user_role" json:"user_role"`
	UserGroup string `gorm:"user_group" json:"user_group"`
	Status    bool   `gorm:"status" json:"status"`
}

func (TradeUser) TableName() string {
	return "trade_users"
}
