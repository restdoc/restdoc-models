package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Team struct {
	Id       int64  `gorm:"type:bigint; not null; primary_key; " json: "id"`
	UserId   int64  `gorm:"type:bigint;not null;" json:"user_id"`
	Name     string `gorm:"type:varchar(128);idx_name;not null" json:"name"`
	Type     int16  `gorm:"type:smallint;not null;" json:"type"`
	Valid    bool   `gorm:"type:boolean;not null;default false;" json:"valid"`
	CreateAt int64  `gorm:"type:bigint;not null;" json:"create_at"`
	UpdateAt int64  `gorm:"type:bigint;not null;" json:"update_at"`
}

func (t *Team) TableName() string {
	return "team"
}

func GetAllTeam(teams *[]Team) (err error) {
	if err = DB.Find(teams).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTeamName(t *[]Team, created int64) (err error) {
	if err = DB.Where("created_at > ?", created).Find(t).Select([]string{"name", "created_at"}).Error; err != nil {
		return err
	}
	return nil
}

func AddNewTeam(t *Team) (err error) {
	if err = DB.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func GetOneTeam(t *Team, id string) (err error) {
	if err := DB.Where("id = ?", id).First(t).Error; err != nil {
		return err
	}
	return nil
}

func GetTeamsByUserId(teams *[]Team, user_id int64) (err error) {
	if err := DB.Where("user_id = ?", user_id).Find(&teams).Error; err != nil {
		return err
	}
	return nil
}

func PutOneTeam(t *Team) (err error) {
	DB.Save(t)
	return nil
}

func UpdateTeam(t *Team, updates map[string]interface{}) error {

	if err := DB.Model(t).Where("user_id = ? AND id = ?", t.UserId, t.Id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTeam(t *Team, teamId int64, userId int64) (err error) {
	//if err := DB.Unscoped().Where("user_id = ? AND id = ?", user_id, id).Delete(t).Error; err != nil {
	if err := DB.Where("id = ? AND user_id = ?", teamId, userId).Delete(t).Error; err != nil {
		return err
	}
	return nil
}
