package xgorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

// config options
type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	Db              string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		Host:            "localhost",
		Port:            3306,
		User:            "root",
		Password:        "123456",
		Db:              "test",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxIdleTime: time.Second * 10,
		ConnMaxLifetime: time.Minute * 10,
	}
}

func (c *Config) Build() *gorm.DB {

	db, err := gorm.Open("mysql", BuildConnString(c.Host, c.Port, c.User, c.Password, c.Db))
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	db.DB().SetMaxIdleConns(c.MaxIdleConns)

	db.DB().SetConnMaxLifetime(c.ConnMaxLifetime)
	db.DB().SetConnMaxIdleTime(c.ConnMaxIdleTime)

	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	if os.Getenv("GORM_DEBUG") == "true" {
		db = db.Debug()
	}
	db.SingularTable(true)
	return db
}

// BuildConnString root:secret@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local
func BuildConnString(host string, port int, user, password, db string) string {
	return fmt.Sprintf(`%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`, user, password, host, port, db)
}
