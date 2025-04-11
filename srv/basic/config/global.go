package config

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	GlobalConfig Config
	DB           *gorm.DB
	RDB          *redis.Client
)
