package model

import (
	"github.com/UniqueStudio/UniqueSSO/database"

	"github.com/sirupsen/logrus"
)

func InitTables() (err error) {
	err = database.DB.AutoMigrate(&User{})
	if err != nil {
		logrus.WithField("table", (&User{}).TableName()).Error("create table failed")
		return err
	}
	return nil
}
