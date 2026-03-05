package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(config.ServerConf.Host)
	fmt.Println(config.ServerConf.Port)
}
