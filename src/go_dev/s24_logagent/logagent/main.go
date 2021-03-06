package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

func getLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "trace":
		return logs.LevelTrace
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "error":
		return logs.LevelError
	default:
		return logs.LevelDebug
	}
}

var appConfig = &AppConfig{}

func main() {
	err := initConfig("./logagent.cfg")
	if err != nil {
		fmt.Println("new config error:",err)
		return
	}

	config := make(map[string]interface{})
	config["filename"] = appConfig.LogFile
	config["level"] = getLevel(appConfig.LogLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.SetLogFuncCall(true)  // 打印文件名、文件行号

	err = initKafka()
	if err != nil {
		logs.Error("init kafka error:%v", err)
		return
	}

	RunServer()
}