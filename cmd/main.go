package main

import (
	"fmt"
	"log"

	"eql/configs"
	"eql/internal/handlers/router"
	"eql/internal/infrastructures/gofiber"
)

func main() {
	// โหลด config จากไฟล์ config.yaml
	conf, err := configs.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	app := gofiber.NewServer() //start go fiber
	router.SetupRouter(app)    //เรียก router ต่างๆ

	// ผ่านตัวแปร conf ที่คุณได้โหลดมาแล้ว
	fmt.Println("ProjectID:", conf.AppConfig.ProjectID)
	fmt.Println("SystemName:", conf.AppConfig.SystemName)
	fmt.Println("WebBaseURL:", conf.AppConfig.WebBaseURL)
	fmt.Println("APIBaseURL:", conf.AppConfig.APIBaseURL)
	fmt.Println("Version:", conf.AppConfig.Version)
	fmt.Println("Environment:", conf.AppConfig.Environment)
	fmt.Println("Release:", conf.AppConfig.Release)
	fmt.Println("Port:", conf.AppConfig.Port)

	fmt.Println("HRIS HostPIS:", conf.APIConfig.HRIS.HostPIS)
	fmt.Println("HRIS HostHRIS:", conf.APIConfig.HRIS.HostHRIS)
	fmt.Println("HRIS HostPSS:", conf.APIConfig.HRIS.HostPSS)
	fmt.Println("HRIS TokenAuth:", conf.APIConfig.HRIS.TokenAuth)
	fmt.Println("HRIS Bank Route:", conf.APIConfig.HRIS.Routes.Bank)
	fmt.Println("HRIS Position Route:", conf.APIConfig.HRIS.Routes.Position)

	fmt.Println("FAS Host:", conf.APIConfig.FAS.Host)
	fmt.Println("FAS TokenAuth:", conf.APIConfig.FAS.TokenAuth)
	fmt.Println("FAS Bank Route:", conf.APIConfig.FAS.Routes.Bank)
	fmt.Println("FAS BankStop Route:", conf.APIConfig.FAS.Routes.BankStop)

	fmt.Println("Database Host:", conf.DatabaseConfig.Main.Host)
	fmt.Println("Database Port:", conf.DatabaseConfig.Main.Port)
	fmt.Println("Database Username:", conf.DatabaseConfig.Main.Username)
	fmt.Println("Database Password:", conf.DatabaseConfig.Main.Password)
	fmt.Println("Database DatabaseName:", conf.DatabaseConfig.Main.DatabaseName)
	fmt.Println("Database DriverName:", conf.DatabaseConfig.Main.DriverName)

	err = app.Listen(fmt.Sprintf(":%v", conf.AppConfig.Port))
	if err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}

}
