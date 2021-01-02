package env

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

type PostgreConfig struct {
	Host     string `envconfig:"POSTGRE_HOST" default:"localhost"`
	Port     int    `envconfig:"POSTGRE_PORT" default:"5432"`
	User     string `envconfig:"POSTGRE_USER" required:"true"`
	Password string `envconfig:"POSTGRE_PASSWORD" required:"true"`
	DBName   string `envconfig:"POSTGRE_DB" required:"true"`
	SSLMode  bool   `envconfig:"POSTGRE_SSL" default:"false"`
}

func (pc* PostgreConfig) DataSourceName() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pc.Host, pc.Port, pc.User, pc.Password, pc.DBName)
}

// Singleton implementation.
type Environment struct {
	PostgreConfig PostgreConfig
}

var (
	envConfig *Environment
	once      sync.Once
)

func Config() *Environment {
	once.Do(func() {
		err := envconfig.Process("", &envConfig)

		if err != nil {
			log.Fatal(err.Error())
		}
	})
	return envConfig
}
