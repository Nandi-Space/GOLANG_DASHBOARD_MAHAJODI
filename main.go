package main

import (
	"Mahajodi_GOLANG_Dashboard/api"
	"Mahajodi_GOLANG_Dashboard/models"
	"Mahajodi_GOLANG_Dashboard/store"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

func main() {

	path := flag.String("c", "/etc/mahajodi", "config file location")
	writeToFile := flag.Bool("f", false, "write logs to file")
	flag.Parse()
	config := parseConfig(*path)
	level, err := logrus.ParseLevel(config.Logging.Level)
	if err != nil {
		level = logrus.ErrorLevel
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)

	if *writeToFile {
		f, err := os.OpenFile("/var/log/mahajodi_dashboard.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		logrus.SetOutput(f)
	}

	store.InitState(&config)
	api.StartServer(config.Server.Listen)

}

// parseConfig uses viper to parse config file.
func parseConfig(path string) models.Config {
	var config models.Config
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		panic("config file not found in " + filepath.Join(path))
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(absPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic("config file invalid: " + err.Error())
	}

	return config
}
