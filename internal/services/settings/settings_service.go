package settings

import (
	"encoding/json"
	"fmt"
	"github.com/wmrodrigues/congress-mailer/internal/structs"
	"log"
	"os"
)

// LoadSettingsFile loads the settings file
func LoadSettingsFile() structs.Settings {
	settings := structs.Settings{}

	dir, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("error getting current path, %s", err.Error())
		log.Fatal(err)
	}

	file, err := os.Open(dir + "/configs/settings.json")
	if err != nil {
		err = fmt.Errorf("error loading settings.json file, %s", err.Error())
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&settings)
	if err != nil {
		err = fmt.Errorf("error decoding settings.json file to struct, %s", err.Error())
		log.Fatal(err)
	}

	return settings
}