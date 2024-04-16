package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type RestProject struct {
	Id        int64  `gorm:"type:bigint; not null; primary_key; " json: "id"`
	TeamId    int64  `gorm:"type:bigint; not null;  " json: "team_id"`
	CreatorId int64  `gorm:"type:bigint; not null;  " json: "creator_id"`
	Name      string `gorm:"type:varchar(255);not null;" json:"name"`
	BaseUrl   string `gorm:"type:varchar(255);not null;" json:"base_url"`
	Weight    string `gorm:"type:varchar(32) ;not null;" json:"weight"`
	Status    int16  `gorm:"type:smallint; not null; " json: "status"`
	NameColor string `gorm:"type:varchar(16);not null;" json:"name_color"`
	Color     string `gorm:"type:varchar(16);not null;" json:"color"`
	Icon      string `gorm:"type:varchar(32);not null;" json:"icon"`
	IconColor string `gorm:"type:varchar(16);not null;" json:"icon_color"`
	CreatedAt int64  `gorm:"type:bigint;not null;" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint;not null;" json:"updated_at"`
}

func (rp *RestProject) TableName() string {
	return "rest_project"
}

func GetUserRestProjects(rps []*RestProject, userId int64) (err error) {
	selectFields := []string{"id", "team_id", "creator_id", "name", "status", "icon", "color", "name_color", "icon_color", "created_at", "updated_at"}
	if err = DB.Where("creator_id = ?", userId).Find(rps).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func GetHomeRestProjects(rps *[]RestProject, teamIds []int64) (err error) {
	selectFields := []string{"id", "team_id", "creator_id", "name", "status", "weight", "created_at", "updated_at"}
	if err = DB.Where("team_id in (?)", teamIds).Find(rps).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func AddNewRestProject(rp *RestProject) error {
	if err := DB.Create(rp).Error; err != nil {
		return err
	}
	return nil
}

func GetOneRestProject(rp *RestProject, id int64) error {
	if err := DB.Where("id = ?", id).First(rp).Error; err != nil {
		return err
	}
	return nil
}

func GetRestProjectsByIds(prs *[]RestProject, ids []int64) (err error) {
	selectFields := []string{"id", "team_id", "creator_id", "name", "status", "icon", "color", "created_at", "updated_at"}
	if err = DB.Where("id in (?)", ids).Find(prs).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func PutOneRestProject(rp *RestProject) (err error) {
	DB.Save(rp)
	return nil
}

func UpdateRestProject(rp *RestProject, updates map[string]interface{}) error {
	if err := DB.Model(rp).Where("id = ? AND creator_id = ?", rp.Id, rp.CreatorId).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProject(pr *RestProject, updates map[string]interface{}) error {
	if err := DB.Model(pr).Where("id = ?", pr.Id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRestProject(rp *RestProject, id int64, creator_id int64) (err error) {
	if err := DB.Unscoped().Where("id = ? AND creator_id = ?", id, creator_id).Delete(rp).Error; err != nil {
		return err
	}
	return nil
}
