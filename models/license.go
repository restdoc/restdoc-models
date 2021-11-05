package Models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type License struct {
	ID       uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	UserId   uint64 `gorm:"type:bigint unsigned;not null;" json:"user_id"`
	Domain   string `gorm:"type:varchar(128);uniqueIndex:idx_domain;not null" json:"domain"`
	License  string `gorm:"type:text;not null" json:"license"`
	Count    uint32 `gorm:"type:int unsigned;not null;" json:"count"`
	CreateAt int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (l *License) TableName() string {
	return "license"
}

func GetAllLicense(ls *[]License) (err error) {
	if err = DB.Find(ls).Error; err != nil {
		return err
	}
	return nil
}

func GetAllLicenseName(ls *[]License, created int64) (err error) {
	if err = DB.Where("created_at > ?", created).Find(ls).Select([]string{"domain", "created_at"}).Error; err != nil {
		return err
	}
	return nil
}

func AddNewLicense(l *License) (err error) {
	if err = DB.Create(l).Error; err != nil {
		return err
	}
	return nil
}

func GetOneLicense(l *License, id string, user_id string) (err error) {
	if err := DB.Where("id = ? AND user_id = ?", id, user_id).First(l).Error; err != nil {
		return err
	}
	return nil
}

func GetLicenseByName(l *License, domain string) (err error) {
	if err := DB.First(l, "domain = ?", domain).Error; err != nil {
		return err
	}
	return nil
}

func GetLicenseByNameAndUser(l *License, domain string, user_id uint64) (err error) {
	if err := DB.Where("user_id = ? AND domain = ?", user_id, domain).First(l).Error; err != nil {
		return err
	}
	return nil
}

func GetLicensesByUserId(ls *[]License, user_id uint64) (err error) {
	if err := DB.Where("user_id = ?", user_id).Find(&ls).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func PutOneLicense(l *License) (err error) {
	DB.Save(l)
	return nil
}

func DeleteLicense(l *License, domain string, user_id uint64) (err error) {
	//if err := DB.Unscoped().Where("user_id = ? AND domain = ?", user_id, domain).Delete(d).Error; err != nil {
	if err := DB.Where("user_id = ? AND domain = ?", user_id, domain).Delete(l).Error; err != nil {
		return err
	}

	return nil
}
