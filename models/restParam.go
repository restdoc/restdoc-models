package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type RestParam struct {
	Id        int64  `gorm:"type:bigint; not null; primary_key; " json: "id"`
	ApiId     int64  `gorm:"type:bigint; not null;  " json: "api_id"`
	Name      string `gorm:"type:varchar(255) ;not null;" json:"name"`
	Title     string `gorm:"type:varchar(255) ;not null;" json:"title"`
	Default   string `gorm:"type:varchar(1024) ;not null;" json:"default"`
	Note      string `gorm:"type:varchar(1024) ;not null;" json:"note"`
	Weight    string `gorm:"type:varchar(32) ;not null;" json:"weight"`
	Required  bool   `gorm:"type:boolean; not null; " json: "required"`
	Status    int16  `gorm:"type:smallint; not null; " json: "status"`
	Type      int16  `gorm:"type:smallint; not null; " json: "type"`
	CreatedAt int64  `gorm:"type:bigint;not null;" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint;not null;" json:"updated_at"`
}

func (rp *RestParam) TableName() string {
	return "rest_param"
}

func GetRestParamsByAPIId(rps *[]RestParam, apiId int64) (err error) {
	selectFields := []string{"id", "api_id", "name", "title", "default", "note", "weight", "required", "status", "type", "created_at"}
	if err = DB.Where("api_id = ?", apiId).Find(rps).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func GetRestParamsByIds(rps *[]RestParam, ids []int64) (err error) {
	selectFields := []string{"id", "api_id", "name", "title", "default", "note", "required", "weight", "status", "type", "created_at"}
	if err = DB.Where("id in (?)", ids).Find(rps).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func AddNewRestParam(rp *RestParam) error {
	if err := DB.Create(rp).Error; err != nil {
		return err
	}
	return nil
}

func GetOneRestParam(rp *RestParam, id int64) error {
	if err := DB.Where("id = ?", id).First(rp).Error; err != nil {
		return err
	}
	return nil
}

func UpdateParam(rp *RestParam, updates map[string]interface{}) error {
	if err := DB.Model(rp).Where("id = ?", rp.Id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func PutOneRestParam(rp *RestParam) (err error) {
	DB.Save(rp)
	return nil
}

func DeleteRestParam(rp *RestParam, id int64, creatorId int64) (err error) {
	if err := DB.Unscoped().Where("id = ? AND creator_id = ?", id, creatorId).Delete(rp).Error; err != nil {
		return err
	}
	return nil
}
