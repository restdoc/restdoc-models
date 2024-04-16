package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type RestGroup struct {
	Id        int64  `gorm:"type:bigint; not null; primary_key; " json: "id"`
	UserId    int64  `gorm:"type:bigint;not null;" json:"user_id"`
	ProjectId int64  `gorm:"type:bigint; not null;  " json: "project_id"`
	Name      string `gorm:"type:varchar(128);idx_name;not null" json:"name"`
	Weight    string `gorm:"type:varchar(32) ;not null;" json:"weight"`
	Type      int16  `gorm:"type:smallint;not null;" json:"type"`
	CreatedAt int64  `gorm:"type:bigint;not null;" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint;not null;" json:"updated_at"`
}

func (rg *RestGroup) TableName() string {
	return "rest_group"
}

func GetAllRestGroup(groups *[]RestGroup) (err error) {
	if err = DB.Find(groups).Error; err != nil {
		return err
	}
	return nil
}

func GetAllGroupName(rgs *[]RestGroup, created int64) (err error) {
	if err = DB.Where("created_at > ?", created).Find(rgs).Select([]string{"name", "created_at"}).Error; err != nil {
		return err
	}
	return nil
}

func AddNewRestGroup(rg *RestGroup) (err error) {
	if err = DB.Create(rg).Error; err != nil {
		return err
	}
	return nil
}

func GetOneRestGroup(rg *RestGroup, id int64, user_id int64) (err error) {
	if err := DB.Where("id = ? AND user_id = ?", id, user_id).First(rg).Error; err != nil {
		return err
	}
	return nil
}

func GetRestGroupsByIds(groups *[]RestGroup, user_id int64, ids []int64) (err error) {
	if err := DB.Where("user_id = ? AND id in (?)", user_id, ids).Find(&groups).Error; err != nil {
		return err
	}
	return nil
}

func GetRestGroupsByUserId(groups *[]RestGroup, user_id int64) (err error) {
	if err := DB.Where("user_id = ?", user_id).Find(&groups).Error; err != nil {
		return err
	}
	return nil
}

func GetRestGroupsByProjectId(groups *[]RestGroup, project_id int64) (err error) {
	if err := DB.Where("project_id = ?", project_id).Find(&groups).Error; err != nil {
		return err
	}
	return nil
}

func PutOneRestGroup(rg *RestGroup) (err error) {
	DB.Save(rg)
	return nil
}

func UpdateRestGroup(rg *RestGroup, updates map[string]interface{}) error {

	if err := DB.Model(rg).Where("user_id = ? AND id = ?", rg.UserId, rg.Id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRestGroupsByProjectId(groups []RestGroup, id int64, creatorId int64) error {
	if err := DB.Unscoped().Where("project_id = ? AND creator_id = ?", id, creatorId).Delete(&groups).Error; err != nil {
		return err
	}
	return nil

}

func DeleteRestGroup(rg *RestGroup, user_id int64) (err error) {
	//if err := DB.Unscoped().Where("user_id = ? AND id = ?", user_id, id).Delete(t).Error; err != nil {
	if err := DB.Where("user_id = ? AND id = ?", user_id, rg.Id).Delete(rg).Error; err != nil {
		return err
	}
	return nil
}
