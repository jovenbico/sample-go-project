package config

import (
	"log"
	"testing"
)

func TestConfig(t *testing.T) {
	log.Println(">>> ", Config.MySQLDBName)
	log.Println(">>> ", Config.MySQLUser)
	log.Println(">>> ", Config.MySQLPass)
}
