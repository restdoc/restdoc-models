package models

import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	RestAPIDefault        int16 = 0
	RestAPIFinished       int16 = 1
	RestAPIDeleted        int16 = 200
	RestAPIForeverDeleted int16 = 201
)

const (
	METHOD_GET    int16 = 0
	METHOD_POST   int16 = 1
	METHOD_OPTION int16 = 2
)

const (
	PARAM_GET       int16 = 1
	PARAM_POST_FORM int16 = 2
	PARAM_HEADER    int16 = 10
)

type RestAPI struct {
	Id        int64  `gorm:"type:bigint; not null; primary_key; " json: "id"`
	ProjectId int64  `gorm:"type:bigint; not null; index:project_id_index;" json: "project_id"`
	GroupId   int64  `gorm:"type:bigint; not null;  " json: "group_id"`
	CreatorId int64  `gorm:"type:bigint;not null;" json:"creator_id"`
	Path      string `gorm:"type:varchar(255) ;not null;" json:"path"`
	Name      string `gorm:"type:varchar(255);not null;" json:"name"`
	BaseUrl   string `gorm:"type:varchar(255) ;not null;" json:"base_url"`
	Method    int16  `gorm:"type:smallint; not null; " json: "method"`
	Weight    string `gorm:"type:varchar(32) ;not null;" json:"weight"`
	Status    int16  `gorm:"type:smallint ; not null; " json: "status"`
	CreatedAt int64  `gorm:"type:bigint;not null;" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint;not null;" json:"updated_at"`
}

func (ra *RestAPI) TableName() string {
	return "rest_api"
}

func GetRestAPIsByProjectId(apis *[]RestAPI, projectId int64) (err error) {
	selectFields := []string{"id", "project_id", "creator_id", "weight", "name", "status", "created_at"}
	if err = DB.Where("project_id = ?", projectId).Find(apis).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func GetRestAPIsByIds(apis *[]RestAPI, ids []int64) (err error) {
	selectFields := []string{"id", "project_id", "creator_id", "weight", "name", "status", "created_at"}
	if err = DB.Where("id in (?)", ids).Find(apis).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func AddNewRestAPI(ls *RestAPI) error {
	if err := DB.Create(ls).Error; err != nil {
		return err
	}
	return nil
}

func GetOneRestAPI(ls *RestAPI, id int64) error {
	if err := DB.Where("id = ?", id).First(ls).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRestAPI(ls *RestAPI, updates map[string]interface{}) error {
	if err := DB.Model(ls).Where("id = ?", ls.Id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func PutOneRestAPI(ls *RestAPI) (err error) {
	DB.Save(ls)
	return nil
}

func DeleteRestAPIsByProjectId(apis []RestAPI, id int64, creatorId int64) error {
	if err := DB.Unscoped().Where("project_id = ? AND creator_id = ?", id, creatorId).Delete(&apis).Error; err != nil {
		return err
	}
	return nil

}

func DeleteRestAPI(ls *RestAPI, id int64, creatorId int64) (err error) {
	if err := DB.Unscoped().Where("id = ? AND creator_id = ?", id, creatorId).Delete(ls).Error; err != nil {
		return err
	}
	return nil
}
