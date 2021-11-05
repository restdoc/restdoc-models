package Models

import (
	//	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Unsub struct {
	ID             uint64 `gorm:"type:bigint unsigned auto_increment; not null; primary_key;" json: "id"`
	EmailID        uint64 `gorm:"type:bigint unsigned;not null" json:"email_id"`
	Hash           uint64 `gorm:"type:bigint unsigned; not null; index;" json:"hash"` //siphash(from+to)
	DomainId       uint64 `gorm:"type:bigint unsigned; not null; index;" json:"domain_id"`
	UserId         uint64 `gorm:"type:bigint unsigned; not null; index;" json:"user_id"`
	From           string `gorm:"type:varchar(320);not null;unique_index:from_to" json:"from"`
	To             string `gorm:"type:varchar(320);not null;unique_index:from_to" json:"to"`
	FromUser       string `gorm:"type:varchar(64);not null" json:"from_user"`
	FromDomain     string `gorm:"type:varchar(255);not null" json:"from_domain"`
	FromRootDomain string `gorm:"type:varchar(255);not null" json:"from_root_domain"`
	ToUser         string `gorm:"type:varchar(64);not null" json:"to_user"`
	ToDomain       string `gorm:"type:varchar(255);not null" json:"to_domain"`
	ToRootDomain   string `gorm:"type:varchar(255);not null" json:"to_root_domain"`
	CreateAt       int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt       int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
	Type           uint8  `gorm:"type:tinyint(8);not null;" json:"type"`
}

func (un *Unsub) TableName() string {
	return "unsub"
}

func AddNewUnsub(un *Unsub) (err error) {
	if err = DB.Create(un).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUnsub(un *Unsub, id string) (err error) {
	if err := DB.Where("id = ?", id).First(un).Error; err != nil {
		return err
	}
	return nil
}

func GetUnsubsByRootDomain(unsubs *[]Unsub, rootDomain string) (err error) {
	if err := DB.Where("from_root_domain = ?", rootDomain).Find(unsubs).Error; err != nil {
		return err
	}
	return nil
}

func GetUnsubByHash(unsubs *[]Unsub, hash uint64) (err error) {
	if err := DB.Where("hash = ?", hash).Find(unsubs).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUnsub(un *Unsub) (err error) {
	DB.Save(un)
	return nil
}

func DeleteUnsub(un *Unsub, id uint64) (err error) {
	if err := DB.Where("id = ?", id).Delete(un).Error; err != nil {
		return err
	}
	return nil
}
