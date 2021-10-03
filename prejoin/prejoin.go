package prejoin

import (
	"gorm-test/config"
)

func Query() {
	db := config.Db().Debug()
	user := User{}
	db.Preload("Address", "IsActive = ?", "1").Find(&user)
}
