package config

import (
	machineryConfig "github.com/RichardKnop/machinery/v2/config"
	"github.com/spf13/viper"
)

type Config struct {
	http      HttpConfig
	redis     RedisConfig
	machinery machineryConfig.Config
	printer   PrinterConfig
}

type HttpConfig struct {
	Port string
}

type RedisConfig struct {
	Host                   string
	Port                   string
	Password               string
	DB                     int
	Pool                   int
	MaxIdle                int
	IdleTimeout            int
	ReadTimeout            int
	WriteTimeout           int
	ConnectTimeout         int
	NormalTasksPollPeriod  int
	DelayedTasksPollPeriod int
}

type PrinterConfig struct {
	PrinterName []string
}

func (config *Config) GetMachineryConfig() *machineryConfig.Config {
	return &config.machinery
}

func (config *Config) GetRedisConfig() *RedisConfig {
	return &config.redis
}

func (config *Config) GetPrinterConfig() *PrinterConfig {
	return &config.printer
}

func (config *Config) GetHttpConfig() *HttpConfig {
	return &config.http
}

func ConfigInit() *Config {
	v := viper.New()
	v.SetConfigType("toml")
	v.AddConfigPath("./config")
	v.SetConfigName("setting")

	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}

	config := v.AllSettings()

	for k, v := range config {
		viper.SetDefault(k, v)
	}

	viper.SetConfigName("setting." + viper.GetString("env.mode"))
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	http := &HttpConfig{Port: viper.GetString("http.port")}

	redisConfig := &RedisConfig{
		Host:                   viper.GetString("redis.host"),
		Port:                   viper.GetString("redis.port"),
		Password:               viper.GetString("redis.password"),
		DB:                     viper.GetInt("redis.db"),
		Pool:                   viper.GetInt("redis.pool"),
		ReadTimeout:            viper.GetInt("redis.readTimeout"),
		WriteTimeout:           viper.GetInt("redis.writeTimeout"),
		ConnectTimeout:         viper.GetInt("redis.connectTimeout"),
		NormalTasksPollPeriod:  viper.GetInt("redis.normalTasksPollPeriod"),
		DelayedTasksPollPeriod: viper.GetInt("redis.delayedTasksPollPeriod"),
	}

	machineryConfig := &machineryConfig.Config{
		DefaultQueue:    viper.GetString("machinery.queue"),
		ResultsExpireIn: viper.GetInt("machinery.expired"),
		Redis: &machineryConfig.RedisConfig{
			MaxIdle:                redisConfig.MaxIdle,
			IdleTimeout:            redisConfig.IdleTimeout,
			ReadTimeout:            redisConfig.ReadTimeout,
			WriteTimeout:           redisConfig.WriteTimeout,
			ConnectTimeout:         redisConfig.ConnectTimeout,
			NormalTasksPollPeriod:  redisConfig.NormalTasksPollPeriod,
			DelayedTasksPollPeriod: redisConfig.DelayedTasksPollPeriod,
		},
	}

	printerConfig := &PrinterConfig{
		PrinterName: viper.GetStringSlice("printer.name"),
	}

	return &Config{
		http:      *http,
		redis:     *redisConfig,
		machinery: *machineryConfig,
		printer:   *printerConfig,
	}
}
