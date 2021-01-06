package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{Name: cfg}

	//初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	//初始化日志包
	c.initLog()

	c.watchConfig()
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) //如果制定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("apiserver/demo4/config") //如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()            //读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") //读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { //解析viper配置文件
		return err
	}
	return nil
}
func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writes"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	log.InitWithConfig(&passLagerCfg)
}
