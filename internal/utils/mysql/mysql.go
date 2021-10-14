package mysql

import (
	"github.com/XC-Zero/common/client"
	"github.com/XC-Zero/common/internal/config"
	"gorm.io/gorm"
	"log"
)

var BatchSize = 20000

var UserDatapackDB *gorm.DB

func InitUserDatapackDB() {
	if UserDatapackDB != nil {
		return
	}
	db, err := client.InitGormV2(config.CONFIG.UserDatapackConfig)
	if err != nil {
		panic(err)
	}
	log.Println("InitDealDatapackDB success")
	UserDatapackDB = db
}
