package config

type Kitex struct {
	Service         string `mapstructure:"service"`
	Address         string `mapstructure:"address"`
	MetricsPort     string `mapstructure:"metrics_port"`
	EnablePprof     bool   `mapstructure:"enable_pprof"`
	EnableGzip      bool   `mapstructure:"enable_gzip"`
	EnableAccessLog bool   `mapstructure:"enable_access_log"`
	LogLevel        string `mapstructure:"log_level"`
	LogFileName     string `mapstructure:"log_file_name"`
	LogMaxSize      int    `mapstructure:"log_max_size"`
	LogMaxBackups   int    `mapstructure:"log_max_backups"`
	LogMaxAge       int    `mapstructure:"log_max_age"`
}