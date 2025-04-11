package config

type Config struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql"`
	Redis Redis `mapstructure:"redis" json:"redis"`
}

type Mysql struct {
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Database string `mapstructure:"database" json:"database"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"   json:"addr"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"DB"`
}
