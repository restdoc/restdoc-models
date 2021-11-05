package Models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type MailLog struct {
	ID         string `gorm:"type:varchar(100);not null;primary_key;" json:"id"`
	UserId     uint64 `gorm:"type:bigint unsigned;not null;" json:"user_id"`
	Domain     string `gorm:"type:varchar(128);not null;" json:"domain"`
	Type       uint8  `gorm:"type:tinyint(8);not null;" json:"type"`
	From       string `gorm:"type:varchar(128);not null;" json:"from"`
	To         string `gorm:"type:varchar(128);not null;" json:"to"`
	Subject    string `gorm:"type:varchar(255);not null;" json:"subject"`
	Created    uint64 `gorm:"type:bigint unsigned;not null;" json:"created"`
	Updated    uint64 `gorm:"type:bigint unsigned;not null;" json:"updated"`
	Status     uint8  `gorm:"type:smallint unsigned;not null;" json:"status"`
	Code       uint16 `gorm:"type:smallint unsigned;not null;" json:"code"`
	Msg        string `gorm:"type:varchar(255);not null;" json:"msg"`
	SmtpServer string `gorm:"type:varchar(128);not null;" json:"smtp_server"`
}

func (m *MailLog) TableName() string {
	return "maillog"
}

func AddNewEmail(m *MailLog) (err error) {
	if err = DB.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func GetOneEmail(m *MailLog, id string) (err error) {
	if err := DB.Where("id = ?", id).First(m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

func GetEmailByDomainName(m *MailLog, domain string) (err error) {
	if err := DB.Where("domain = ?", domain).First(m).Error; err != nil {
		return err
	}
	return nil
}

func GetEmailsByUserId(emails *[]MailLog, user_id uint64, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ?", user_id).Find(&emails).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetEmailsByDomain(emails *[]MailLog, domain string, limit int, offset int) (err error) {
	if err := DB.Where("domain = ?", domain).Find(&emails).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func PutOneEmail(m *MailLog) (err error) {
	DB.Save(m)
	return nil
}

/*
func UpdateCheckResult(m *MailLog) error {
	if err := DB.Model(m).Where("user_id = ? AND domain = ?", m.UserId, m.MailLog).Updates(MailLog{Valid: m.Valid, CurrentSpfRecord: m.CurrentSpfRecord, CurrentDkimRecord: m.CurrentDkimRecord}).Error; err != nil {
		return err
	}
	return nil
}
*/

func UpdateStatus(m *MailLog) error {
	if err := DB.Model(m).Where("id = ?", m.ID).Updates(MailLog{Status: m.Status, Code: m.Code, Msg: m.Msg, Updated: m.Updated}).Error; err != nil {
		return err
	}
	return nil
}

func SummaryEmails(emails *[]MailLog, begin int64, end int64) (err error) {
	if err := DB.Where("created > ? and created < ?", begin, end).Select("user_id, status, domain").Find(&emails).Error; err != nil {
		return err
	}
	return nil
}
