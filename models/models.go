package RestdocModels

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"

	//"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	// _ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"gorm.io/gorm/schema"
)

type ModelConfig struct {
	Debug  bool
	Mock   bool
	IsSaaS bool
	SqlDB  string
}

var DefaultConfig *ModelConfig

var DB *gorm.DB

func Init(config *ModelConfig) error {

	//todo  check config if saas reflect
	DefaultConfig = config

	err := connect()
	if err != nil {
		return err
	}

	return nil

}

func CreateTables() {

	refl()

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&VerifyCode{})
	DB.AutoMigrate(&RestAPI{})
	DB.AutoMigrate(&RestParam{})
	DB.AutoMigrate(&RestProject{})
	DB.AutoMigrate(&RestGroup{})
	DB.AutoMigrate(&RestAPI{})
	DB.AutoMigrate(&RestEndpoint{})
	DB.AutoMigrate(&Team{})
	DB.AutoMigrate(&TeamUser{})
	//add index labelRelation labelId

}

func refl() {
	var _ = reflect.TypeOf(User{})
	_ = reflect.TypeOf(Session{})
	_ = reflect.TypeOf(VerifyCode{})
	_ = reflect.TypeOf(RestAPI{})
	_ = reflect.TypeOf(RestParam{})
	_ = reflect.TypeOf(RestProject{})
	_ = reflect.TypeOf(RestGroup{})
	_ = reflect.TypeOf(RestAPI{})
	_ = reflect.TypeOf(RestEndpoint{})
	_ = reflect.TypeOf(Team{})
	_ = reflect.TypeOf(TeamUser{})
}

func connect() error {

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

			host := DefaultConfig.SqlDB
			if host != "" {
				if strings.HasPrefix(host, "mysql://") {
					host = strings.TrimPrefix(host, "mysql://")
					DB, err = gorm.Open(mysql.Open(host), &gorm.Config{})
				}

				if strings.HasPrefix(host, "postgresql://") {
					DB, err = gorm.Open(postgres.Open(host), &gorm.Config{})
				}

				if strings.HasPrefix(host, "cockroach://") {
					host = strings.Replace(host, "cockroach://", "postgresql://", -1)
					DB, err = gorm.Open(postgres.Open(host), &gorm.Config{})
				}

			}

			if err != nil {
				fmt.Println(err)
				return err
			}
		}

		//DB.LogMode(true)
		db, err := DB.DB()
		db.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		db.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		db.SetConnMaxLifetime(time.Minute * 20)
		return err
	}

	return nil
}

func Close() {
	if DB != nil {
	}
}
