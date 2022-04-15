package RestdocModels

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type TeamUser struct {
	Id       int64  `gorm:"type:bigint; not null; primary_key; " json: "id"`
	TeamId   int64  `gorm:"type:bigint; not null; " json: "team_id"`
	UserId   int64  `gorm:"type:bigint; not null; uniqueIndex:idx_team_id_user_id" json: "user_id"`
	Locale   string `gorm:"type:varchar(64);not null;" json:"locale"`
	Valid    bool   `gorm:"type:boolean;not null;default false;" json:"valid"`
	Type     int16  `gorm:"type:smallint;not null;default 0" json:"type"`
	CreateAt int64  `gorm:"type:bigint;not null;" json:"create_at"`
	UpdateAt int64  `gorm:"type:bigint;not null;" json:"update_at"`
}

func (u *TeamUser) TableName() string {
	return "team_user"
}

func MembersCount() int64 {
	var count int64
	DB.Model(&TeamUser{}).Where("id > ?", 0).Count(&count)
	return count
}

func GetAllTeamUser(u *[]TeamUser) (err error) {
	if err = DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewTeamUser(u *TeamUser) (err error) {
	if err = DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneTeamUser(u *TeamUser, id int64) (err error) {
	if err := DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetTeamUserByTeamIdAndUser(u *TeamUser, teamId int64, email string) (err error) {
	if err := DB.Where("team_id = ? AND email = ?", teamId, email).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetTeamUsersByTeamId(u *[]TeamUser, team_id string) (err error) {
	if err := DB.Where("team_id = ?", team_id).Find(u).Error; err != nil {
		return err
	}
	return nil
}

func VerifyTeamUserByCode(u *TeamUser, teamId int64, email string, code string) (err error) {
	if err := DB.Where("team_id = ? AND email = ? AND verify_code = ?", teamId, email, code).First(u).Error; err != nil {
		return err
	}
	if u.Valid == true {
		return nil
	}

	ts := time.Now().Unix()
	if err := DB.Model(u).Where("id = ?", u.Id).Updates(map[string]interface{}{"update_at": ts, "valid": true, "verify_code": ""}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTeamUserLocale(u *TeamUser, locale string) (err error) {
	ts := time.Now().Unix()
	if err := DB.Model(u).Where("id = ?", u.Id).Updates(map[string]interface{}{"locale": locale, "update_at": ts}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTeamUserPassword(u *TeamUser, team_id int64, password string) (err error) {
	ts := time.Now().Unix()
	if err := DB.Model(u).Where("id = ? AND team_id = ?", u.Id, u.TeamId).Updates(map[string]interface{}{"password": password, "update_at": ts}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTeamUserVerifyCode(u *TeamUser, code string) (err error) {
	if err := DB.Model(u).Where("id = ?", u.Id).Updates(map[string]interface{}{"verify_code": code}).Error; err != nil {
		return err
	}
	return nil
}

func PutOneTeamUser(u *TeamUser) (err error) {
	DB.Save(u)
	return nil
}

func DeleteTeamUser(u *TeamUser, id int64, user_id int64) (err error) {
	DB.Where("id = ? AND user_id = ?", id, user_id).Delete(u)
	return nil
}
