package Models

import (
	"errors"
	"sort"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	NormalProduct  = uint8(1)
	ExpireProduct  = uint8(5)
	DeletedProduct = uint8(10)
)

type Product struct {
	ID          uint64 `gorm:"type:bigint unsigned; not null; primary_key; " json: "id"`
	UserId      uint64 `gorm:"type:bigint unsigned;not null;" json:"user_id"`
	OrderId     string `gorm:"type:varchar(64);not null;uniqueIndex:idx_order_id;" json:"order_id"`
	AccountId   string `gorm:"type:varchar(64);not null" json:"account_id"`
	OpenId      string `gorm:"type:varchar(32);not null" json:"open_id"`
	ProductId   uint64 `gorm:"type:bigint unsigned;not null" json:"product_id"`
	RequestId   string `gorm:"type:varchar(64);not null" json:"request_id"`
	SignId      string `gorm:"type:varchar(12);not null;" json:"sign_id"`
	ProductName string `gorm:"type:varchar(128);not null" json:"product_name"`
	IsTrial     bool   `gorm:"type:boolean;not null; default false" json:"is_trial"`
	Spec        string `gorm:"type:varchar(64);not null" json:"spec"`
	Timespan    uint64 `gorm:"type:bigint unsigned;not null;" json:"time_span"`
	Timeunit    string `gorm:"type:varchar(32);not null" json:"time_unit"`
	Comment     string `gorm:"type:varchar(128);not null" json:"comment"`
	TemplateId  string `gorm:"type:varchar(32);not null" json:"template_id"`
	Ccnid       string `gorm:"type:varchar(32);not null" json:"ccn_id"`
	ExpireTime  string `gorm:"type:varchar(32);not null" json:"expire_time"`
	Type        uint8  `gorm:"type:tinyint(8);not null;" json:"type"`
	CreateAt    int64  `gorm:"type:bigint unsigned;not null;" json:"create_at"`
	UpdateAt    int64  `gorm:"type:bigint unsigned;not null;" json:"update_at"`
}

func (p *Product) TableName() string {
	return "product"
}

func AddNewProduct(p *Product) (err error) {

	if err = DB.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func GetOneProduct(p *Product, id uint64, user_id uint64) (err error) {
	if err := DB.Where("id = ? AND user_id = ?", id, user_id).First(p).Error; err != nil {
		return err
	}
	return nil
}

func GetOneProductByOrderId(p *Product, order_id string, user_id uint64) (err error) {
	if err := DB.Where("order_id = ? ", order_id).First(p).Error; err != nil {
		return err
	}
	return nil
}

func PutOneProduct(p *Product) (err error) {
	DB.Save(p)
	return nil
}

func UpdateExpireTime(p *Product, expire string) error {
	ts := time.Now().Unix()
	updates := map[string]interface{}{
		"expire_time": expire,
		"update_at":   ts,
	}
	if err := DB.Model(p).Where("order_id = ?", p.OrderId).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func UpdateType(p *Product, ptype uint8) error {
	ts := time.Now().Unix()
	updates := map[string]interface{}{
		"type":      ptype,
		"update_at": ts,
	}
	if err := DB.Model(p).Where("order_id = ?", p.OrderId).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(products *[]Product, user_id uint64) (err error) {

	if len(*products) == 0 {
		return errors.New("wrong length")
	}

	ids := []uint64{}
	for i := range *products {
		m := (*products)[i]
		id := m.ID
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	m := Mail{UserId: user_id}

	if err := DB.Unscoped().Where("user_id = ? AND id in (?)", user_id, ids).Delete(m).Error; err != nil {
		return err
	}
	return nil
}
