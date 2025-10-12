package config

import (
	"context"
	"fmt"
	"log"
	"sync"
	"travel/biz/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

// InitDatabase 初始化数据库连接
func InitDatabase(ctx context.Context) {
	once.Do(func() {
		// 从 viper 配置中读取数据库配置
		host := viper.GetString("mysql.host")
		port := viper.GetInt("mysql.port")
		database := viper.GetString("mysql.database")
		user := viper.GetString("mysql.user")
		password := viper.GetString("mysql.password")

		// 构建 DSN (Data Source Name)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, database)

		// 连接数据库
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
			return
		}

		// 获取底层的 sql.DB 对象来配置连接池
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("Failed to get underlying sql.DB: %v", err)
			return
		}

		// 设置连接池参数
		sqlDB.SetMaxIdleConns(10)  // 最大空闲连接数
		sqlDB.SetMaxOpenConns(100) // 最大打开连接数

		log.Printf("Successfully connected to MySQL database: %s", database)
		log.Printf("Executing Migration")
		DB.AutoMigrate(&model.Conversation{}, &model.ChatMemory{}, &model.Feedback{})
		log.Printf("Migration completed")
	})
}