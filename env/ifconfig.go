package env

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Dependency struct {
	Err    error
	DB     *sql.DB
	Params EnvironmentParameters
}

type (
	EnvironmentParameters struct {
		Ports struct {
			Gin string `yaml:"gin"`
		} `yaml:"ports"`
		Database struct {
			MySQLDB MySQLDB `yaml:"mysqldb"`
		} `yaml:"database_entity"`
		Schedular Schedular `yaml:"schedular"`
	}

	MySQLDB struct {
		DBURL             string `yaml:"database_url"`
		DBType            string `yaml:"database_type"`
		DBUser            string `yaml:"database_user"`
		DBPass            string `yaml:"database_password"`
		DBHost            string `yaml:"database_host"`
		DBHostWithoutPort string `yaml:"database_host_without_port"`
		DBPort            string `yaml:"database_port"`
		DBName            string `yaml:"database"`
		DBAdditional      struct {
			ParseTime string `yaml:"database_parse_time"`
		} `yaml:"database_additional"`
		DBConfig string `yaml:"db_config"`
	}

	Schedular struct {
		JobExecTime string `yaml:"job_exec_time"`
	}
)

func getENV(envName string) string {
	return os.Getenv(envName)
}

// NewENV : reading through provided config
func NewENV(configPath string) (*Dependency, error) {
	var settings Dependency
	config, err := os.Open(configPath)
	if err != nil {
		return NewENVFromMap()
	}

	defer func(config *os.File) {
		err = config.Close()
	}(config)
	d := yaml.NewDecoder(config)
	if err = d.Decode(&settings.Params); err != nil {
		return nil, err
	}
	log.Println(settings.Params)
	return &settings, err
}

func NewENVFromMap() (*Dependency, error) {
	var configs = "ports : \n" +
		"  gin : " + getENV("PORT_GIN") + "\n" +
		"database_entity : \n" +
		"  mysqldb : \n" +
		"   database_url : " + fmt.Sprintf("%q", getENV("DATABASE_URL")) + "\n" +
		"   database_type : " + getENV("DATABASE_TYPE") + "\n" +
		"   database_user : " + getENV("DATABASE_USER") + "\n" +
		"   database_password: " + getENV("DATABASE_PASSWORD") + "\n" +
		"   database_host: " + getENV("DATABASE_HOST") + "\n" +
		"   database_host_without_port: " + getENV("DATABASE_HOST_WITHOUT_PORT") + "\n" +
		"   database_port: " + getENV("DATABASE_PORT") + "\n" +
		"   database: " + getENV("DATABASE") + "\n" +
		"   database_additional : \n" +
		"      database_parse_time: " + getENV("DATABASE_PARSE_TIME") + "\n" +
		"schedular : \n" +
		"  job_exec_time : " + getENV("JOB_EXEC_TIME") + "\n"
	fmt.Println(configs)
	var settings Dependency
	d := yaml.NewDecoder(strings.NewReader(configs))
	err := d.Decode(&settings.Params)
	if err != nil {
		return nil, err
	}
	return &settings, err
}

// SetupMySQLDBConnection : parse database parameters
func (eP *EnvironmentParameters) SetupMySQLDBConnection() *EnvironmentParameters {
	dbInfo := eP.Database.MySQLDB
	eP.Database.MySQLDB.DBConfig = dbInfo.DBUser + ":" + dbInfo.DBPass + "@tcp(" + dbInfo.DBHostWithoutPort + ":" + dbInfo.DBPort + ")/" + dbInfo.DBName
	fmt.Println(eP.Database.MySQLDB.DBConfig)
	return eP
}
