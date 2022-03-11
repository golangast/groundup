package createserver

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	. "gitlab.com/zendrulat123/groundup/cmd/utserver"
)

func Createservers() {
	viper.SetConfigName("persis") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config/") // config file path
	viper.AutomaticEnv()           // read value ENV variable
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	p := viper.Get("home.path")
	f := viper.Get("home.file")
	s := viper.Get("home.script")
	sp := fmt.Sprintf("%v", p)
	sf := fmt.Sprintf("%v", f)
	ss := fmt.Sprintf("%v", s)

	CreateServer(sp, sf, ss)

}
