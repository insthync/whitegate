package configurations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// AppConfiguration ... This is configuration for whole app
type (
	AppConfiguration struct {
		Port             int    `json:"port"`
		DatabaseHost     string `json:"dbHost"`
		DatabaseUsername string `json:"dbUser"`
		DatabasePassword string `json:"dbPass"`
		DatabaseName     string `json:"dbName"`
		UserPasswordSalt string `json:"userPasswordSalt"`
	}
)

var appConfig *AppConfiguration

// ReadConfig ... Function to read config from conf.json
func ReadConfig() (*AppConfiguration, error) {

	fmt.Println("Reading configuration file...")
	raw, err := ioutil.ReadFile("./conf.json")

	if err != nil {
		fmt.Println("Failed to read configuration file.")
		fmt.Println(err.Error())
		return appConfig, err
	}

	err = json.Unmarshal(raw, &appConfig)
	if err != nil {
		fmt.Println("Failed to read configuration file.")
		fmt.Println(err.Error())
		return appConfig, err
	}
	fmt.Println("Configuration file was read successfully.")

	return appConfig, nil
}

// GetAppConfiguration ... Get main configuration data
func GetAppConfiguration() *AppConfiguration {
	return appConfig
}
