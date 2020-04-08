package models

import (
	"github.com/satori/go.uuid"
	"grpc-sai-test/store/mysql"
	"time"
)

type TUser struct {
	UserID            string     `gorm:"primary_key;column:UserID"`
	UserName          string     `gorm:"column:UserName"`
	Password          string     `gorm:"column:Password"`
	UserType          int64      `gorm:"column:UserType"`
	Email             string     `gorm:"column:Email"`
	Mobile            string     `gorm:"column:Mobile"`
	MobileCountryCode int64      `gorm:"column:MobileCountryCode"`
	Gender            int64      `gorm:"column:Gender"`
	Avatar            string     `gorm:"column:Avatar"`
	AvatarThumb       string     `gorm:"column:AvatarThumb"`
	AvatarLarge       string     `gorm:"column:AvatarLarge"`
	Grade             int64      `gorm:"column:Grade"`
	SupportNum        int64      `gorm:"column:SupportNum"`
	CountryID         int64      `gorm:"column:CountryID"`
	ProvinceID        int64      `gorm:"column:ProvinceID"`
	CityID            int64      `gorm:"column:CityID"`
	Status            int64      `gorm:"column:Status"`
	CreatedAt         time.Time  `gorm:"column:CreatedAt"`
	UpdatedAt         time.Time  `gorm:"column:UpdatedAt"`
	DeletedAt         *time.Time `gorm:"column:DeletedAt"`
}

func (t TUser) Create() (string, error) {
	uid := uuid.NewV4().String()
	t.UserID = uid

	mysql.Engine.AutoMigrate(&TUser{})
	//if primary key is not existing
	err := mysql.Engine.Create(&t).Error
	return uid, err
}
