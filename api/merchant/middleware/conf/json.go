package conf

import (
	"encoding/json"
	"github.com/go-chassis/openlog"
	"io/ioutil"
	"os"
)

type Config struct {
	Mysql struct {
		DSN   string `json:"dsn"`
		Debug bool   `json:debug`
	} `json:"mysql"`

	Redis struct {
		DSN string `json:"dsn"`
	} `json:"redis"`
	Cookie struct {
		Name       string `json:"name"`
		Prefix     string `json:"prefix"`
		EncryptKey string `json:"encrypt_key"`
	}
	Cache struct {
		CookiePrefix string `json:"cookie_prefix"`
	} `json:"cache"`
}
type Conf struct {
	Config *Config
}

var defaultJsonConfPath string = "conf/config.json"

func LoadConfig() *Conf {
	confFile, err := os.Open(defaultJsonConfPath)
	if err != nil {
		openlog.Fatal("open Json Config failed. " + err.Error())
		return nil
	}
	confContent, err := ioutil.ReadAll(confFile)
	if err != nil {
		openlog.Fatal("read Json Config failed. " + err.Error())
		return nil
	}
	config := &Config{}
	err = json.Unmarshal([]byte(confContent), &config)
	if err != nil {
		openlog.Fatal("unmarshal Json Config failed. " + err.Error())
		return nil
	}
	return &Conf{
		Config: config,
	}
}
