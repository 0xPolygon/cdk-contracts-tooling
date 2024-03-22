package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func LoadFromFile(configFilePath string, target interface{}) error {
	dirName, fileName := filepath.Split(configFilePath)

	fileExtension := strings.TrimPrefix(filepath.Ext(fileName), ".")
	fileNameWithoutExtension := strings.TrimSuffix(fileName, "."+fileExtension)

	viper.AddConfigPath(dirName)
	viper.SetConfigName(fileNameWithoutExtension)
	viper.SetConfigType(fileExtension)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("CDK_CONTRACTS")
	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			fmt.Println("config file not found")
		} else {
			fmt.Println("error reading config file: ", err)
			return err
		}
	}

	decodeHooks := []viper.DecoderConfigOption{
		// this allows arrays to be decoded from env var separated by ",", example: MY_VAR="value1,value2,value3"
		viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(mapstructure.TextUnmarshallerHookFunc(), mapstructure.StringToSliceHookFunc(","))),
	}

	return viper.Unmarshal(target, decodeHooks...)
}
