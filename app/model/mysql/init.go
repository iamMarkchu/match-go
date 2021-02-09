package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"match-go/app/model/mysql/match"
)

var (
	DB  *gorm.DB
	err error

	MatchSQL *match.SQL
)

func Init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/match_base?charset=utf8mb4&parseTime=True&loc=Local"
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		panic("gorm.Open error")
	}

	// migration
	if err = DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&match.Model{}); err != nil {
		log.Println("migration err", err)
	}
	// 初始化model
	MatchSQL = match.NewMatchSQL(DB)
}
