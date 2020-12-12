package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"

var DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
