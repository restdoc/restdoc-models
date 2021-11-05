package Models

import (
	_ "github.com/go-sql-driver/mysql"
)

type KanbanProject struct {
	Id        uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	TeamId    uint64 `gorm:"type:bigint unsigned; not null;  " json: "team_id"`
	CreatorId uint64 `gorm:"type:bigint unsigned; not null;  " json: "creator_id"`
	Name      string `gorm:"type:varchar(255);not null;" json:"name"`
	Status    uint8  `gorm:"type:tinyint unsigned; not null; " json: "status"`
	CreatedAt int64  `gorm:"type:bigint unsigned;not null;" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint unsigned;not null;" json:"updated_at"`
}

func (pr *KanbanProject) TableName() string {
	return "kanban_project"
}

func GetUserProjects(prs *[]KanbanProject, userId uint64) (err error) {
	selectFields := []string{"id", "team_id", "creator_id", "name", "status", "created_at", "updated_at"}
	if err = DB.Where("creator_id = ?", userId).Find(prs).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func GetHomeProjects(prs *[]KanbanProject, teamIds []uint64) (err error) {
	selectFields := []string{"id", "team_id", "creator_id", "name", "status", "created_at", "updated_at"}
	if err = DB.Where("team_id in (?)", teamIds).Find(prs).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func AddNewProject(pr *KanbanProject) error {
	if err := DB.Create(pr).Error; err != nil {
		return err
	}
	return nil
}

func GetOneProject(pr *KanbanProject, id uint64) error {
	if err := DB.Where("id = ?", id).First(pr).Error; err != nil {
		return err
	}
	return nil
}

func PutOneProject(pr *KanbanProject) (err error) {
	DB.Save(pr)
	return nil
}

func DeleteProject(pr *KanbanProject, id uint64, creator_id uint64) (err error) {
	if err := DB.Unscoped().Where("id = ? AND creator_id = ?", id, creator_id).Delete(pr).Error; err != nil {
		return err
	}
	return nil
}
