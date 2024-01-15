package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config represents the configuration data
type Config struct {
	Name        string        `mapstructure:"name"`
	Environment string        `mapstructure:"environment"`
	Listen      string        `mapstructure:"listen"`
	Nats        string        `mapstructure:"nats"`
	Postgres    string        `mapstructure:"postgres"`
	Sendgrid    string        `mapstructure:"sendgrid_api_key"`
	SigninURL   string        `mapstructure:"signin_url"`
	Redis       redisConfig   `mapstructure:"redis"`
	Storage     storageConfig `mapstructure:"storage"`
}

type redisConfig struct {
	Address string `mapstructure:"address"`
	MaxConn int    `mapstructure:"max_conn"`
}

type storageConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKey       string `mapstructure:"access_key"`
	SecretKey       string `mapstructure:"secret_key"`
	SignatureExpiry uint   `mapstructure:"signature_expiry"`
	Bucket          string `mapstructure:"bucket"`
}

// Init conf with defaults
var _conf = Config{
	Listen: "0.0.0.0:80",
}

// Get returns the global config
func Get() Config {
	return _conf
}

// Read reads the global config from a json file
func Read(filePath string) {

	// Viper setup
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Bind config to env vars
	viper.BindEnv("environment", "ENV")
	viper.BindEnv("listen", "LISTEN")
	viper.BindEnv("postgres", "POSTGRES")
	viper.BindEnv("sendgrid_api_key", "SENDGRID_API_KEY")
	viper.BindEnv("signin_url", "SIGNIN_URL")
	viper.BindEnv("nats", "NATS")
	viper.BindEnv("redis.address", "REDIS_ADDRESS")
	viper.BindEnv("redis.maxconn", "REDIS_MAX_CONN")
	viper.BindEnv("storage.endpoint", "STORAGE_ENDPOINT")
	viper.BindEnv("storage.access_key", "STORAGE_ACCESS_KEY")
	viper.BindEnv("storage.secret_key", "STORAGE_SECRET_KEY")
	viper.BindEnv("storage.signature_expiry", "STORAGE_SIGNATURE_EXPIRY")
	viper.BindEnv("storage.bucket", "STORAGE_BUCKET")

	// Reads the config
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file\n", err)
	}

	err = viper.Unmarshal(&_conf)
	if err != nil {
		log.Fatal("Cound not unmarshall config\n", err)
	}
}
