package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	// Config main config
	Config config
	// ProjectPath project absolute path
	ProjectPath string
	// LoggerPath logger absolute path
	LoggerPath string
	configPath string
)

// Config struct
type config struct {
	MySQLDBName string
	MySQLUser   string
	MySQLPass   string
}

func init() {
	log.Println("init config")

	ProjectPath = os.Getenv("GO_PARKBY")
	if ProjectPath == "" {
		panic("Set env variable $GO_PARKBY")
	}

	sep := string(os.PathSeparator) // shorten os.PathSeparator

	configPath = ProjectPath + sep + "config"
	configPath += sep + "config.json"

	LoggerPath = ProjectPath + sep + "log" + sep + "parkby.log"

	initMySQL()
}

func initMySQL() {
	log.Println("init config MySQL")
	file, err := os.Open(configPath)
	printlnErr(err)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	printlnErr(err)
}

func printlnErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
