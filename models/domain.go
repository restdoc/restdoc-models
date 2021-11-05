package Models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Domain struct {
	ID                  uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	UserId              uint64 `gorm:"type:bigint unsigned;not null;" json:"user_id"`
	Domain              string `gorm:"type:varchar(128);uniqueIndex:idx_domain;not null" json:"domain"`
	RootDomain          string `gorm:"type:varchar(128);uniqueIndex:idx_root_domain;not null" json:"root_domain"`
	ApiKey              string `gorm:"type:varchar(64);not null;" json:"api_key"`               //todo
	DefaultSmtpPassword string `gorm:"type:varchar(64);not null;" json:"default_smtp_password"` //todo
	CreateAt            int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt            int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
	Type                uint8  `gorm:"type:tinyint(8);not null;" json:"type"`
	Valid               bool   `gorm:"type:boolean;not null;default false;" json:"valid"`
	Amount              uint32 `gorm:"type:int unsigned;not null;" json:"amount"`
	InboundSpamFilter   bool   `gorm:"type:boolean;not null;default false;" json:"inbound_spam_filter"`
	WildcardDomain      bool   `gorm:"type:boolean;not null;default false;" json:"wildcard_domain"`
	Unsubscribes        bool   `gorm:"type:boolean;not null;default true;" json:"unsubscribes"`
	OpenTracking        bool   `gorm:"type:boolean;not null;default true;" json:"open_tracking"`
	TrackingHostname    string `gorm:"type:varchar(128);not null;" json:"tracking_hostname"`

	ClickTracking      bool   `gorm:"type:boolean;not null;default true;" json:"click_tracking"`
	DkimPubKey         string `gorm:"type:text;not null;" json:"dkim_pub_key"`
	DkimPrivateKey     string `gorm:"type:text;not null;" json:"dkim_private_key"`
	DkimRecord         string `gorm:"type:text;not null;" json:"dkim_record"`
	CurrentDkimRecord  string `gorm:"type:text;not null;" json:"current_dkim_record"`
	SpfRecord          string `gorm:"type:text;not null;" json:"spf_record"`
	CurrentSpfRecord   string `gorm:"type:text;not null;" json:"current_spf_record"`
	MxaRecord          string `gorm:"type:text;not null;" json:"mxa_record"`
	MxbRecord          string `gorm:"type:text;not null;" json:"mxb_record"`
	CurrentMxaRecord   string `gorm:"type:text;not null;" json:"current_mxa_record"`
	CurrentMxbRecord   string `gorm:"type:text;not null;" json:"current_mxb_record"`
	DmarcRecord        string `gorm:"type:text;not null;" json:"dmarc_record"`
	CurrentDmarcRecord string `gorm:"type:text;not null;" json:"current_dmarc_record"`
	SpfValid           bool   `gorm:"type:boolean;not null;default false;" json:"spf_valid"`
	Hostname           string `gorm:"type:varchar(128);not null;" json:"hostname"`
	SmtpServer         string `gorm:"type:varchar(128);not null;" json:"smtp_server"`
	ForeignSmtpServer  string `gorm:"type:varchar(128);not null;" json:"foreign_smtp_server"`
	InboxRules         string `gorm:"type:text;not null;default '';" json:"inbox_rules"`
}

func (d *Domain) TableName() string {
	return "domain"
}

func GetAllDomain(d *[]Domain) (err error) {
	if err = DB.Find(d).Error; err != nil {
		return err
	}
	return nil
}

func GetAllDomainName(d *[]Domain, created int64) (err error) {
	if err = DB.Where("created_at > ?", created).Find(d).Select([]string{"domain", "created_at"}).Error; err != nil {
		return err
	}
	return nil
}

func AddNewDomain(d *Domain) (err error) {
	if err = DB.Create(d).Error; err != nil {
		return err
	}
	return nil
}

func GetOneDomain(d *Domain, id string, user_id string) (err error) {
	if err := DB.Where("id = ? AND user_id = ?", id, user_id).First(d).Error; err != nil {
		return err
	}
	return nil
}

func GetDomainsByRootDomain(domains *[]Domain, rootDomain string) (err error) {
	if err := DB.Where("root_domain = ?", rootDomain).Select("user_id,root_domain").Find(domains).Error; err != nil {
		return err
	}
	return nil
}

func GetDomainByName(d *Domain, domain string) (err error) {
	if err := DB.Where("domain = ?", domain).First(d).Error; err != nil {
		return err
	}
	return nil
}

func GetDomainByNameAndUser(d *Domain, domain string, user_id uint64) (err error) {
	if err := DB.Where("user_id = ? AND domain = ?", user_id, domain).First(d).Error; err != nil {
		return err
	}
	return nil
}

func GetDomainsByUserId(domains *[]Domain, user_id uint64, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ?", user_id).Find(&domains).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func PutOneDomain(d *Domain) (err error) {
	DB.Save(d)
	return nil
}

func UpdateCheckResult(d *Domain) error {
	ts := time.Now().Unix()

	updates := map[string]interface{}{
		"valid":                d.Valid,
		"current_spf_record":   d.CurrentSpfRecord,
		"current_dkim_record":  d.CurrentDkimRecord,
		"current_dmarc_record": d.CurrentDmarcRecord,
		"current_mxa_record":   d.CurrentMxaRecord,
		"current_mxb_record":   d.CurrentMxbRecord,
		"spf_valid":            d.SpfValid,
		"update_at":            ts,
	}
	if err := DB.Model(d).Where("user_id = ? AND domain = ?", d.UserId, d.Domain).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDomainInboxRules(d *Domain) error {
	ts := time.Now().Unix()

	updates := map[string]interface{}{
		"inbox_rules": d.InboxRules,
		"update_at":   ts,
	}

	if err := DB.Model(d).Where("user_id = ? AND domain = ?", d.UserId, d.Domain).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDomain(d *Domain, domain string, user_id uint64) (err error) {
	//if err := DB.Unscoped().Where("user_id = ? AND domain = ?", user_id, domain).Delete(d).Error; err != nil {
	if err := DB.Where("user_id = ? AND domain = ?", user_id, domain).Delete(d).Error; err != nil {
		return err
	}
	return nil
}
