package global

import "time"

type ServerConfig struct {
	ServerInfo BaseConfig  `mapstructure:"server"`
	MysqlInfo  MysqlConfig `mapstructure:"mysql"`
	RedisInfo  RedisConfig `mapstructure:"redis"`
	LogInfo    LogConfig   `mapstructure:"logInfo"`
}

type BaseConfig struct {
	Name                 string        `mapstructure:"name"`
	Port                 int           `mapstructure:"port"`
	LogPath              string        `mapstructure:"logPath"`
	WriteTimeout         time.Duration `mapstructure:"writeTimeout"`
	ReadTimeout          time.Duration `mapstructure:"readTimeout"`
	UploadSavePath       string        `mapstructure:"UploadSavePath"`
	UploadServerUrl      string        `mapstructure:"UploadServerUrl"`
	UploadImageMaxSize   int           `mapstructure:"UploadImageMaxSize"`
	UploadImageAllowExts []string      `mapstructure:"UploadImageAllowExts"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"Password"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups"`
}
