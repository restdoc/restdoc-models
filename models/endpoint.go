package RestdocModels

import (
	"gorm.io/gorm/clause"
)

type RestEndpoint struct {
	Id        int64  `gorm:"type:bigint; not null; primary_key; " json: "id"`
	ProjectId int64  `gorm:"type:bigint; not null; index:project_id_index;" json: "project_id"`
	Name      string `gorm:"type:varchar(64);not null;" json:"name"`
	Value     string `gorm:"type:varchar(255) ;not null;" json:"value"`
	Weight    string `gorm:"type:varchar(32) ;not null;" json:"weight"`
	Status    int16  `gorm:"type:smallint ; not null; " json: "status"`
	CreatedAt int64  `gorm:"type:bigint;not null;" json:"created_at"`
	UpdatedAt int64  `gorm:"type:bigint;not null;" json:"updated_at"`
}

func (re *RestEndpoint) TableName() string {
	return "rest_endpoint"
}

func GetRestEndpointsByProjectId(endpoints *[]RestEndpoint, projectId int64) (err error) {
	selectFields := []string{"id", "project_id", "name", "value", "weight", "status", "created_at"}
	if err = DB.Where("project_id = ?", projectId).Find(endpoints).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func GetRestEndpointsByProjectIds(endpoints *[]RestEndpoint, projectIds []int64) (err error) {
	selectFields := []string{"id", "project_id", "name", "value", "weight", "status", "created_at"}
	if err = DB.Where("project_id in (?)", projectIds).Find(endpoints).Select(selectFields).Error; err != nil {
		return err
	}
	return nil
}

func AddNewRestEndpoint(re *RestEndpoint) error {
	if err := DB.Create(re).Error; err != nil {
		return err
	}
	return nil
}

func AddRestEndpoints(res *[]RestEndpoint) (err error) {
	DB.Clauses(clause.OnConflict{DoNothing: true}).Create(*res)
	return nil
}

func GetOneRestEndpoint(re *RestEndpoint, id int64) error {
	if err := DB.Where("id = ?", id).First(re).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRestEndpoint(re *RestEndpoint, updates map[string]interface{}) error {
	if err := DB.Model(re).Where("id = ?", re.Id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func PutOneRestEndpoint(re *RestEndpoint) (err error) {
	DB.Save(re)
	return nil
}

func DeleteRestEndpointsByProjectId(endpoints []RestEndpoint, id int64) error {
	if err := DB.Unscoped().Where("project_id = ?", id).Delete(&endpoints).Error; err != nil {
		return err
	}
	return nil

}

func DeleteRestEndpoint(re *RestEndpoint, id int64) (err error) {
	if err := DB.Unscoped().Where("id = ?", id).Delete(re).Error; err != nil {
		return err
	}
	return nil
}
