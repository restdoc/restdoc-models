package Models

import (
	_ "github.com/go-sql-driver/mysql"
)

type AppUser struct {
	ID           uint64 `gorm:"type:bigint unsigned; not null;primary_key; " json: "id"`
	DomainID     uint64 `gorm:"type:bigint unsigned; not null;" json: "domain_id"`
	DomainUserID uint64 `gorm:"type:bigint unsigned; not null;" json: "domain_user_id"`
	AppName      string `gorm:"type:varchar(64);not null" json:"app_name"`
	Password     string `gorm:"type:varchar(64);not null;" json:"password"`
	AllowIPs     string `gorm:"type:varchar(128);not null;default ''" json:"allow_ips"`
	UserType     uint8  `gorm:"type:tinyint unsigned;not null;default 0;" json:"user_type"`
	CreateAt     int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt     int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (u *AppUser) TableName() string {
	return "app_user"
}

func GetAllAppUser(u *[]AppUser) (err error) {
	if err = DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewAppUser(a *AppUser) (err error) {
	if err = DB.Create(a).Error; err != nil {
		return err
	}
	return nil
}

func GetOneAppUser(a *AppUser, id uint64) (err error) {
	if err := DB.Where("id = ?", id).First(a).Error; err != nil {
		return err
	}
	return nil
}

func GetOneAppUserByPassword(a *AppUser, domain_user_id uint64, domain_id uint64, password string) (err error) {
	if err := DB.Where("domain_user_id = ? AND domain_id = ? AND password = ?", domain_user_id, domain_id, password).First(a).Error; err != nil {
		return err
	}
	return nil
}

func GetAppUsersByDomainUserId(users *[]AppUser, domain_user_id uint64) (err error) {
	if err := DB.Where("domain_user_id = ?", domain_user_id).Find(users).Error; err != nil {
		return err
	}
	return nil
}

func PutOneAppUser(a *AppUser) (err error) {
	DB.Save(a)
	return nil
}

func DeleteAppUser(a *AppUser, id uint64, user_id uint64) (err error) {
	DB.Where("id = ? AND domain_user_id = ? ", id, user_id).Delete(a)
	return nil
}
