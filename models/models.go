package Models

import (
	"reflect"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ModelConfig struct {
	Debug  bool
	Mock   bool
	IsSaaS bool
	Mysql  string
}

var DefaultConfig *ModelConfig

var DB *gorm.DB

func Init(config *ModelConfig) error {

	//todo  check config if saas reflect
	DefaultConfig = config

	err := connect(DefaultConfig.Mysql)
	if err != nil {
		return err
	}

	return nil
}

func CreateTables() {

	var _ = reflect.TypeOf(AppUser{})
	_ = reflect.TypeOf(Contact{})
	_ = reflect.TypeOf(ContactLabel{})
	_ = reflect.TypeOf(ContactLabelRelation{})
	_ = reflect.TypeOf(Domain{})
	_ = reflect.TypeOf(DomainUser{})
	_ = reflect.TypeOf(KanbanCard{})
	_ = reflect.TypeOf(KanbanList{})
	_ = reflect.TypeOf(KanbanProject{})
	_ = reflect.TypeOf(Label{})
	_ = reflect.TypeOf(LabelRelation{})
	_ = reflect.TypeOf(License{})
	_ = reflect.TypeOf(Mail{})
	_ = reflect.TypeOf(MailLog{})
	_ = reflect.TypeOf(Product{})
	_ = reflect.TypeOf(Rule{})
	_ = reflect.TypeOf(Server{})
	_ = reflect.TypeOf(Session{})
	_ = reflect.TypeOf(Summary{})
	_ = reflect.TypeOf(User{})
	_ = reflect.TypeOf(Unsub{})
	_ = reflect.TypeOf(RestAPI{})
	_ = reflect.TypeOf(RestParam{})
	_ = reflect.TypeOf(RestProject{})

	DB.AutoMigrate(&AppUser{})
	DB.AutoMigrate(&Contact{})
	DB.AutoMigrate(&ContactLabel{})
	DB.AutoMigrate(&ContactLabelRelation{})
	DB.AutoMigrate(&Domain{})
	DB.AutoMigrate(&DomainUser{})
	DB.AutoMigrate(&KanbanProject{})
	DB.AutoMigrate(&KanbanList{})
	DB.AutoMigrate(&KanbanCard{})
	DB.AutoMigrate(&Label{})
	DB.AutoMigrate(&LabelRelation{})
	DB.AutoMigrate(&Mail{})
	DB.AutoMigrate(&MailLog{})
	DB.AutoMigrate(&Rule{})
	DB.AutoMigrate(&Server{})
	DB.AutoMigrate(&Summary{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Unsub{})
	DB.AutoMigrate(&RestAPI{})
	DB.AutoMigrate(&RestParam{})
	DB.AutoMigrate(&RestProject{})
	//add index labelRelation labelId

	if DefaultConfig.IsSaaS {
		DB.AutoMigrate(&License{})
		DB.AutoMigrate(&Product{})
	}
}

func connect(host string) error {

	var err error
	if DB == nil {

		if DefaultConfig.Mock {
			sqlDB, _, err := sqlmock.New()
			if err != nil {
				return err
			}
			DB, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB, SkipInitializeWithVersion: true}), // auto configure based on currently MySQL version
				&gorm.Config{})
		} else {
			DB, err = gorm.Open(mysql.Open(host), &gorm.Config{})
			if err != nil {
				return err
			}

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
