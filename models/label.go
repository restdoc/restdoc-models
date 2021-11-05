package Models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Label struct {
	ID               uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	UserId           uint64 `gorm:"type:bigint unsigned;uniqueIndex:idx_user_id_name;not null;" json:"user_id"`
	Name             string `gorm:"type:varchar(128);uniqueIndex:idx_user_id_name;not null;" json:"name"`
	Count            uint32 `gorm:"type:int unsigned;not null; default 0" json:"count"`
	UnreadCount      uint32 `gorm:"type:int unsigned;not null; default 0" json:"unread_count"`
	Type             uint8  `gorm:"type:tinyint unsigned;not null;default 0" json:"type"`
	Color            string `gorm:"type:char(7);not null;default ''" json:"color"`
	MailListDisplay  uint8  `gorm:"type:tinyint unsigned;not null; default 1" json:"mail_list_display"`
	ImapDisplay      uint8  `gorm:"type:tinyint unsigned;not null; default 1" json:"imap_display"`
	LabelListDisplay uint8  `gorm:"type:tinyint unsigned;not null; default 1" json:"label_list_display"`
	CreateAt         int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt         int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (l *Label) TableName() string {
	return "label"
}

func AddNewLabel(l *Label) (err error) {
	if err = DB.Create(l).Error; err != nil {
		return err
	}
	return nil
}

func GetOneLabel(l *Label, id uint64, user_id uint64) (err error) {
	if err := DB.Where("id = ? AND user_id = ?", id, user_id).First(l).Error; err != nil {
		return err
	}
	return nil
}

func GetLabelNamed(l *Label, name string, user_id uint64) (err error) {
	if err := DB.Where("user_id = ? AND name = ?", user_id, name).First(l).Error; err != nil {
		return err
	}
	return nil
}

func GetAllLabelsByUserId(labels *[]Label, user_id uint64) (err error) {
	if err := DB.Where("user_id = ?", user_id).Order("id").Find(labels).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetLabelsByUserId(labels *[]Label, user_id uint64, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ?", user_id).Order("id").Find(labels).Limit(limit).Offset(offset).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateLabelStates(label *Label, userId uint64, changes map[string]interface{}) error {
	if err := DB.Model(label).Where("id= ? AND user_id = ?", label.ID, userId).Updates(changes).Error; err != nil {
		return err
	}
	return nil
}

func PutOneLabel(l *Label) (err error) {
	DB.Save(l)
	return nil
}

/*
func UpdateStatus(l *Label) error {
	ts := time.Now().Unix()
	if err := DB.Model(l).Where("id = ? AND user_id = ?", m.ID, m.UserId).Updates(Label{Valid: d.Valid, CurrentSpfRecord: d.CurrentSpfRecord, CurrentDkimRecord: d.CurrentDkimRecord, CurrentDmarcRecord: d.CurrentDmarcRecord, CurrentMxaRecord: d.CurrentMxaRecord, CurrentMxbRecord: d.CurrentMxbRecord, UpdateAt: ts}).Error; err != nil {
		return err
	}
	return nil
}
*/

func DeleteLabel(l *Label, id uint64, user_id uint64) (err error) {
	if err := DB.Unscoped().Where("id = ? AND user_id = ?", id, user_id).Delete(l).Error; err != nil {
		return err
	}
	return nil
}
