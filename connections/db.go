package connections

import (
	"strconv"
	"time"

	"github.com/huprince/quick-gin/config"
	"github.com/huprince/quick-gin/modules/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// InitDB 初始化DB
func InitDB() (*gorm.DB, error) {
	dbType := config.GetEnv().DbConfig.DbType
	dbHost := config.GetEnv().DbConfig.DbHost
	dbPort := config.GetEnv().DbConfig.DbPort
	dbName := config.GetEnv().DbConfig.DbName
	dbUser := config.GetEnv().DbConfig.DbUser
	dbPassword := config.GetEnv().DbConfig.DbPassword
	var dsn string
	var db *gorm.DB
	var err error
	switch(dbType) {
	case "mysql":
		dsn = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		break
	case "postgres":
		dsn = "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + strconv.Itoa(dbPort) + " vsslmode=disable TimeZone=Asia/Shanghai"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		break
	case "sqlserver":
		dsn = "sqlserver://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + strconv.Itoa(dbPort) + "?database=" + dbName
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		break
	default:
		dbName := config.GetEnv().DbConfig.DbName
		db, err = gorm.Open(sqlite.Open(dbName + ".db"), &gorm.Config{})
		break
	}
	if dbType != "sqlite" {
		sqlDb, err := db.DB()
		if err != nil {
			log.Logger.Error(err.Error())
		}
		sqlDb.SetMaxIdleConns(config.GetEnv().DbConfig.DbMaxIdleConns)
		sqlDb.SetMaxOpenConns(config.GetEnv().DbConfig.DbMaxOpenConns)
		sqlDb.SetConnMaxLifetime(time.Duration(config.GetEnv().DbConfig.DbConnMaxLifetime) * time.Minute)
	}
	return db, err

}