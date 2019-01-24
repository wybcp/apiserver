package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

// Database 数据库
type Database struct {
	Self *gorm.DB
	// Docker *gorm.DB
}

// DB 数据库
var DB *Database

// Init 初始化连接
func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
		// Docker: GetDockerDB(),
	}
}

// GetSelfDB 获取self 数据库连接
func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

// InitSelfDB 初始化self数据库连接
func InitSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

// InitDockerDB 初始化docker数据库连接
func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

//GetDockerDB 获取docker数据库连接
func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}
func openDB(username, password, addr, name string) *gorm.DB {
	// gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	// set for db connection
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name is: %s", name)
	}
	log.Infof("%s Database was connected", name)
	setupDB(db)

	return db
}
func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

// Close 关闭数据库连接
func (db *Database) Close() {
	DB.Self.Close()
	// DB.Docker.Close()
}
