package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DbAddr string `yarm: "dbaddr"`
	DbPort string `yarm: "dbport"`
	DbKind string `yarm: "dbkind"`
	DbVer  string `yarm: "dbver"`
}

func (cfg *Config) ReadConfig() error {
	bytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Printf("Config File Read Failed. err = '%v'\n", err.Error())
		return err
	}

	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		fmt.Printf("yaml.Unmarshal Failed. err = '%v'\n", err.Error())
		return err
	}

	return nil
}
