package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DefaultDB *gorm.DB

func InitGorm(dsn string) error {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("[Warn] gorm init from dsl[%s] errr[%s]", dsn, err)
		return err
	}

	DefaultDB = db
	fmt.Printf("[Info] gorm init from dsl[%s] success", dsn)
	return nil
}
