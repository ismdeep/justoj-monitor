package config

import (
	"github.com/BurntSushi/toml"
	"github.com/ismdeep/args"
	"io/ioutil"
)

// configStruct 配置结构体
type configStruct struct {
	JustOJ struct {
		Site string `json:"site"`
	} `json:"justoj"`
	Notification struct {
		Host  string   `json:"host"`
		Token string   `json:"token"`
		To    []string `json:"to"`
	} `json:"notification"`
	PendingMonitor struct {
		Duration         string `json:"duration"`
		ErrDuration      string `json:"errDuration"`
		PendingThreshold uint   `json:"pendingThreshold"`
	} `json:"pendingMonitor"`
}

var Global = &configStruct{}
var JustOJ = &Global.JustOJ
var Notification = &Global.Notification
var PendingMonitor = &Global.PendingMonitor

func LoadConfig() {
	filePath := "./server.toml"
	if args.Exists("-c") {
		filePath = args.GetValue("-c")
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	if err := toml.Unmarshal(content, Global); err != nil {
		panic(err)
	}
}
