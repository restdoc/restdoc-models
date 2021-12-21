package Models

import (
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type LabelType int8

const (
	SystemUsefulLabel  LabelType = 0 // inbox, important, starred
	SystemUselessLabel LabelType = 1 // spam, deleted
	CustomLabel        LabelType = 2
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
*/

func UpdateLabelName(l *Label, newName string) error {
	updates := map[string]interface{}{
		"label_name": newName,
	}
	if err := DB.Model(l).Where("id = ? AND user_id = ?", l.ID, l.UserId).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func DeleteLabel(l *Label, id uint64, user_id uint64) (err error) {
	if err := DB.Unscoped().Where("id = ? AND user_id = ?", id, user_id).Delete(l).Error; err != nil {
		return err
	}
	return nil
}

func DeleteLabelByName(l *Label, name string, user_id uint64) (err error) {
	if err := DB.Unscoped().Where("user_id = ? AND name = ?", user_id, name).Delete(l).Error; err != nil {
		return err
	}
	return nil
}

func GetLabelType(labelName string, labels map[string]Label) (Label, LabelType, error) {

	lbType := SystemUsefulLabel

	switch strings.ToLower(labelName) {
	case AllMailbox.LowerName:
		lb := Label{ID: uint64(AllMailbox.Id), Name: AllMailbox.Name}
		return lb, lbType, nil
	case InboxMailbox.LowerName:
		lb := Label{ID: uint64(InboxMailbox.Id), Name: InboxMailbox.Name}
		return lb, lbType, nil
	case SnoozedMailbox.LowerName:
		lb := Label{ID: uint64(SnoozedMailbox.Id), Name: SnoozedMailbox.Name}
		return lb, lbType, nil
	case SentMailbox.LowerName:
		lb := Label{ID: uint64(SentMailbox.Id), Name: SentMailbox.Name}
		return lb, lbType, nil
	case DraftsMailbox.LowerName:
		lb := Label{ID: uint64(DraftsMailbox.Id), Name: DraftsMailbox.Name}
		return lb, lbType, nil
	case StarredMailbox.LowerName:
		lb := Label{ID: uint64(StarredMailbox.Id), Name: StarredMailbox.Name}
		return lb, lbType, nil
	case ImportantMailbox.LowerName:
		lb := Label{ID: uint64(ImportantMailbox.Id), Name: ImportantMailbox.Name}
		return lb, lbType, nil
	case DeletedMailbox.LowerName:
		lb := Label{ID: uint64(DeletedMailbox.Id), Name: DeletedMailbox.Name}
		lbType = SystemUselessLabel
		return lb, lbType, nil
	case SpamMailbox.LowerName:
		lb := Label{ID: uint64(SpamMailbox.Id), Name: SpamMailbox.Name}
		lbType = SystemUselessLabel
		return lb, lbType, nil
	case ReadedMailbox.LowerName:
		lb := Label{ID: uint64(ReadedMailbox.Id), Name: ReadedMailbox.Name}
		return lb, lbType, nil
	case DeletedForeverMailbox.LowerName:
		lb := Label{ID: uint64(DeletedForeverMailbox.Id), Name: DeletedForeverMailbox.Name}
		lbType = SystemUselessLabel
		return lb, lbType, nil
	default:
	}

	l, ok := labels[labelName]
	if !ok {
		err := errors.New("invalid label")
		return l, lbType, err
	} else {
		lbType = CustomLabel
	}
	return l, lbType, nil
}

func GenLabelUpdates(fromType LabelType, toType LabelType, fromLabel Label, toLabel Label, user_id uint64) (map[string]interface{}, error) {
	var err error
	key := fmt.Sprintf("%d-%d", fromType, toType)
	updates := map[string]interface{}{}
	switch key {
	case "0-1":
		switch strings.ToLower(toLabel.Name) {
		case "deleted", "trash":
			updates = map[string]interface{}{
				"is_inbox": false,
				"type":     DeletedMailbox.Id,
			}
		case "spam":
			updates = map[string]interface{}{
				"is_inbox": false,
				"type":     SpamMailbox.Id,
			}
		default:
			err = errors.New("错误的新标签")
		}
		return updates, err
	case "1-0":
		switch strings.ToLower(toLabel.Name) {
		case "inbox":
			updates = map[string]interface{}{
				"is_inbox": true,
				"type":     AllMailbox.Id,
			}
		default:
			err = errors.New("错误的新标签")
		}
		return updates, err
	case "0-2":
		if strings.ToLower(fromLabel.Name) == "inbox" {
			updates = map[string]interface{}{
				"is_inbox": false,
			}
		} else {
			updates = map[string]interface{}{}
		}
		return updates, err
	case "2-0":
		updates = map[string]interface{}{
			"is_inbox": true,
		}
		/*
			err = removeLabel(c, relationships, user_id, fromLabel.ID)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error(), "code": 1, "message": "删除标签失败 更新失败"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除标签成功 更新成功"})
		*/
		return updates, err
	case "1-2":
		switch strings.ToLower(fromLabel.Name) {
		case "deleted", "trash", "spam":
			updates = map[string]interface{}{
				"type": AllMailbox.Id,
			}
		default:
			err = errors.New("错误的旧标签")
		}
		return updates, err
	case "1-1":
		switch strings.ToLower(toLabel.Name) {
		case "deleted", "trash":
			updates = map[string]interface{}{
				"is_inbox": false,
				"type":     DeletedMailbox.Id,
			}
		case "spam":
			updates = map[string]interface{}{
				"is_inbox": false,
				"type":     SpamMailbox.Id,
			}
		default:
			err = errors.New("错误的旧标签")
		}
		return updates, err
	case "2-1":
		switch strings.ToLower(toLabel.Name) {
		case "deleted", "trash":
			updates = map[string]interface{}{
				"is_inbox": false,
				"type":     DeletedMailbox.Id,
			}
		case "spam":
			updates = map[string]interface{}{
				"is_inbox": false,
				"type":     SpamMailbox.Id,
			}
		default:
			err = errors.New("错误的旧标签")
		}
		return updates, err
	case "2-2":
		return updates, nil
	}
	return updates, nil
}
