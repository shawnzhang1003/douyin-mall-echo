package config

type MySQL struct {
	DSN      string `mapstructure:"dsn"`
	HOST     string `mapstructure:"host"`
	PORT     string `mapstructure:"port"`
	USER     string `mapstructure:"user"`
	PASSWORD string `mapstructure:"password"`
	DBNAME   string `mapstructure:"dbname"`
}
