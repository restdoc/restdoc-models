package Models

import (
	_ "github.com/go-sql-driver/mysql"
)

type KanbanList struct {
	Id        uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	ProjectId uint64 `gorm:"type:bigint unsigned; not null;  " json: "project_id"`
	CreatorId uint64 `gorm:"type:bigint unsigned;not null;" json:"creator_id"`
	Weight    string `gorm:"type:varchar(32) ;not null;" json:"weight"`
	Name      string `gorm:"type:varchar(255);not null;" json:"name"`
	Status    uint8  `gorm:"type:tinyint unsigned; not null; " json: "status"`
	CreatedAt int64  `gorm:"type:bigint unsigned;not null;" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint unsigned;not null;" json:"updated_at"`
}

func (l *KanbanList) TableName() string {
	return "kanban_list"
}

func GetKanbanListByProjectId(ls *[]KanbanList, projectId uint64) (err error) {
	selectFields := []string{"id", "project_id", "creator_id", "weight", "name", "status", "created_at"}
	if err = DB.Where("project_id = ?", projectId).Find(ls).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func GetKanbanListByIds(ls *[]KanbanList, ids []uint64) (err error) {
	selectFields := []string{"id", "project_id", "creator_id", "weight", "name", "status", "created_at"}
	if err = DB.Where("id in (?)", ids).Find(ls).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func AddNewKanbanList(ls *KanbanList) error {
	if err := DB.Create(ls).Error; err != nil {
		return err
	}
	return nil
}

func GetOneKanbanList(ls *KanbanList, id uint64) error {
	if err := DB.Where("id = ?", id).First(ls).Error; err != nil {
		return err
	}
	return nil
}

func UpdateList(ls *KanbanList, updates map[string]interface{}) error {
	if err := DB.Model(ls).Where("id = ?", ls.Id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func PutOneKanbanList(ls *KanbanList) (err error) {
	DB.Save(ls)
	return nil
}

func DeleteKanbanList(ls *KanbanList, id uint64, creatorId uint64) (err error) {
	if err := DB.Unscoped().Where("id = ? AND creator_id = ?", id, creatorId).Delete(ls).Error; err != nil {
		return err
	}
	return nil
}

func DeleteKanbanListsByProjectId(projectId uint64, creatorId uint64) (err error) {
	DB.Where("project_id = ? AND creator_Id = ?", projectId, creatorId).Delete(KanbanList{})
	return nil
}
