package RestdocModels

import (
	_ "github.com/go-sql-driver/mysql"
)

type VerifyCode struct {
	Id         int64  `gorm:"type:bigint; not null; primary_key; " json: "id"`
	Email      string `gorm:"type:varchar(255);uniqueIndex:idx_email;not null" json:"email"`
	VerifyCode string `gorm:"type:varchar(128);not null;index:verify_code_index" json:"verify_code"`
	IP         int64  `gorm:"type:bigint ;not null;" json:"ip"`
	Type       int16  `gorm:"type:smallint;not null;" json:"type"`
	CreateAt   int64  `gorm:"type:bigint ;not null;" json:"create_at"`
	UpdateAt   int64  `gorm:"type:bigint ;not null;" json:"update_at"`
}

func (v *VerifyCode) TableName() string {
	return "verify_code"
}

func AddNewVerifyCode(v *VerifyCode) (err error) {
	if err = DB.Create(v).Error; err != nil {
		return err
	}
	return nil
}

func GetOneVerifyCode(v *VerifyCode, id string) (err error) {
	if err := DB.Where("id = ?", id).First(v).Error; err != nil {
		return err
	}
	return nil
}

func GetVerifyCodeByEmail(v *VerifyCode, email string) (err error) {
	if err := DB.Where("email = ?", email).First(v).Error; err != nil {
		return err
	}
	return nil
}

func GetVerifyCodeByEmailAndCode(v *VerifyCode, email string, code string) (err error) {
	if err := DB.Where("email = ? AND verify_code = ?", email, code).First(v).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVerifyCode(v *VerifyCode, email string, code string) (err error) {
	if err := DB.Model(v).Where("email = ?", email).Updates(map[string]interface{}{"verify_code": code}).Error; err != nil {
		return err
	}
	return nil
}

func PutOneVerifyCode(v *VerifyCode) (err error) {
	DB.Save(v)
	return nil
}

func DeleteVerifyCode(v *VerifyCode, id string) (err error) {
	DB.Where("id = ?", id).Delete(v)
	return nil
}
