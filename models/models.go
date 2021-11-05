package Models

import (
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ModelConfig struct {
	Debug      bool
	Timeout    int
	SaaSDomain string
	Mysql      string
}

var DefaultConfig ModelConfig

var DB *gorm.DB

func Init(config *ModelConfig) error {

	//todo  check config if saas reflect
	if config.SaaSDomain != "" {

		var _ = reflect.TypeOf(Contact{})
		_ = reflect.TypeOf(ContactLabel{})
		_ = reflect.TypeOf(ContactLabelRelation{})
		_ = reflect.TypeOf(Domain{})
		_ = reflect.TypeOf(DomainUser{})
		_ = reflect.TypeOf(KanbanCard{})
		_ = reflect.TypeOf(KanbanList{})
		_ = reflect.TypeOf(KanbanProject{})
		_ = reflect.TypeOf(Label{})
		_ = reflect.TypeOf(LabelRelation{})
		_ = reflect.TypeOf(Mail{})
		_ = reflect.TypeOf(Product{})
		_ = reflect.TypeOf(Rule{})
		_ = reflect.TypeOf(Server{})
		_ = reflect.TypeOf(User{})
		_ = reflect.TypeOf(Session{})
		_ = reflect.TypeOf(AppUser{})
		_ = reflect.TypeOf(License{})
	}

	err := connect(config.Mysql)
	if err != nil {
		return err
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&DomainUser{})
	DB.AutoMigrate(&AppUser{})
	DB.AutoMigrate(&Rule{})
	DB.AutoMigrate(&Domain{})
	DB.AutoMigrate(&Server{})
	DB.AutoMigrate(&Mail{})
	DB.AutoMigrate(&Label{})
	DB.AutoMigrate(&LabelRelation{})
	DB.AutoMigrate(&KanbanProject{})
	DB.AutoMigrate(&KanbanList{})
	DB.AutoMigrate(&KanbanCard{})
	DB.AutoMigrate(&Contact{})
	DB.AutoMigrate(&ContactLabel{})
	DB.AutoMigrate(&ContactLabelRelation{})

	if config.SaaSDomain != "" {
		DB.AutoMigrate(&Product{})
		DB.AutoMigrate(&License{})
	}

	return err
}

func connect(host string) error {

	var err error
	if DB == nil {
		DB, err = gorm.Open(mysql.Open(host), &gorm.Config{})
		if err != nil {
			return err
		}
		//DB.LogMode(true)
		db, err := DB.DB()
		db.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		db.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		db.SetConnMaxLifetime(time.Hour * 87600)
		return err
	}

	return nil
}

func Close() {
	if DB != nil {
	}
}
