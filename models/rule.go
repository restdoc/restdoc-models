package Models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Rule struct {
	ID              uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	UserId          uint64 `gorm:"type:bigint unsigned;not null;" json:"user_id"`
	Type            uint8  `gorm:"type:tinyint(8);not null;" json:"type"`
	Valid           bool   `gorm:"type:boolean;not null;default false;" json:"valid"`
	From            string `gorm:"type:varchar(256);not null;" json:"from"`
	To              string `gorm:"type:varchar(256);not null;" json:"to"`
	LargerSize      uint64 `gorm:"type:bigint unsigned;not null;default 0" json:"larger_size"`
	SmallerSize     uint64 `gorm:"type:bigint unsigned;not null;default 0" json:"smaller_size"`
	SubjectContains string `gorm:"type:varchar(256);not null;" json:"subject_contains"`
	ContentContains string `gorm:"type:varchar(256);not null;" json:"content_contains"`
	NotContains     string `gorm:"type:varchar(256);not null;" json:"not_contains"`
	Archive         bool   `gorm:"type:boolean;not null;default false;" json:"archive"`
	MarkRead        bool   `gorm:"type:boolean;not null;default false;" json:"mark_read"`
	ImportantState  uint8  `gorm:"type:tinyint(8);not null;" json:"important_state"`
	Star            bool   `gorm:"type:boolean;not null;default false;" json:"star"`
	Delete          bool   `gorm:"type:boolean;not null;default false;" json:"delete"`
	NeverSpam       bool   `gorm:"type:boolean;not null;default false;" json:"never_spam"`
	LabelId         uint64 `gorm:"type:bigint unsigned;not null;default 0;" json:"label_id"`
	ForwardTo       string `gorm:"type:varchar(256);not null;" json:"forward_to"`
	CreateAt        int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt        int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (r *Rule) TableName() string {
	return "rule"
}

func GetAllRule(r *[]Rule) (err error) {
	if err = DB.Find(r).Error; err != nil {
		return err
	}
	return nil
}

func AddNewRule(r *Rule) (err error) {
	if err = DB.Create(r).Error; err != nil {
		return err
	}
	return nil
}

func GetOneRule(r *Rule, id string, user_id string) (err error) {
	if err := DB.Where("id = ? AND user_id = ?", id, user_id).First(r).Error; err != nil {
		return err
	}
	return nil
}

func GetRulesByUserId(rules *[]Rule, user_id uint64) (err error) {
	if err := DB.Where("user_id = ?", user_id).Find(&rules).Error; err != nil {
		return err
	}
	return nil
}

func PutOneRule(r *Rule) (err error) {
	DB.Save(r)
	return nil
}

func UpdateRule(r *Rule, updates map[string]interface{}) error {
	if err := DB.Model(r).Where("id = ? AND user_id = ?", r.ID, r.UserId).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRule(r *Rule) (err error) {
	//if err := DB.Unscoped().Where("user_id = ? AND domain = ?", user_id, domain).Delete(r).Error; err != nil {
	if err := DB.Where("id = ? AND user_id = ?", r.ID, r.UserId).Delete(r).Error; err != nil {
		return err
	}

	return nil
}
