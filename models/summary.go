package Models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Summary struct {
	ID         uint64 `gorm:"type:bigint unsigned auto_increment; not null; primary_key;" json: "id"`
	UserId     uint64 `gorm:"type:bigint unsigned;not null;uniqueIndex:user_id_date_domain_id" json:"user_id"`
	Date       string `gorm:"type:char(8); not null;uniqueIndex:user_id_date_domain_id" json:"date"`
	DomainId   uint64 `gorm:"type:bigint unsigned;not null;uniqueIndex:user_id_date_domain_id" json:"domain_id"`
	Processed  uint64 `gorm:"type:bigint unsigned;not null;" json:"processed"`
	Delivered  uint64 `gorm:"type:bigint unsigned;not null;" json:"delivered"`
	Suppressed uint64 `gorm:"type:bigint unsigned;not null;" json:"suppressed"` //软退信
	Dropped    uint64 `gorm:"type:bigint unsigned;not null;" json:"dropped"`    //无效邮件
	Spam       uint64 `gorm:"type:bigint unsigned;not null;" json:"spam"`       //垃圾邮件举报
}

func (s *Summary) TableName() string {
	return "summary"
}

func AddNewSummary(s *Summary) (err error) {
	if err = DB.Create(s).Error; err != nil {
		return err
	}
	return nil
}

func GetOneSummary(s *Summary, id string) (err error) {
	if err := DB.Where("id = ?", id).First(s).Error; err != nil {
		return err
	}
	return nil
}

func GetSummaryByDomainName(s *Summary, domain string) (err error) {
	if err := DB.Where("domain = ?", domain).First(s).Error; err != nil {
		return err
	}
	return nil
}

func GetSummariesByUserId(summaries *[]Summary, user_id uint64) (err error) {
	if err := DB.Where("user_id = ?", user_id).Find(&summaries).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetSummariesByDomain(summaries *[]Summary, domain string, limit int, offset int) (err error) {
	if err := DB.Where("domain = ?", domain).Find(&summaries).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func PutOneSummary(s *Summary) (err error) {
	DB.Save(s)
	return nil
}

func SaveSummary(s *Summary) error {

	//db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrCreate(&user)

	rows := DB.Model(s).Where("user_id = ? and date = ? and domain = ?", s.UserId, s.Date, s.Domain).Updates(Summary{Processed: s.Processed, Delivered: s.Delivered, Suppressed: s.Suppressed, Dropped: s.Dropped, Spam: s.Spam}).RowsAffected
	if rows == 0 {
		DB.Save(s)
	}
	return nil
}
