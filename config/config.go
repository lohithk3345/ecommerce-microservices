package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var EnvMap = getENV()

func isEmpty(str string) bool {
	return str == ""
}

func getENV() map[string]string {
	envMap := make(map[string]string)
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
	envMap["MONGODB_URL"] = os.Getenv("MONGODB_URL")
	envMap["API_KEY"] = os.Getenv("API_KEY")
	// log.Println(envMap)
	for key, value := range envMap {
		if isEmpty(value) {
			log.Fatalf("ENV not found for %s", key)
		}
	}
	return envMap
}
