package Models

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*

inbox starred important

all    flag: inbox starred important answerd snoozed draft sent
spam
deleted

*/

type PrivateLabel struct {
	Id        uint8
	Name      string
	LowerName string
	Flag      string
}

var AllMailbox = PrivateLabel{Id: 0, Name: "All", LowerName: "all", Flag: ""}
var InboxMailbox = PrivateLabel{Id: 1, Name: "INBOX", LowerName: "inbox", Flag: "\\inbox"}
var SnoozedMailbox = PrivateLabel{Id: 10, Name: "Snoozed", LowerName: "snoozed", Flag: ""}
var SentMailbox = PrivateLabel{Id: 11, Name: "Sent", LowerName: "sent", Flag: ""}
var DraftsMailbox = PrivateLabel{Id: 12, Name: "Drafts", LowerName: "drafts", Flag: "\\draft"}
var StarredMailbox = PrivateLabel{Id: 190, Name: "Starred", LowerName: "starred", Flag: "\\flagged"}
var ImportantMailbox = PrivateLabel{Id: 191, Name: "Important", LowerName: "important", Flag: "\\important"}
var DeletedMailbox = PrivateLabel{Id: 201, Name: "Trash", LowerName: "trash", Flag: "\\deleted"}
var SpamMailbox = PrivateLabel{Id: 202, Name: "Spam", LowerName: "spam"}

var ReadedMailbox = PrivateLabel{Id: 192, Name: "Readed", LowerName: "readed", Flag: "\\seen"}
var UnreadMailbox = PrivateLabel{Id: 192, Name: "Unread", LowerName: "unread", Flag: "\\unseen"}
var DeletedForeverMailbox = PrivateLabel{Id: 200, Name: "Deleteforever", LowerName: "deleteforever", Flag: ""}
var AnsweredMailbox = PrivateLabel{Id: 193, Name: "Answered", LowerName: "answered", Flag: "\\answered"}

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
	IsInbox          bool   `gorm:"type:boolean;not null; default false" json:"is_inbox"`
	IsImportant      bool   `gorm:"type:boolean;not null; default false" json:"is_important"`
	IsStarred        bool   `gorm:"type:boolean;not null; default false" json:"is_starred"`
	IsRead           bool   `gorm:"type:boolean;not null; default false" json:"is_read"`
	IsDraft          bool   `gorm:"type:boolean;not null; default false" json:"is_draft"`
	IsSent           bool   `gorm:"type:boolean;not null; default false" json:"is_sent"`
	IsAnswerd        bool   `gorm:"type:boolean;not null; default false" json:"is_answerd"`
	IsSnoozed        bool   `gorm:"type:boolean;not null; default false" json:"is_snoozed"`
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

func GetMailsByIdArray(mails *[]Mail, ids []int64) (err error) {
	if err := DB.Where("id in (?)", ids).Order("id desc").Find(mails).Error; err != nil {
		return err
	}
	return nil
}

func GetMailIds(list *[]Mail) (err error) {
	if err := DB.Where("id > 0").Select([]string{"id", "user_id", "uid"}).Find(list).Error; err != nil {
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
	if err := DB.Where("user_id = ? AND `type` = ?", user_id, InboxMailbox.Id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
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

func GetFlagedMailsByUserId(flag string, mails *[]Mail, user_id uint64, limit int, offset int) (err error) {
	switch flag {
	case "all":
		if err := DB.Where("user_id = ? AND type = 0", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
			fmt.Println(err)
			return err
		}
	case "inbox":
		if err := DB.Where("user_id = ? AND is_inbox = 1 AND type = 0", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
			fmt.Println(err)
			return err
		}
	case "important":
		if err := DB.Where("user_id = ? AND is_important = 1 AND type = 0", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
			fmt.Println(err)
			return err
		}
	case "starred":
		if err := DB.Where("user_id = ? AND is_starred = 1 AND type = 0", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
			fmt.Println(err)
			return err
		}
	case "read":
		if err := DB.Where("user_id = ? AND is_read = 1 AND type = 0", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
			fmt.Println(err)
			return err
		}
	case "draft", "drafts":
		if err := DB.Where("user_id = ? AND is_draft = 1 AND type = 0", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
			fmt.Println(err)
			return err
		}
	case "sent", "send":
		if err := DB.Where("user_id = ? AND is_sent = 1 AND type = 0", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
			fmt.Println(err)
			return err
		}
	case "answerd":
		if err := DB.Where("user_id = ? AND is_answerd = 1 AND type = 0", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
			fmt.Println(err)
			return err
		}
	case "snoozed":
		if err := DB.Where("user_id = ? AND is_snoozed = 1 AND type = 0", user_id).Order("id desc").Limit(limit).Offset(offset).Find(mails).Error; err != nil {
			fmt.Println(err)
			return err
		}
	default:
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

func GetAllMails(mails *[]Mail) (err error) {
	if err := DB.Select([]string{"id", "user_id", "type", "is_read", "is_starred", "is_important"}).Find(mails).Error; err != nil {
		return err
	}
	return nil
}

func GetAllMailsId(mails *[]Mail) (err error) {
	if err := DB.Where("id > 0").Select([]string{"id", "user_id"}).Find(mails).Error; err != nil {
		return err
	}
	return nil
}

func CheckMailboxName(name string) bool {
	ok := true
	switch strings.ToLower(name) {
	case AllMailbox.LowerName,
		InboxMailbox.LowerName,
		SnoozedMailbox.LowerName,
		SentMailbox.LowerName,
		DraftsMailbox.LowerName,
		StarredMailbox.LowerName,
		ImportantMailbox.LowerName,
		DeletedMailbox.LowerName,
		SpamMailbox.LowerName,
		ReadedMailbox.LowerName,
		DeletedForeverMailbox.LowerName:
		ok = false
	default:
	}
	return ok
}

func CheckMailboxId(id uint64) bool {
	ok := true
	switch id {
	case uint64(AllMailbox.Id),
		uint64(InboxMailbox.Id),
		uint64(SnoozedMailbox.Id),
		uint64(SentMailbox.Id),
		uint64(DraftsMailbox.Id),
		uint64(StarredMailbox.Id),
		uint64(ImportantMailbox.Id),
		uint64(DeletedMailbox.Id),
		uint64(SpamMailbox.Id),
		uint64(ReadedMailbox.Id),
		uint64(DeletedForeverMailbox.Id):
		ok = false
	default:
	}
	return ok
}
