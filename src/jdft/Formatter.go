package jdft

import (
	"go.uber.org/zap/zapcore"
	"log"
)

var FormatterMap map[string]*LogFormatter

func init() {
	FormatterMap = make(map[string]*LogFormatter)
	parseFomatterFromGlobalSettings()
}

// parseFomatterFromGlobalSettings 将GlobalSetting中的LOGGER.format解析为FormatterMap
func parseFomatterFromGlobalSettings() {
	logger, _ := GlobalSettings["LOGGER"]
	if logger == nil {
		log.Fatalln("[error]解析日志配置错误 logger")
	}
	formatters, _ := logger.(map[string]interface{})["FORMATTERS"]
	if formatters == nil {
		log.Fatalln("[error]解析日志配置错误 format")
	}
	for key, value := range formatters.(map[string]interface{}) {
		m := value.(map[string]interface{})
		time := ""
		level := ""
		caller := ""
		message := ""
		timeLayout := "2006-01-02 15:04:05"
		if t, ok := m["time"]; ok {
			time = t.(string)
		}
		if l, ok := m["level"]; ok {
			level = l.(string)
		}
		if c, ok := m["caller"]; ok {
			caller = c.(string)
		}
		if m, ok := m["message"]; ok {
			message = m.(string)
		}
		if t, ok := m["timeLayout"]; ok {
			timeLayout = t.(string)
		}
		formatter := LogFormatter{
			TimeKey:    time,
			LevelKey:   level,
			CallerKey:  caller,
			MessageKey: message,
			TimeLayout: timeLayout,
		}
		FormatterMap[key] = &formatter
	}
}

type LogFormatter struct {
	TimeKey    string `json:"time"`
	LevelKey   string `json:"level"`
	CallerKey  string `json:"caller"`
	MessageKey string `json:"message"`
	TimeLayout string `json:"timeLayout"`
}

func (l *LogFormatter) ToEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        l.TimeKey,
		LevelKey:       l.LevelKey,
		NameKey:        "logger",
		CallerKey:      l.CallerKey,
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     l.MessageKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(l.TimeLayout),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
