package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const envPath = "./assets/"
const envType = "yaml"
const envPrefix = "demo"

var Env *viper.Viper

func init() {
	Env = initViper()
}

func initViper() *viper.Viper {
	v := viper.New()

	loadConfig(v, "env")
	if len(os.Args) >= 2 {
		mergeConfig(v, "env."+os.Args[1])
	}

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	fmt.Println("============ Config ============")
	for key, value := range v.AllSettings() {
		fmt.Println(key + ":" + fmt.Sprint(value))
	}
	fmt.Println("============ Config ============")

	return v
}

func loadConfig(v *viper.Viper, fileName string) {
	v.SetConfigName(fileName)
	v.SetConfigType(envType)
	v.AddConfigPath(envPath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("[Error] Loading config failed: ", err)
		panic(err)
	}
}

func mergeConfig(v *viper.Viper, fileName string) {
	v.SetConfigName(fileName)
	v.SetConfigType(envType)
	v.AddConfigPath(envPath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := v.MergeInConfig()
	if err != nil {
		fmt.Println("[Error] Merge config failed: ", err)
		panic(err)
	}
}
