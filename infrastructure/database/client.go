package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ederj98/hex-movies-microservice/infrastructure/adapter/repository/entity"

	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/jinzhu/gorm"
)

const (
	MysqlUsername = "MYSQL_USERNAME"
	MysqlPassword = "MYSQL_PASSWORD"
	MysqlHost     = "MYSQL_HOST"
	MysqlSchema   = "MYSQL_SCHEMA"
	MysqlPort     = "MYSQL_PORT"
)

func GetDatabaseInstance() *gorm.DB {

	userName := os.Getenv(MysqlUsername)
	password := os.Getenv(MysqlPassword)
	host := os.Getenv(MysqlHost)
	schema := os.Getenv(MysqlSchema)
	port, _ := strconv.ParseInt(os.Getenv(MysqlPort), 10, 64)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC", userName, password, host, port, schema)
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		logger.Error(err.Error(), err)
		_ = db.Close()
		panic("database not working")
	}
	db.SingularTable(true)
	migrateDatabase(db)

	return db
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&entity.MovieEntity{})
}
