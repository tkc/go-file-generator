package yaml

import (
	"fmt"
	"github.com/spf13/viper"
)

func getViper(confName string) *viper.Viper {
	vp := viper.New()
	vp.SetConfigName(confName)
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./yamls")
	vp.AddConfigPath("../yamls")
	vp.AddConfigPath("../../yamls")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error: %s", err))
	}
	return vp
}

func GetAccount(confName string) string {
	return getViper(confName).Get("account").(string)
}
