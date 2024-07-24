package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	// 方式1： 直接指定配置文件路径（相对路径或者绝对路径）
	// 相对路径：相对于可执行文件的路径
	// 绝对路径：系统中实际的文件路径
	//viper.SetConfigFile("./config.yaml") // 指定配置文件路径

	// 方式2：指定配置文件名和位置  viper自行查找可用的配置文件
	// 配置文件名不需要带后缀
	// 配置文件位置可配置多个
	viper.SetConfigName("config") // 指定配置文件名称
	viper.SetConfigType("yaml")   // 指定配置文件类型（专用于从远程获取配置信息时指定配置文件类型的）
	viper.AddConfigPath(".")      // 指定配置文件路径
	err = viper.ReadInConfig()    // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("ReadInConfig failed, err: %v\n", err)
		return
	}

	// 热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) { //回调函数
		fmt.Println("配置文件修改了...")
	})
	return
}
