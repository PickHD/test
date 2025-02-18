package config

import "test/internal/helper"

type (
	Configuration struct {
		General  *GeneralConf
		Database *DatabaseConf
		Jwt      *JwtConf
	}

	GeneralConf struct {
		AppPort int
	}

	DatabaseConf struct {
		Host     string
		Port     int
		Username string
		Password string
		DbName   string
	}

	JwtConf struct {
		Secret string
		Expire int
	}
)

func NewConfig() *Configuration {
	return &Configuration{
		General: &GeneralConf{
			AppPort: helper.GetEnvInt("APP_PORT"),
		},
		Database: &DatabaseConf{
			Host:     helper.GetEnvString("DB_HOST"),
			Port:     helper.GetEnvInt("DB_PORT"),
			Username: helper.GetEnvString("DB_USERNAME"),
			Password: helper.GetEnvString("DB_PASSWORD"),
			DbName:   helper.GetEnvString("DB_NAME"),
		},
		Jwt: &JwtConf{
			Secret: helper.GetEnvString("JWT_SECRET"),
			Expire: helper.GetEnvInt("JWT_EXPIRE"),
		},
	}
}
