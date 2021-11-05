package Models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DomainUser struct {
	ID          uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	DomainID    uint64 `gorm:"type:bigint unsigned; not null; " json: "domain_id"`
	Domain      string `gorm:"type:varchar(64);not null; uniqueIndex:idx_user_domain;" json:"domain"`
	UserID      uint64 `gorm:"type:bigint unsigned; not null; " json: "user_id"`
	User        string `gorm:"type:varchar(64);not null; uniqueIndex:idx_user_domain;" json:"user"`
	Name        string `gorm:"type:varchar(64);not null" json:"name"`
	Password    string `gorm:"type:varchar(64);not null;" json:"password"`
	Locale      string `gorm:"type:varchar(64);not null;" json:"locale"`
	UniqId      uint64 `gorm:"type:bigint unsigned;not null;" json:"uniq_id"`
	VerifyCode  string `gorm:"type:varchar(128);not null;" json:"verify_code"`
	BackupEmail string `gorm:"type:varchar(64);not null;default ''" json:"backup_email"`
	Valid       bool   `gorm:"type:boolean;not null;default false;" json:"valid"`
	Type        uint8  `gorm:"type:tinyint unsigned;not null;default 0" json:"type"`
	CreateAt    int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt    int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (u *DomainUser) TableName() string {
	return "domain_user"
}

func MembersCount() int64 {
	var count int64
	DB.Model(&DomainUser{}).Where("id > ?", 0).Count(&count)
	return count
}

func GetAllDomainUser(u *[]DomainUser) (err error) {
	if err = DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewDomainUser(u *DomainUser) (err error) {
	if err = DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneDomainUser(u *DomainUser, id uint64) (err error) {
	if err := DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetDomainUserByDomainAndUser(u *DomainUser, domain string, user string) (err error) {
	if err := DB.Where("domain = ? AND user = ?", domain, user).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetDomainUserByDomainIdAndUser(u *DomainUser, domainId uint64, user string) (err error) {
	if err := DB.Where("domain_id = ? AND user = ?", domainId, user).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetDomainUsersByDomain(u *[]DomainUser, domain string) (err error) {
	if err := DB.Where("domain = ?", domain).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func VerifyDomainUserByCode(u *DomainUser, domainId uint64, user string, code string) (err error) {
	if err := DB.Where("domain_id = ? AND user = ? AND verify_code = ?", domainId, user, code).First(u).Error; err != nil {
		return err
	}
	if u.Valid == true {
		return nil
	}

	ts := time.Now().Unix()
	if err := DB.Model(u).Where("id = ?", u.ID).Updates(map[string]interface{}{"update_at": ts, "valid": true, "verify_code": ""}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDomainUserLocale(u *DomainUser, locale string) (err error) {
	ts := time.Now().Unix()
	if err := DB.Model(u).Where("id = ?", u.ID).Updates(map[string]interface{}{"locale": locale, "update_at": ts}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDomainUserPassword(u *DomainUser, domain_id uint64, password string) (err error) {
	ts := time.Now().Unix()
	if err := DB.Model(u).Where("id = ? AND domain_id = ?", u.ID, u.DomainID).Updates(map[string]interface{}{"password": password, "update_at": ts}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDomainUserVerifyCode(u *DomainUser, code string) (err error) {
	if err := DB.Model(u).Where("id = ?", u.ID).Updates(map[string]interface{}{"verify_code": code}).Error; err != nil {
		return err
	}
	return nil
}

func PutOneDomainUser(u *DomainUser) (err error) {
	DB.Save(u)
	return nil
}

func DeleteDomainUser(u *DomainUser, id uint64, user_id uint64) (err error) {
	DB.Where("id = ? AND user_id = ?", id, user_id).Delete(u)
	return nil
}
