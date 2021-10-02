package conf

import (
	"bytes"
	"errors"
	"log"
	"os"

	"gin-i18n/pkg/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	// SupportedConfigTypes --
	SupportedConfigTypes = []string{"json", "yaml", "toml"}
)

// ReadConfig read the config file regard to the file type (file extension)
// Also, viper can watch changes on the file --> allow us to hot reload the application
func ReadConfig(fileName string, configPaths ...string) bool {
	viper.SetConfigName(fileName)
	if len(configPaths) < 1 {
		// look for current dir
		viper.AddConfigPath(".")
	} else {
		for _, configPath := range configPaths {
			viper.AddConfigPath(configPath)
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Cannot read config file. %s\n", err)
		return false
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	return true
}

// ReadConfigByFile read config file by file path
func ReadConfigByFile(file string) bool {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Cannot read config file. %s\n", err)
		return false
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	return true
}

// LoadConfig reads configuration from environment variables.
func ReadConfigFromEnvVariables() bool {
	viper.AutomaticEnv()
	return true
}

// ReadFromBytes -- read config from bytes
// configType: "json", "toml", "yaml". If not set, will accept all kind of config types
func ReadFromBytes(configType string, value []byte, isMerge bool) (err error) {
	if configType != "" {
		if !utils.IsStringSliceContains(SupportedConfigTypes, configType) {
			return errors.New("Not supported config type " + configType)
		}
		viper.SetConfigType(configType)
	}

	if !isMerge {
		err = viper.ReadConfig(bytes.NewBuffer(value))
	} else {
		err = viper.MergeConfig(bytes.NewBuffer(value))
	}
	return
}

func MustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

// GetConfigLocation --
func GetConfigLocation() string {
	return MustGetenv("CONFIG_LOCATION")
}

// GetPodName -- get pod name in k8s
func GetPodName() string {
	return os.Getenv("HOSTNAME")
}

// GetLogOutput --
func GetLogOutput() string {
	return viper.GetString("LOGGER_OUTPUT")
}

// GetLogFortmat --
func GetLogFortmat() string {
	return viper.GetString("LOGGER_FORMAT")
}

// GetLogEnableDebug --
func GetLogEnableDebug() bool {
	return viper.GetBool("LOGGER_ENABLE_DEBUG")
}

// GetLogLevel --
func GetLogLevel() int {
	return viper.GetInt("LOGGER_LEVEL")
}

// GetHTTPPort --
func GetHTTPPort() int64 {
	return viper.GetInt64("PORT_HTTP")
}

// GetGrpcPort --
func GetGrpcPort() int64 {
	return viper.GetInt64("PORT_GRPC")
}
