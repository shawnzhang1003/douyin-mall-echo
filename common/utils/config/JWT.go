package config

type JWTConfig struct {
	RefreshSecretKey string   `mapstructure:"refresh_secret_key"`
	SecretKey        string   `mapstructure:"secret_key"`
	Whitelist        []string `mapstructure:"whitelist"`
}
