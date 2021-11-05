package Models

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"
)

type ContactLabelRelation struct {
	ID        uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	ContactID uint64 `gorm:"type:bigint unsigned; not null;uniqueIndex:idx_contact_id_user_id_label_id" json: "contact_id"`
	UserId    uint64 `gorm:"type:bigint unsigned;not null;uniqueIndex:idx_contact_id_user_id_label_id" json:"user_id"`
	LabelId   uint64 `gorm:"type:bigint unsigned;not null;uniqueIndex:idx_contact_id_user_id_label_id" json:"label_id"`
	Type      uint8  `gorm:"type:tinyint(8);not null;" json:"type"`
	CreateAt  int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt  int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (lr *ContactLabelRelation) TableName() string {
	return "contact_label_relation"
}

func AddNewContactLabelRelation(lr *ContactLabelRelation) (err error) {
	if err = DB.Create(lr).Error; err != nil {
		return err
	}
	return nil
}

func GetOneContactLabelRelation(lr *ContactLabelRelation, id uint64, user_id uint64) (err error) {
	if err := DB.Where("id = ? AND user_id = ?", id, user_id).First(lr).Error; err != nil {
		return err
	}
	return nil
}

func GetContactLabelRelationsByUserId(lrs *[]ContactLabelRelation, user_id uint64, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ?", user_id).Order("id").Find(lrs).Limit(limit).Offset(offset).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetContactLabelRelationByUserIdContactIdLabelId(lr *ContactLabelRelation, user_id uint64, contactIds uint64, labelId uint64) (err error) {
	if err := DB.Where("contact_id = ? AND user_id = ? AND label_id = ?", contactIds, user_id, labelId).First(lr).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetContactLabelRelationsByUserIdAndContactIds(lrs *[]ContactLabelRelation, user_id uint64, contactIds []uint64) (err error) {
	if err := DB.Where("contact_id in (?) AND user_id = ?", contactIds, user_id).Find(lrs).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetContactLabelRelationsByUserIdAndContactId(lrs *[]ContactLabelRelation, user_id uint64, contactId uint64) (err error) {
	if err := DB.Where("contact_id = ? AND user_id = ?", contactId, user_id).Find(lrs).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetContactLabelRelationsByUserIdAndLabelId(lrs *[]ContactLabelRelation, user_id uint64, labelId uint64, limit int, offset int) (err error) {
	if err := DB.Where("label_id = ? AND user_id = ?", labelId, user_id).Order("id").Limit(limit).Offset(offset).Find(lrs).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ChangeContactLabelRelation(lrs *[]ContactLabelRelation, fromId uint64, toId uint64, user_id uint64) (err error) {

	ts := time.Now().Unix()
	contactIds := []uint64{}
	length := len(*lrs)
	if length == 0 {
		err := errors.New("wrong length")
		return err
	}

	if length == 1 {
		updates := map[string]interface{}{
			"label_id":  toId,
			"update_at": ts,
		}

		lr := (*lrs)[0]
		lr.LabelId = fromId
		contactId := lr.ContactID
		if err := DB.Debug().Model(lr).Where("contact_id = ? AND label_id = ? AND user_id = ?", contactId, fromId, user_id).Updates(updates).Error; err != nil {
			return err
		}
	} else {
		for i := range *lrs {
			lr := (*lrs)[i]
			cid := lr.ContactID
			contactIds = append(contactIds, cid)
		}
		lr := ContactLabelRelation{LabelId: fromId, UserId: user_id}
		updates := map[string]interface{}{
			"label_id":  toId,
			"update_at": ts,
		}

		if err := DB.Debug().Model(lr).Where("contact_id in (?) AND label_id = ? AND user_id = ?", contactIds, fromId, user_id).Updates(updates).Error; err != nil {
			return err
		}
	}

	return nil
}

func UpdateContactLabelRelationType(lr *ContactLabelRelation, rType uint8, user_id uint64) (err error) {

	ts := time.Now().Unix()

	updates := map[string]interface{}{
		"type":      rType,
		"update_at": ts,
	}
	if lr.ID == 0 {
		err := errors.New("invalid id")
		return err
	}

	if err := DB.Model(lr).Where("id = ? AND user_id = ?", lr.ID, user_id).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func AddContactRelations(lrs *[]ContactLabelRelation) (err error) {

	//clause.OnConflict{DoNothing: true}
	DB.Clauses(clause.OnConflict{DoNothing: true}).Create(*lrs)
	return nil
}

func RemoveContactRelations(lrs *[]ContactLabelRelation, user_id uint64) (err error) {

	contactIds := []uint64{}
	arr := []interface{}{}
	for i := range *lrs {
		lr := (*lrs)[i]
		contactId := lr.ContactID
		contactIds = append(contactIds, contactId)
		item := []interface{}{contactId, user_id, lr.LabelId}
		arr = append(arr, item)
	}

	lr := ContactLabelRelation{UserId: user_id}
	if err := DB.Unscoped().Where("(contact_id, user_id, label_id) in (?)", arr).Delete(lr).Error; err != nil {
		return err
	}
	return nil
}

func RemoveContactLabeledRelations(lrs *[]ContactLabelRelation, user_id uint64, label_id uint64) (err error) {

	contactIds := []uint64{}
	for i := range *lrs {
		lr := (*lrs)[i]
		contactId := lr.ContactID
		contactIds = append(contactIds, contactId)
	}
	lr := ContactLabelRelation{UserId: user_id, LabelId: label_id}
	if err := DB.Unscoped().Where("contact_id in (?) AND user_id = ? AND label_id = ?", contactIds, user_id, label_id).Delete(lr).Error; err != nil {
		return err
	}
	return nil
}

func PutOneContactLabelRelation(lr *ContactLabelRelation) (err error) {
	DB.Save(lr)
	return nil
}

func DeleteContactLabelRelation(lr *ContactLabelRelation, id uint64, user_id uint64) (err error) {
	if err := DB.Unscoped().Where("id = ? AND user_id = ?", id, user_id).Delete(lr).Error; err != nil {
		return err
	}
	return nil
}
