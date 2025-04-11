package init

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"srv/basic/config"
)

func init() {
	InitConfig()
	InitMysql()
	InitRedis()
}

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("F:\\gocode\\src\\zhuangaosi\\rk\\day16\\srv\\appcattion.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&config.GlobalConfig)
	if err != nil {
		panic(err)
	}

}

func InitMysql() {
	con := config.GlobalConfig.Mysql
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", con.User, con.Password, con.Host, con.Port, con.Database)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
func InitRedis() {
	config.RDB = redis.NewClient(&redis.Options{
		Addr:     config.GlobalConfig.Redis.Addr,
		Password: config.GlobalConfig.Redis.Password, // no password set
		DB:       config.GlobalConfig.Redis.DB,       // use default DB
	})

}
