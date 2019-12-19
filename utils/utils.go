package utils

import "log"

import "github.com/spf13/viper"

import "encoding/json"

import "os"

type config struct {
	DatabaseHost string `mapstructure:"database_host"`
	DatabasePort int    `mapstructure:"database_port"`
}

var (
	// Config holds overall Database Config
	Config config
)

// InitConfig setup Configuration
func InitConfig() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	viper.SetConfigName("config") //name of the config file without extension
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(&Config)

	log.Printf("\n\n CONFIGURATION\n")
	log.Printf("\n======================================================================\n")
	displayConfig := Config
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	_ = enc.Encode(displayConfig)
	log.Printf("\n======================================================================\n")
}
