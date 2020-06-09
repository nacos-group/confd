package main

import (
	"reflect"
	"testing"

	"github.com/kelseyhightower/confd/log"
	"github.com/satori/go.uuid"
)

func TestInitConfigDefaultConfig(t *testing.T) {
	log.SetLevel("warn")
	want := Config{
		BackendsConfig: BackendsConfig{
			Backend:      "etcd",
			BackendNodes: []string{"http://127.0.0.1:4001"},
			Scheme:       "http",
			Filter:       "*",
		},
		TemplateConfig: TemplateConfig{
			ConfDir:     "/etc/confd",
			ConfigDir:   "/etc/confd/conf.d",
			TemplateDir: "/etc/confd/templates",
			Noop:        false,
		},
		ConfigFile: "/etc/confd/confd.toml",
		Interval:   600,
	}
	if err := initConfig(); err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(want, config) {
		t.Errorf("initConfig() = %v, want %v", config, want)
	}
	u2, err1 := uuid.NewV4()
	if err1 != nil {
		println(`生成一个UUID v4时出现错误：`)
		panic(err1)
	}
	println(`生成的UUID v4：`)
	println(u2.String())
}
