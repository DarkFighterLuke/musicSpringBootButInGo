package utils

import (
	"fmt"
	"io/ioutil"

	"github.com/dzwvip/oracle"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

type Credentials struct {
	Username string
	Password string
}

func GetDatabaseConnection() (*gorm.DB, error) {
	credentials, err := getDatabaseCredentials()
	if err != nil {
		return nil, err
	}

	return gorm.Open(oracle.Open(fmt.Sprintf("%s/%s@marte.etlforma.com:8524/xe", credentials.Username, credentials.Password)), &gorm.Config{})
}

func getDatabaseCredentials() (*Credentials, error) {
	var credentials Credentials
	file, err := ioutil.ReadFile("conf/conf.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &credentials)
	if err != nil {
		return nil, err
	}

	return &credentials, nil
}
