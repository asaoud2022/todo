package config

import (
	"github.com/spf13/viper"
)

type Config struct {

	// ServerConfig
	Port string `mapstructure:"PORT"`
	//Name        string `env:"APP_NAME" env-default:"fibo"`
	//Host        string `env:"APP_HOST" env-default:"localhost"`
	//ProxyHeader string `mapstructure:"PROXY_HEADER" env:"PROXY_HEADER" env-default:"*"`
	//UploadSize  int    `mapstructure:"UPLOAD_SIZE" env:"UPLOAD_SIZE" env-default:"400"`

	// DatabaseConfig
	DBHost string `mapstructure:"DB_HOST"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
	DBPort string `mapstructure:"DB_PORT"`
	//SSLMode     string `env:"DB_SSL_MODE" env-default:"disable"`

	// LogConfig.
	//FilePath string `env:"LOG_FILE_FORMAT" env-default:"./tmp/logs/%s-%s.log"`

	//# See MW Fiber Session config:
	SessionStorageProvider         string `mapstructure:"MW_FIBER_SESSION_STORAGE_PROVIDER"`
	SessionStorageHost             string `mapstructure:"MW_FIBER_SESSION_STORAGE_HOST"`
	SessionStoragePort             int    `mapstructure:"MW_FIBER_SESSION_STORAGE_PORT"`
	SessionStorageTlsEnabled       bool   `mapstructure:"MW_FIBER_SESSION_STORAGE_TLS_ENABLED"`
	SessionStorageUserName         string `mapstructure:"MW_FIBER_SESSION_STORAGE_USERNAME"`
	SessionStoragePassword         string `mapstructure:"MW_FIBER_SESSION_STORAGE_PASSWORD"`
	SessionStorageDatabase         int    `mapstructure:"MW_FIBER_SESSION_STORAGE_DATABASE"`
	SessionStorageTable            string `mapstructure:"MW_FIBER_SESSION_STORAGE_TABLE"`
	SessionSesionCookieName string `mapstructure:"MW_FIBER_SESSION_COOKIENAME"`
	SessionCookieDomain     string `mapstructure:"MW_FIBER_SESSION_COOKIEDOMAIN"`
	SessionCookiePath       string `mapstructure:"MW_FIBER_SESSION_COOKIEPATH"`
	SessionCookieHttpOnly   bool `mapstructure:"MW_FIBER_SESSION_COOKIEHTTPONLY"`
	// Options: "Lax", "Strict", "None"
	SessionStorageCookieSameSite string `mapstructure:"MW_FIBER_SESSION_COOKIESAMESITE"`
	SessionStorageReset          bool   `mapstructure:"MW_FIBER_SESSION_STORAGE_RESET"`
	SessionCookieSecure          bool `mapstructure:"MW_FIBER_SESSION_COOKIESECURE"`
	SessionExpiration            string `mapstructure:"MW_FIBER_SESSION_EXPIRATION"`
	SessionStorageGCInterval     string `mapstructure:"MW_FIBER_SESSION_STORAGE_GCINTERVAL"`

	
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
