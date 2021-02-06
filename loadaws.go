package loadaws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Config - shape of the json obj stored on disk
type Config struct {
	AWS struct {
		SNS struct {
			ARN string
		}
	}
}

// FromENV - loads config from ~/.env.json
func FromENV() (Config, error) {
	homedir, ok := os.LookupEnv("HOME")
	if !ok {
		log.Fatal("Failed To Lookup $HOME")
	}
	pathToJSON := fmt.Sprintf("%s/%s", homedir, ".env.json")

	jsonString, err := ioutil.ReadFile(pathToJSON)
	if err != nil {
		log.Fatal("Failed To Load AWS JSON File: ", err)
	}

	var conf Config

	err = json.Unmarshal(jsonString, &conf)

	return conf, err
}
