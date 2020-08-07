package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"goApi/global"
)

var ( // 命令行标志的定义
	env = flag.String("env", "prod", "env of the system")
)

// 在 main 函数之前被调用，根据调用关系决定执行的顺序
func init() {
	v := viper.New()
	flag.Parse()
	var configName string
	switch *env {
	case "prod":
		configName = "prod"
	case "test":
		configName = "test"
	case "devl":
		configName = "devl"
	default:
		configName = "prod"
	}

	configFile := "env/" + configName + ".yaml"
	fmt.Printf("watch the value from env :%s configFile :%s\n", *env, configFile)
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.SERVER_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.SERVER_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.VIPER_CONFIG = v
}
