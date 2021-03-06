package clogger

import (
	"fmt"
	"runtime"
	"strings"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/goccy/go-json"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var purple = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
		purple = ""
		cyan = ""
		gray = ""
		white = ""
	}
}

func Default(message *map[string]interface{}) {
	logger.Log.Level = "DEFAULT"
	logger.Log.Message = message
	pc, file, line, ok := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	fn := funcName[strings.LastIndex(funcName, ".")+1:]

	if ok {
		logger.Log.Line = file + ":" + fmt.Sprintf("%v", line)
		logger.Log.Method = fn
	}
	res2B, _ := json.MarshalIndent(logger.Log, "", "\t")
	println(white + string(res2B) + reset)
}

func Info(message *map[string]interface{}) {
	logger.Log.Level = "INFO"
	logger.Log.Message = message
	pc, file, line, ok := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	fn := funcName[strings.LastIndex(funcName, ".")+1:]

	if ok {
		logger.Log.Line = file + ":" + fmt.Sprintf("%v", line)
		logger.Log.Method = fn
	}
	res2B, _ := json.MarshalIndent(logger.Log, "", "\t")
	println(blue + string(res2B) + reset)
}

func Warning(message *map[string]interface{}) {
	logger.Log.Level = "WARNING"
	logger.Log.Message = message
	pc, file, line, ok := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	fn := funcName[strings.LastIndex(funcName, ".")+1:]

	if ok {
		logger.Log.Line = file + ":" + fmt.Sprintf("%v", line)
		logger.Log.Method = fn
	}
	res2B, _ := json.MarshalIndent(logger.Log, "", "\t")
	println(yellow + string(res2B) + reset)
}

func Error(message *map[string]interface{}) {
	logger.Log.Level = "ERROR"
	for k, v := range *message {
		val, ok := v.(error)
		if ok {
			v := val.Error()
			(*message)[k] = v
		}
	}
	logger.Log.Message = message
	pc, file, line, ok := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	fn := funcName[strings.LastIndex(funcName, ".")+1:]

	if ok {
		logger.Log.Line = file + ":" + fmt.Sprintf("%v", line)
		logger.Log.Method = fn
	}
	res2B, _ := json.MarshalIndent(logger.Log, "", "\t")
	println(red + string(res2B) + reset)
}
