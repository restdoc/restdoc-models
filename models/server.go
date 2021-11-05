package Models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	Id      uint64 `gorm:"type:bigint unsigned;not null;primary key;" json:"id"`
	Ip      string `gorm:"type:varchar(20);not null;" json:"ip"`
	Port    int32  `gorm:"type:int unsigned;not null;uniqueIndex:idx_rdns_port" json:"port"`
	Rdns    string `gorm:"type:varchar(100);not null;uniqueIndex:idx_rdns_port;" json:"rdns"`
	Prefix  string `gorm:"type:varchar(20);not null;" json:"prefix"`
	Type    uint8  `gorm:"type:tinyint(8);not null;" json:"type"`
	Foreign uint8  `gorm:"type:tinyint(8);not null;" json:"foreign"`
	State   string `gorm:"type:varchar(20);not null;" json:"state"`
}

func (s *Server) TableName() string {
	return "server"
}

func GetAllServer(s *[]Server) (err error) {
	if err = DB.Find(s).Error; err != nil {
		return err
	}
	return nil
}

func AddNewServer(s *Server) (err error) {
	if err = DB.Create(s).Error; err != nil {
		return err
	}
	return nil
}

func GetOneServer(s *Server, id uint64) (err error) {
	if err := DB.Where("id = ?", id).First(s).Error; err != nil {
		return err
	}
	return nil
}

func GetServerByIp(s *Server, ip string) (err error) {
	if err := DB.Where("ip = ?", ip).First(s).Error; err != nil {
		return err
	}
	return nil
}

func GetServerByRdns(s *Server, rdns string) (err error) {
	if err := DB.Where("rdns = ?", rdns).First(s).Error; err != nil {
		return err
	}
	return nil
}

func GetServersByType(servers *[]Server, _type uint8) (err error) {
	if err := DB.Where("type = ?", _type).Find(&servers).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func PutOneServer(s *Server) (err error) {
	DB.Save(s)
	return nil
}

func DeleteServer(s *Server, id uint64) (err error) {
	if err := DB.Unscoped().Where("id = ? ", id).Delete(s).Error; err != nil {
		return err
	}
	return nil
}
