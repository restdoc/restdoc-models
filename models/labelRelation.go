package Models

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"
)

type LabelRelation struct {
	ID       uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	MailID   uint64 `gorm:"type:bigint unsigned; not null;uniqueIndex:idx_mail_id_label_id;" json: "mail_id"`
	UserId   uint64 `gorm:"type:bigint unsigned;not null;" json:"user_id"`
	LabelId  uint64 `gorm:"type:bigint unsigned;not null;uniqueIndex:idx_mail_id_label_id;" json:"label_id"`
	Type     uint8  `gorm:"type:tinyint(8);not null;" json:"type"`
	CreateAt int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (lr *LabelRelation) TableName() string {
	return "label_relation"
}

func AddNewLabelRelation(lr *LabelRelation) (err error) {
	if err = DB.Create(lr).Error; err != nil {
		return err
	}
	return nil
}

func GetOneLabelRelation(lr *LabelRelation, id uint64, user_id uint64) (err error) {
	if err := DB.Where("id = ? AND user_id = ?", id, user_id).First(lr).Error; err != nil {
		return err
	}
	return nil
}

func GetAllLabelRelations(lrs *[]LabelRelation) (err error) {
	if err := DB.Select([]string{"mail_id", "user_id", "label_id"}).Find(lrs).Error; err != nil {
		return err
	}
	return nil
}

func GetLabelRelationsByUserId(lrs *[]LabelRelation, user_id uint64, limit int, offset int) (err error) {
	if err := DB.Where("user_id = ?", user_id).Order("id").Find(lrs).Limit(limit).Offset(offset).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetLabelRelationsByUserIdAndMailIds(lrs *[]LabelRelation, user_id uint64, mailIds []uint64) (err error) {
	if err := DB.Where("mail_id in (?) AND user_id = ?", mailIds, user_id).Find(lrs).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetLabelRelationsByUserIdAndMailId(lrs *[]LabelRelation, user_id uint64, mailId uint64) (err error) {
	if err := DB.Where("mail_id = ? AND user_id = ?", mailId, user_id).Find(lrs).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetLabelRelationsByUserIdAndLabelId(lrs *[]LabelRelation, user_id uint64, labelId uint64, limit int, offset int) (err error) {
	if err := DB.Where("label_id = ? AND user_id = ?", labelId, user_id).Order("id").Limit(limit).Offset(offset).Find(lrs).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ChangeMailLabelRelation(lrs *[]LabelRelation, fromId uint64, toId uint64, user_id uint64) (err error) {

	ts := time.Now().Unix()
	mailIds := []uint64{}
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
		mailId := lr.MailID
		if err := DB.Debug().Model(lr).Where("mail_id = ? AND label_id = ? AND user_id = ?", mailId, fromId, user_id).Updates(updates).Error; err != nil {
			return err
		}
	} else {
		for i := range *lrs {
			lr := (*lrs)[i]
			mid := lr.MailID
			mailIds = append(mailIds, mid)
		}
		lr := LabelRelation{LabelId: fromId, UserId: user_id}
		updates := map[string]interface{}{
			"label_id":  toId,
			"update_at": ts,
		}

		if err := DB.Debug().Model(lr).Where("mail_id in (?) AND label_id = ? AND user_id = ?", mailIds, fromId, user_id).Updates(updates).Error; err != nil {
			return err
		}
	}

	return nil
}

func AddRelations(lrs *[]LabelRelation) (err error) {
	DB.Clauses(clause.OnConflict{DoNothing: true}).Create(*lrs)
	return nil
}

func RemoveRelations(lrs *[]LabelRelation, user_id uint64) (err error) {

	mailIds := []uint64{}
	arr := []interface{}{}
	for i := range *lrs {
		lr := (*lrs)[i]
		mailId := lr.MailID
		mailIds = append(mailIds, mailId)
		item := []interface{}{lr.MailID, user_id, lr.LabelId}
		arr = append(arr, item)
	}

	lr := LabelRelation{UserId: user_id}
	if err := DB.Unscoped().Where("(mail_id, user_id, label_id) in (?)", arr).Delete(lr).Error; err != nil {
		return err
	}
	return nil
}

func RemoveLabeledRelations(lrs *[]LabelRelation, user_id uint64, label_id uint64) (err error) {

	mailIds := []uint64{}
	for i := range *lrs {
		lr := (*lrs)[i]
		mailId := lr.MailID
		mailIds = append(mailIds, mailId)
	}
	lr := LabelRelation{UserId: user_id, LabelId: label_id}
	if err := DB.Unscoped().Where("mail_id in (?) AND user_id = ? AND label_id = ?", mailIds, user_id, label_id).Delete(lr).Error; err != nil {
		return err
	}
	return nil
}

func RemoveLabeledRelationsByLabelId(user_id uint64, label_id uint64) (err error) {
	lr := LabelRelation{UserId: user_id, LabelId: label_id}
	if err := DB.Unscoped().Where("label_id = ? AND user_id = ?", label_id, user_id).Delete(lr).Error; err != nil {
		return err
	}
	return nil
}

func PutOneLabelRelation(lr *LabelRelation) (err error) {
	DB.Save(lr)
	return nil
}

/*
func UpdateStatus(lr *LabelRelation) error {
	ts := time.Now().Unix()
	if err := DB.Model(lr).Where("id = ? AND user_id = ?", lr.ID, lr.UserId).Updates(LabelRelation{Valid: d.Valid, CurrentSpfRecord: d.CurrentSpfRecord, CurrentDkimRecord: d.CurrentDkimRecord, CurrentDmarcRecord: d.CurrentDmarcRecord, CurrentMxaRecord: d.CurrentMxaRecord, CurrentMxbRecord: d.CurrentMxbRecord, UpdateAt: ts}).Error; err != nil {
		return err
	}
	return nil
}
*/

func DeleteLabelRelation(lr *LabelRelation, id uint64, user_id uint64) (err error) {
	if err := DB.Unscoped().Where("id = ? AND user_id = ?", id, user_id).Delete(lr).Error; err != nil {
		return err
	}
	return nil
}
