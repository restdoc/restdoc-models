package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64  `gorm:"type:bigint; not null; primary_key; " json: "id"`
	Email    string `gorm:"type:varchar(255);uniqueIndex:idx_email;not null" json:"email"`
	Company  string `gorm:"type:varchar(20);not null;" json:"company"`
	Name     string `gorm:"type:varchar(50);not null;" json:"name"`
	Locale   string `gorm:"type:varchar(50);not null;" json:"locale"`
	Password string `gorm:"type:varchar(255);not null;" json:"password"`
	Valid    bool   `gorm:"type:boolean;not null;default false;" json:"valid"`
	CreateAt int64  `gorm:"type:bigint;not null;" json:"create_at"`
	UpdateAt int64  `gorm:"type:bigint;not null;" json:"update_at"`
}

func (u *User) TableName() string {
	return "user"
}

func GetAllUser(u *[]User) (err error) {
	if err = DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func UserCount() int64 {
	var count int64
	DB.Model(&User{}).Where("id > ?", 0).Count(&count)
	return count
}

func AddNewUser(u *User) (err error) {
	if err = DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(u *User, id string) (err error) {
	if err := DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(u *User, email string) (err error) {
	if err := DB.Where("email = ?", email).First(u).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByName(u *User, name string) (err error) {
	if err := DB.Where("name = ?", name).First(u).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePassword(u *User, email string, password string) (err error) {
	if err := DB.Model(u).Where("email = ?", email).Updates(map[string]interface{}{"password": password}).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUser(u *User) (err error) {
	DB.Save(u)
	return nil
}

func DeleteUser(u *User, id string) (err error) {
	DB.Where("id = ?", id).Delete(u)
	return nil
}
