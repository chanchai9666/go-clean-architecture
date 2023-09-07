package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ProjectID   string `mapstructure:"PROJECT_ID"`
	SystemName  string `mapstructure:"SYSTEM_NAME"`
	WebBaseURL  string `mapstructure:"WEB_BASE_URL"`
	APIBaseURL  string `mapstructure:"API_BASE_URL"`
	Version     string `mapstructure:"VERSION"`
	Environment string `mapstructure:"ENVIRONMENT"`
	Release     bool   `mapstructure:"RELEASE"`
	Port        int    `mapstructure:"PORT"`
}

type APIConfig struct {
	HRIS struct {
		HostPIS   string `mapstructure:"HOSTPIS"`
		HostHRIS  string `mapstructure:"HOSTHRIS"`
		HostPSS   string `mapstructure:"HOSTPSS"`
		TokenAuth string `mapstructure:"TOKEN_AUTH"`
		Routes    struct {
			Bank              string `mapstructure:"BANK"`
			Position          string `mapstructure:"POSITION"`
			Employee          string `mapstructure:"EMPLOYEE"`
			EmployeeIDs       string `mapstructure:"EMPLOYEEIDS"`
			EmployeeIDCompany string `mapstructure:"EMPLOYEEIDCOMPANY"`
			Rewards           string `mapstructure:"REWARDS"`
			GroupRewards      string `mapstructure:"GROUPREWARDS"`
			CancelPatment     string `mapstructure:"CANCELPATMENT"`
			PathDepartment    string `mapstructure:"PATHDEPARTMENT"`
			PathPosition      string `mapstructure:"PATHPOSITION"`
			PathCompany       string `mapstructure:"PATHCOMPANY"`
			PathSection       string `mapstructure:"PATHSECTION"`
			PathSegments      string `mapstructure:"PATHSEGMENTS"`
		} `mapstructure:"ROUTES"`
	} `mapstructure:"HRIS"`
	FAS struct {
		Host      string `mapstructure:"HOST"`
		TokenAuth string `mapstructure:"TOKEN_AUTH"`
		Routes    struct {
			Bank     string `mapstructure:"BANK"`
			BankStop string `mapstructure:"BANKSTOP"`
		} `mapstructure:"ROUTES"`
	} `mapstructure:"FAS"`
}

type DatabaseConfig struct {
	Main struct {
		Host         string `mapstructure:"HOST"`
		Port         int    `mapstructure:"PORT"`
		Username     string `mapstructure:"USERNAME"`
		Password     string `mapstructure:"PASSWORD"`
		DatabaseName string `mapstructure:"DATABASE_NAME"`
		DriverName   string `mapstructure:"DRIVER_NAME"`
	} `mapstructure:"Main"`
	Main2 struct {
		Host         string `mapstructure:"HOST"`
		Port         int    `mapstructure:"PORT"`
		Username     string `mapstructure:"USERNAME"`
		Password     string `mapstructure:"PASSWORD"`
		DatabaseName string `mapstructure:"DATABASE_NAME"`
		DriverName   string `mapstructure:"DRIVER_NAME"`
	} `mapstructure:"Main2"`
}

type Config struct {
	AppConfig      AppConfig      `mapstructure:"APP"`
	APIConfig      APIConfig      `mapstructure:"API"`
	DatabaseConfig DatabaseConfig `mapstructure:"DATABASE"`
}

var CF *Config

func LoadConfig() (*Config, error) {
	if CF != nil {
		return CF, nil
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	CF = &config
	return CF, nil
}

// สร้างฟังก์ชันควบคุมในการโหลด Config
// เพื่อให้คุณสามารถเรียกใช้ LoadConfig ได้จากทุกที่ในโปรแกรม
func GetConfig() (*Config, error) {
	return LoadConfig()
}
