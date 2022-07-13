package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)
/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description 
 **/
func init()  {
	loadConfiguration()
}

var C= struct {
	Iris iris.Configuration
	Addr struct{
		Internal struct{
			IP     string
			Plain  int
			Secure int
		}
	}
	Locale struct {
		Pattern   string
		Default   string
		Supported []string
	}

}{
	Iris: iris.DefaultConfiguration(),
}
func loadConfiguration() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/app/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.app") // call multiple times to add many search paths
    viper.AddConfigPath(".")
   err:=	viper.ReadInConfig()
	if err:=viper.ReadInConfig();err!=nil{
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			panic(fmt.Errorf("load configuration: %w", err))
		}
	}
	err:=viper.Unmarshal(&C)
	if err != nil {
		panic(fmt.Errorf("load configuration: unmarshal: %w", err))
	}
}