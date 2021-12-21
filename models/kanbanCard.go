package Models

import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	CardDefault        uint8 = 0
	CardFinished       uint8 = 1
	CardDeleted        uint8 = 200
	CardForeverDeleted uint8 = 201
)

type KanbanCard struct {
	Id          uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	ProjectId   uint64 `gorm:"type:bigint unsigned; not null; index:project_id_index;" json: "project_id"`
	ListId      uint64 `gorm:"type:bigint unsigned; not null; index:list_id_index;" json: "list_id"`
	Title       string `gorm:"type:varchar(256);not null" json:"title"`
	Status      uint8  `gorm:"type:tinyint unsigned; not null; " json: "status"`
	CreatorId   uint64 `gorm:"type:bigint unsigned; not null; " json: "creator_id"`
	PrincipalId uint64 `gorm:"type:bigint unsigned; not null; " json: "principal_id"`
	Due         string `gorm:"type:varchar(128);not null;" json:"due"`
	Desc        string `gorm:"type:varchar(1024);not null;" json:"desc"`
	Weight      string `gorm:"type:varchar(32) ;not null;" json:"weight"`
	CreatedAt   int64  `gorm:"type:bigint unsigned;not null;" json:"created_at"`
	UpdatedAt   int64  `gorm:"type:bigint unsigned;not null;" json:"updated_at"`
}

func (cd *KanbanCard) TableName() string {
	return "kanban_card"
}

func GetKanbanCardsByProjectId(list *[]KanbanCard, project_id uint64) (err error) {
	if err = DB.Where("project_id = ?", project_id).Find(list).Error; err != nil {
		return err
	}
	return nil
}

func GetKanbanCardsByIds(list *[]KanbanCard, ids []uint64) (err error) {
	if err = DB.Where("id in (?)", ids).Find(list).Error; err != nil {
		return err
	}
	return nil
}

func AddNewKanbanCard(c *KanbanCard) (err error) {
	if err = DB.Create(c).Error; err != nil {
		return err
	}
	return nil
}

func GetOneKanbanCard(c *KanbanCard, id string) (err error) {
	if err := DB.Where("id = ?", id).First(c).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCard(c *KanbanCard, updates map[string]interface{}) error {
	if err := DB.Model(c).Where("id = ?", c.Id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func PutOneKanbanCard(c *KanbanCard) (err error) {
	DB.Save(c)
	return nil
}

func DeleteKanbanCard(c *KanbanCard, id string) (err error) {
	DB.Where("id = ?", id).Delete(c)
	return nil
}

func DeleteKanbanCards(ids []uint64) (err error) {
	DB.Delete(&KanbanCard{}, ids)
	return nil
}

func DeleteKanbanCardsByProjectId(projectId uint64, creatorId uint64) (err error) {
	DB.Where("project_id = ? AND creator_Id = ?", projectId, creatorId).Delete(KanbanCard{})
	return nil
}
