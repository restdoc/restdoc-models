package Models

import (
	"errors"
	"fmt"
	"sort"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MailAll            uint8 = 0
	MailInbox          uint8 = 1
	MailInboxSend      uint8 = 2 //收发
	MailSnoozed        uint8 = 10
	MailSend           uint8 = 11
	MailDraft          uint8 = 12
	MailDeletedForever uint8 = 200
	MailDeleted        uint8 = 201
	MailSpam           uint8 = 202
	MailStarred        uint8 = 190
	MailImportant      uint8 = 191
	MailReaded         uint8 = 192
)

type Mail struct {
	ID               uint64 `gorm:"type:bigint  unsigned; not null; primary_key; " json: "id"`
	UserId           uint64 `gorm:"type:bigint unsigned;not null;" json:"user_id"`
	MailID           string `gorm:"type:varchar(255);not null" json:"mail_id"`
	Date             string `gorm:"type:varchar(32);not null" json:"date"`
	RecvDate         string `gorm:"type:varchar(32);not null" json:"recv_date"`
	From             string `gorm:"type:varchar(255);not null" json:"from"`
	FromName         string `gorm:"type:varchar(64);not null" json:"from_name"`
	Sender           string `gorm:"type:varchar(255);not null" json:"sender"`
	ReplyTo          string `gorm:"type:varchar(255);not null" json:"reply_to"`
	To               string `gorm:"type:varchar(255);not null" json:"to"`
	ToName           string `gorm:"type:varchar(255);not null" json:"to_name"`
	Cc               string `gorm:"type:varchar(255);not null" json:"cc"`
	Bcc              string `gorm:"type:varchar(255);not null" json:"bcc"`
	InReplyTo        string `gorm:"type:varchar(128);not null" json:"in_reply_to"`
	Subject          string `gorm:"type:varchar(255);not null" json:"subject"`
	Size             uint32 `gorm:"type:int unsigned;not null; default 0" json:"size"`
	IsImportant      bool   `gorm:"type:boolean;not null; default false" json:"is_important"`
	IsStarred        bool   `gorm:"type:boolean;not null; default false" json:"is_starred"`
	IsRead           bool   `gorm:"type:boolean;not null; default false" json:"is_read"`
	HasAttachment    bool   `gorm:"type:boolean;not null; default false" json:"has_attachment"`
	Type             uint8  `gorm:"type:tinyint unsigned;not null;default 0" json:"type"`
	ExternalResource uint64 `gorm:"type:bigint unsigned;not null;" json:"external_resource"`
	ParentId         uint64 `gorm:"type:bigint unsigned;not null;" json:"parent_id"`
	Uid              uint32 `gorm:"type:int unsigned;not null;" json:"uid"`
	Headers          string `gorm:"type:text;" json:"headers"`
	BodyStructure    string `gorm:"type:text;" json:"body_struct"`
	CreateAt         int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt         int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (m *Mail) TableName() string {
	return "mail"
}

func AddNewMail(m *Mail) (err error) {

	if err = DB.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func GetOneMail(m *Mail, id uint64, user_id uint64) (err error) {
	if err := DB.Where("id = ? AND user_id = ?", id, user_id).First(m).Error; err != nil {
		return err
	}
	return nil
}

func GetMailsByIds(mails *[]Mail, ids []uint64) (err error) {
	if err := DB.Where("id in (?)", ids).Order("id desc").Find(mails).Error; err != nil {
		return err
	}
	return nil
}

func GetMailsByUserId(mails *[]Mail, user_id uint64, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ?", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetMailsByUserIdAndParentId(mails *[]Mail, user_id uint64, parent_id uint64, limit int, offset int) (err error) {
	if err := DB.Debug().Where("parent_id = ? AND user_id = ?", parent_id, user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
		return err
	}
	return nil
}

func GetInboxMailsByUserId(mails *[]Mail, user_id uint64, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ? AND `type` = ?", user_id, MailInbox).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetMailsByTypeAndUserId(mails *[]Mail, user_id uint64, _type uint8, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ? AND `type` = ?", user_id, _type).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetMailsByTypesAndUserId(mails *[]Mail, user_id uint64, types []uint8, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ? AND `type` in (?)", user_id, types).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetImportantMailsByUserId(mails *[]Mail, user_id uint64, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ? AND is_important = 1", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetStarredMailsByUserId(mails *[]Mail, user_id uint64, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ? AND is_starred = 1", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetMailsByConditions(mails *[]Mail, user_id uint64, conditions map[string]interface{}) (err error) {

	if err := DB.Where(conditions).Select("id").Find(&mails).Error; err != nil {
		return err
	}

	return nil
}

func PutOneMail(m *Mail) (err error) {
	DB.Save(m)
	return nil
}

func UpdateReadState(mails *[]Mail, user_id uint64, state bool) error {

	ts := time.Now().Unix()
	if len(*mails) == 0 {
		return errors.New("wrong length")
	}

	if len(*mails) == 1 {
		m := (*mails)[0]
		updates := map[string]interface{}{
			"is_read":   state,
			"update_at": ts,
		}

		if err := DB.Model(m).Where("id = ? AND user_id = ?", m.ID, user_id).Updates(updates).Error; err != nil {
			return err
		}
		return nil

	} else {
		updates := map[string]interface{}{
			"is_read":   state,
			"update_at": ts,
		}

		ids := []uint64{}
		for i := range *mails {
			m := (*mails)[i]
			id := m.ID
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

		m := Mail{UserId: user_id}
		if err := DB.Model(m).Where("user_id = ? AND id in (?)", user_id, ids).Updates(updates).Error; err != nil {
			return err
		}
		return nil
	}
}

func UpdateStarredState(mails *[]Mail, user_id uint64, state bool) error {

	ts := time.Now().Unix()
	if len(*mails) == 0 {
		return errors.New("wrong length")
	}

	if len(*mails) == 1 {
		m := (*mails)[0]
		updates := map[string]interface{}{
			"is_starred": state,
			"update_at":  ts,
		}

		if err := DB.Model(m).Where("id = ? AND user_id = ?", m.ID, user_id).Updates(updates).Error; err != nil {
			return err
		}
		return nil

	} else {
		updates := map[string]interface{}{
			"is_starred": state,
			"update_at":  ts,
		}

		ids := []uint64{}
		for i := range *mails {
			m := (*mails)[i]
			id := m.ID
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

		m := Mail{UserId: user_id}
		if err := DB.Model(m).Where("user_id = ? AND id in (?)", user_id, ids).Updates(updates).Error; err != nil {
			return err
		}
		return nil
	}
}

func UpdateImportantState(mails *[]Mail, user_id uint64, state bool) error {

	ts := time.Now().Unix()
	if len(*mails) == 0 {
		return errors.New("wrong length")
	}

	if len(*mails) == 1 {
		m := (*mails)[0]
		updates := map[string]interface{}{
			"is_important": state,
			"update_at":    ts,
		}

		if err := DB.Model(m).Where("id = ? AND user_id = ?", m.ID, user_id).Updates(updates).Error; err != nil {
			return err
		}
		return nil

	} else {
		updates := map[string]interface{}{
			"is_important": state,
			"update_at":    ts,
		}

		ids := []uint64{}
		for i := range *mails {
			m := (*mails)[i]
			id := m.ID
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

		m := Mail{UserId: user_id}
		if err := DB.Model(m).Where("user_id = ? AND id in (?)", user_id, ids).Updates(updates).Error; err != nil {
			return err
		}
		return nil
	}
}

func UpdateMailStates(mails *[]Mail, user_id uint64, updates map[string]interface{}) error {

	if len(*mails) == 0 {
		return errors.New("wrong length")
	}

	if len(*mails) == 1 {
		m := (*mails)[0]
		if err := DB.Model(m).Where("id = ? AND user_id = ?", m.ID, user_id).Updates(updates).Error; err != nil {
			return err
		}
		return nil

	} else {
		ids := []uint64{}
		for i := range *mails {
			m := (*mails)[i]
			id := m.ID
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

		m := Mail{UserId: user_id}
		if err := DB.Model(m).Where("user_id = ? AND id in (?)", user_id, ids).Updates(updates).Error; err != nil {
			return err
		}
		return nil
	}
}

func UpdateFlags(mails *[]Mail, user_id uint64, field string, state bool) error {

	ts := time.Now().Unix()
	if len(*mails) == 0 {
		return errors.New("wrong length")
	}

	if len(*mails) == 1 {
		m := (*mails)[0]
		updates := map[string]interface{}{
			field:       state,
			"update_at": ts,
		}

		if err := DB.Model(m).Where("id = ? AND user_id = ?", m.ID, user_id).Updates(updates).Error; err != nil {
			return err
		}
		return nil

	} else {
		updates := map[string]interface{}{
			field:       state,
			"update_at": ts,
		}

		ids := []uint64{}
		for i := range *mails {
			m := (*mails)[i]
			id := m.ID
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

		m := Mail{UserId: user_id}
		if err := DB.Model(m).Where("user_id = ? AND id in (?)", user_id, ids).Updates(updates).Error; err != nil {
			return err
		}
		return nil
	}
}

func DeleteMail(mails *[]Mail, user_id uint64) (err error) {

	if len(*mails) == 0 {
		return errors.New("wrong length")
	}

	ids := []uint64{}
	for i := range *mails {
		m := (*mails)[i]
		id := m.ID
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	m := Mail{UserId: user_id}

	if err := DB.Unscoped().Where("user_id = ? AND id in (?)", user_id, ids).Delete(m).Error; err != nil {
		return err
	}
	return nil
}
