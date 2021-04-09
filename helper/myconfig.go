package helper

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppServer        string
	DbDriver         string
	HostDatabase     string
	UserDatabase     string
	PasswordDatabase string
	DatabaseName     string
	EncryptDatabase  string
}

func LoadConfiguration() (Config, error) {
	var config Config
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

/*
func Find() myConnectionStrings {
	var err error
	var cnf = flag.String("cnf", "config.json", "Database Config")
	flag.Parse()
	file, err := os.Open(*cnf)
	if err != nil {
		log.Fatal("can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := myConnectionStrings{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}

	constrings := myConnectionStrings{
		Host:     Config.Host,
		User:     Config.User,
		Password: Config.Password,
		Database: Config.Database,
		Encrypt:  Config.Encrypt,
	}
	return constrings
}
*/
