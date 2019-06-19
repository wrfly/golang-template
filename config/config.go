package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/wrfly/ecp"
)

// Config structure
type Config struct {
	Port   int `env:"ENV_PORT" default:"1234"`
	Config struct {
		HHH string `default:"hhh"`
		JJJ string `default:"jjj"`
		KKK int    `default:"666"`
		LLL bool   `default:"true"`
	}
}

// New config from file
func New(cfgPath string) (*Config, error) {
	c := new(Config)

	bs, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, wrapE("read file error: %s", err)
	}
	if err := yaml.Unmarshal(bs, c); err != nil {
		return nil, wrapE("unmarshal config file error: %s", err)
	}

	if err := ecp.Parse(c, "ENV"); err != nil {
		return nil, wrapE("parse env config error: %s", err)
	}
	if err := ecp.Default(c); err != nil {
		return nil, wrapE("set default value error: %s", err)
	}
	return c, nil
}

func wrapE(format string, e error) error {
	return fmt.Errorf(format, e)
}

// Example print the config example
func Example() error {
	c := new(Config)
	if err := ecp.Default(c); err != nil {
		return err
	}
	bs, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	fmt.Printf("%s", bs)
	return nil
}

// EnvList print the config EnvList
func EnvList() {
	c := new(Config)
	for _, x := range ecp.List(c, "ENV") {
		fmt.Println(x)
	}
}
