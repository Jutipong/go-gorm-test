package prejoin

import "time"

type User struct {
	Id         string    `gorm:"primaryKey; column:beast_id"`
	Name       string    `gorm:"column:Name"`
	Last       string    `gorm:"column:Last"`
	CreateDate time.Time `gorm:"column:CreateDate"`
	CreateBy   string    `gorm:"column:CreateBy"`
	AddressId  string    `gorm:"column:AddressId"`
	Address    []Address `gorm:"foreignKey:AddressId;references:AddressId"`
}

func (User) TableName() string {
	return "User"
}

type Address struct {
	AddressId  string    `gorm:"primary;column:AddressId"`
	Addresses  string    `gorm:"column:Addresses"`
	CreateDate time.Time `gorm:"column:CreateDate"`
	CreateBy   string    `gorm:"column:CreateBy"`
	UpdateDate time.Time `gorm:"column:UpdateDate"`
	UpdateBy   string    `gorm:"column:UpdateBy"`
}

func (Address) TableName() string {
	return "Address"
}
