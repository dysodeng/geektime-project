package db

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	initDB()
}

// 初始化数据库连接
func initDB() {
	var err error
	var dsn string

	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbUser := os.Getenv("db_user")
	dbPasswd := os.Getenv("db_passwd")
	dbDatabase := os.Getenv("db_database")

	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPasswd,
		dbHost,
		dbPort,
		dbDatabase,
	) + "?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"

	// db日志
	logFilename := "var/db.log"
	dbLogFile, _ := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	dbLogger := logger.New(
		log.New(io.MultiWriter(os.Stdout, dbLogFile), "", log.LstdFlags),
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, // 慢查询时间
			LogLevel:      logger.Warn,
			Colorful:      false,
		},
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true, // 禁用默认事务
		PrepareStmt:                              true, // 预编译sql
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁止表名复数
		},
		Logger: dbLogger, // db日志
	})
	if err != nil {
		panic("failed to connect database " + err.Error())
	}

	sqlDB, _ := db.DB()

	// 连接池
	sqlDB.SetMaxIdleConns(50)                   // 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于该值，超过的连接会被连接池关闭。
	sqlDB.SetMaxOpenConns(200)                  // 连接池最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 300) // 连接空闲超时
}

// DB 获取数据库连接
func DB() *gorm.DB {
	return db
}
