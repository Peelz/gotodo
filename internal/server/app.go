package server

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start(ctx context.Context, conf map[string]string) error {
	dsn := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Bangkok"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
