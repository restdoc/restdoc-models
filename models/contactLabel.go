package Models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ContactLabel struct {
	ID           uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	UserId       uint64 `gorm:"type:bigint unsigned;uniqueIndex:idx_user_id_name;not null;" json:"user_id"`
	Name         string `gorm:"type:varchar(128);uniqueIndex:idx_user_id_name;not null" json:"name"`
	Count        uint32 `gorm:"type:int unsigned;not null; default 0" json:"count"`
	UnreadCount  uint32 `gorm:"type:int unsigned;not null; default 0" json:"unread_count"`
	Type         uint8  `gorm:"type:tinyint(8);not null;" json:"type"`
	Color        string `gorm:"type:char(7);not null" json:"color"`
	InboxDisplay bool   `gorm:"type:boolean;not null; default false" json:"inbox_display"`
	SideDisplay  bool   `gorm:"type:boolean;not null; default false" json:"side_display"`
	CreateAt     int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt     int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (l *ContactLabel) TableName() string {
	return "contact_label"
}

func AddNewContactLabel(l *ContactLabel) (err error) {
	if err = DB.Create(l).Error; err != nil {
		return err
	}
	return nil
}

func GetOneContactLabel(l *ContactLabel, id uint64, user_id uint64) (err error) {
	if err := DB.Where("id = ? AND user_id = ?", id, user_id).First(l).Error; err != nil {
		return err
	}
	return nil
}

func GetContactLabelNamed(l *ContactLabel, name string, user_id uint64) (err error) {
	if err := DB.Where("user_id = ? AND name = ?", user_id, name).First(l).Error; err != nil {
		return err
	}
	return nil
}

func GetAllContactLabelsByUserId(labels *[]ContactLabel, user_id uint64) (err error) {
	if err := DB.Where("user_id = ?", user_id).Order("id").Find(labels).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetContactLabelsByUserId(labels *[]ContactLabel, user_id uint64, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ?", user_id).Order("id").Find(labels).Limit(limit).Offset(offset).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func PutOneContactLabel(l *ContactLabel) (err error) {
	DB.Save(l)
	return nil
}

func DeleteContactLabel(l *ContactLabel, id uint64, user_id uint64) (err error) {
	if err := DB.Unscoped().Where("id = ? AND user_id = ?", id, user_id).Delete(l).Error; err != nil {
		return err
	}
	return nil
}
