package network_cache

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"path"
	"pgo/pikapi/const_value"
	"time"
)

var db *gorm.DB

type NetworkCache struct {
	gorm.Model
	K string `gorm:"index:uk_k,unique"`
	V string
}

func InitDBConnect(databaseDir string) {
	var err error
	db, err = gorm.Open(sqlite.Open(path.Join(databaseDir, "network_cache.db")), const_value.GormConfig)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&NetworkCache{})
}

func LoadCache(key string, expire time.Duration) string {
	var cache NetworkCache
	err := db.First(&cache, "k = ? AND updated_at > ?", key, time.Now().Add(expire*-1)).Error
	if err == nil {
		return cache.V
	}
	if gorm.ErrRecordNotFound == err {
		return ""
	}
	panic(errors.New("?"))
}

func SaveCache(key string, value string) {
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "k"}},
		DoUpdates: clause.AssignmentColumns([]string{"created_at", "updated_at", "v"}),
	}).Create(&NetworkCache{
		K: key,
		V: value,
	})
}
