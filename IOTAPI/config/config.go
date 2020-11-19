package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var defaultConf = []byte(`
app:
  port: 80
database:
  driver: sqlite,
  host: 127.0.0.1,
  port: 3306,
  username: test,
  password: test,
  database: iot,  
mqtt:
  host: 104.199.185.7
  port: 1883
  username: iot
  password:i iot 
`)

type ConfYaml struct {
	App  SectionApp      `yaml:"app"`
	DB   SectionDatabase `yaml:"db"`
	MQTT SectionMQTT     `yaml:"mqtt"`
}

type SectionApp struct {
	Port string `yaml:"port"`
}

type SectionDatabase struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type SectionMQTT struct {
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

func New(cfgPath string) *ConfYaml {
	config, err := LoadConf(cfgPath)
	if err != nil {
		log.Panic("Failed to load configuration file !")
	}
	return config
}

func LoadConf(confPath string) (*ConfYaml, error) {
	var conf ConfYaml

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()      // read in environment variables that match
	viper.SetEnvPrefix("IOT") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			return &conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return &conf, err
		}
	} else {
		viper.AddConfigPath("$HOME/.app")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			log.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			// load default config
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return &conf, err
			}
		}
	}

	//APP
	conf.App.Port = viper.GetString("app.port")

	//Database
	conf.DB.Driver = viper.GetString("database.driver")
	conf.DB.Host = viper.GetString("database.host")
	conf.DB.Port = viper.GetInt("database.port")
	conf.DB.Username = viper.GetString("database.username")
	conf.DB.Password = viper.GetString("database.password")
	conf.DB.Database = viper.GetString("database.database")

	//mqtt
	conf.MQTT.Host = viper.GetString("mqtt.host")
	conf.MQTT.Port = viper.GetInt("mqtt.port")
	conf.MQTT.Username = viper.GetString("mqtt.username")
	conf.MQTT.Password = viper.GetString("mqtt.password")

	return &conf, nil
}
