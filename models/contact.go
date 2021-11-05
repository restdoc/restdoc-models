package Models

import (
	"errors"
	"sort"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	ContactPersonal uint8 = 1
	ContactStarred  uint8 = 2
	ContactTeam     uint8 = 3
	ContactDeleted  uint8 = 20
)

type Contact struct {
	ID               uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	DomainID         uint64 `gorm:"type:bigint unsigned; not null; " json: "domain_id"`
	DomainUserID     uint64 `gorm:"type:bigint unsigned; not null; " json: "domain_user_id"`
	FirstName        string `gorm:"type:varchar(64);not null;default ''" json:"first_name"`
	MiddleName       string `gorm:"type:varchar(64);not null;default ''" json:"middle_name"`
	LastName         string `gorm:"type:varchar(64);not null;default ''" json:"last_name"`
	FirstNamePinyin  string `gorm:"type:varchar(64);not null;default ''" json:"first_name_pinyin"`
	MiddleNamePinyin string `gorm:"type:varchar(64);not null;default ''" json:"middle_name_pinyin"`
	LastNamePinyin   string `gorm:"type:varchar(64);not null;default ''" json:"last_name_pinyin"`
	NickName         string `gorm:"type:varchar(64);not null;default ''" json:"nick_name"`
	Company          string `gorm:"type:varchar(64);not null;default ''" json:"company"`
	Title            string `gorm:"type:varchar(64);not null;default ''" json:"title"`
	Email            string `gorm:"type:varchar(64);not null;default ''" json:"email"`
	Phone            string `gorm:"type:varchar(64);not null;default ''" json:"phone"`
	Country          string `gorm:"type:varchar(64);not null;default ''" json:"country"`
	Provicne         string `gorm:"type:varchar(64);not null;default ''" json:"province"`
	City             string `gorm:"type:varchar(64);not null;default ''" json:"city"`
	Address          string `gorm:"type:varchar(64);not null;default ''" json:"address"`
	Postcode         string `gorm:"type:varchar(64);not null;default ''" json:"postcode"`
	Birthday         string `gorm:"type:varchar(64);not null;default ''" json:"birthday"`       //MM/DD/YYYY
	LunarBirthday    string `gorm:"type:varchar(64);not null;default ''" json:"lunar_birthday"` //MM/DD/YYYY
	Relationship     string `gorm:"type:varchar(64);not null;default ''" json:"relationship"`
	QQ               string `gorm:"type:varchar(64);not null;default ''" json:"qq"`
	Wechat           string `gorm:"type:varchar(64);not null;default ''" json:"wechat"`
	Fb               string `gorm:"type:varchar(64);not null;default ''" json:"fb"`
	CreateAt         int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt         int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (c *Contact) TableName() string {
	return "contact"
}

func GetAllContact(cs *[]Contact) (err error) {
	if err = DB.Find(cs).Error; err != nil {
		return err
	}
	return nil
}

func AddNewContact(c *Contact) (err error) {
	if err = DB.Create(c).Error; err != nil {
		return err
	}
	return nil
}

func GetOneContactById(c *Contact, id uint64) (err error) {
	if err := DB.Where("id = ?", id).First(c).Error; err != nil {
		return err
	}
	return nil
}

func GetOneContact(c *Contact, id uint64, user_id uint64) (err error) {
	if err := DB.Where("id = ? AND domain_user_id = ?", id, user_id).First(c).Error; err != nil {
		return err
	}
	return nil
}

func GetOneTeamContact(c *Contact, id uint64, domain_id uint64) (err error) {
	if err := DB.Where("id = ? AND domain_id = ?", id, domain_id).First(c).Error; err != nil {
		return err
	}
	return nil
}

func GetContactsByUserId(cs *[]Contact, userId uint64) (err error) {
	if err := DB.Where("domain_user_id = ?", userId).Find(cs).Error; err != nil {
		return err
	}
	return nil
}

func GetContactsByDomainId(cs *[]Contact, domainId uint64) (err error) {
	if err := DB.Where("domain_id = ?", domainId).Find(cs).Error; err != nil {
		return err
	}
	return nil
}

func GetContactsByUserIdAndIds(cs *[]Contact, userId uint64, contact_ids []uint64) (err error) {
	if err := DB.Where("domain_user_id = ? AND id in (?)", userId, contact_ids).Find(cs).Error; err != nil {
		return err
	}
	return nil
}

func GetContactsByIds(cs *[]Contact, contact_ids []uint64) (err error) {
	if err := DB.Where("id in (?)", contact_ids).Find(cs).Error; err != nil {
		return err
	}
	return nil
}

func UpdateContactStates(cls *[]Contact, user_id uint64, updates map[string]interface{}) error {

	if len(*cls) == 0 {
		return errors.New("wrong length")
	}

	if len(*cls) == 1 {
		m := (*cls)[0]
		if err := DB.Model(m).Where("id = ? AND domain_user_id = ?", m.ID, user_id).Updates(updates).Error; err != nil {
			return err
		}
		return nil

	} else {
		ids := []uint64{}
		for i := range *cls {
			m := (*cls)[i]
			id := m.ID
			ids = append(ids, id)
		}
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

		m := Contact{DomainUserID: user_id}
		if err := DB.Model(m).Where("domain_user_id = ? AND id in (?)", user_id, ids).Updates(updates).Error; err != nil {
			return err
		}
		return nil
	}
}

func UpdateContact(c *Contact) error {
	ts := time.Now().Unix()
	updates := map[string]interface{}{
		"first_name":     c.FirstName,
		"last_name":      c.LastName,
		"nick_name":      c.NickName,
		"email":          c.Email,
		"phone":          c.Phone,
		"birthday":       c.Birthday,
		"lunar_birthday": c.LunarBirthday,
		"wechat":         c.Wechat,
		"company":        c.Company,
		"title":          c.Title,
		"update_at":      ts,
	}
	if err := DB.Model(c).Where("id = ? AND domain_user_id = ?", c.ID, c.DomainUserID).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func PutOneContact(c *Contact) (err error) {
	DB.Save(c)
	return nil
}

func DeleteContact(c *Contact, id uint64, user_id uint64) (err error) {
	DB.Where("id = ? AND domain_user_id = ?", id, user_id).Delete(c)
	return nil
}
