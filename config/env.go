package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const envType = "yaml"
const envPrefix = "demo"

var Env *viper.Viper

func InitConfig(path string) {
	Env = initViper(path)
}

func initViper(path string) *viper.Viper {
	v := viper.New()

	loadConfig(v, path, "env")
	if len(os.Args) >= 2 {
		mergeConfig(v, path, "env."+os.Args[1])
		v.Set("mode", os.Args[1])
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	fmt.Println("============ Config ============")
	for key, value := range v.AllSettings() {
		fmt.Println(key + ":" + fmt.Sprint(value))
	}
	fmt.Println("============ Config ============")

	return v
}

func loadConfig(v *viper.Viper, path string, fileName string) {
	v.SetConfigName(fileName)
	v.SetConfigType(envType)
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("[Error] Loading config failed: ", err)
		panic(err)
	}
}

func mergeConfig(v *viper.Viper, path string, fileName string) {
	v.SetConfigName(fileName)
	v.SetConfigType(envType)
	v.AddConfigPath(path)
	err := v.MergeInConfig()
	if err != nil {
		fmt.Println("[WARN] Merge config failed: ", err)
	}
}
