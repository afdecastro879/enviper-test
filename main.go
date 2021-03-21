package main

import (
	"fmt"
	"github.com/iamolegga/enviper"
	"github.com/spf13/viper"
	"os"
)

type barry struct {
	Bar int `mapstructure:"bar"`
}
type bazzy struct {
	Baz bool
}
type quxxy struct {
	Qux string
}
type config struct {
	Foo     string
	Barry   barry
	Barries map[string]barry
	Bazzy   bazzy `mapstructure:",squash"`
	Quxxy   *quxxy
}

func main() {
	var c config

	e := enviper.New(viper.New())
	e.SetEnvPrefix("MYAPP")
	e.AddConfigPath(".")
	e.SetConfigName("config")
	setEnvs()

	e.Unmarshal(&c)
	fmt.Printf("%#v", c)
}

func setEnvs() {
	os.Setenv("FOO", "foo")
	os.Setenv("BARRY_BAR", "42")
	os.Setenv("BAZ", "true")
	os.Setenv("BARRIES_KEY1_BAR", "42")
	os.Setenv("QUXXY_QUX", "ipsum")
}
