package logger

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.InfoLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// getLoggerLevel 根据级别获取zapcore.Level
func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

var LoggerMap map[string]mLogger

type mLogger *zap.SugaredLogger

func init() {
	LoggerMap = make(map[string]mLogger)
	parseLoggerFromGlobalSettings()
	fmt.Println("bbbbbbbbbbbbb", LoggerMap)
}

// parseLoggerFromGlobalSettings 将GlobalSetting中的LOGGER解析为LoggerMap
func parseLoggerFromGlobalSettings() {
	logger, _ := jdft.GlobalSettings["LOGGER"]
	if logger == nil {
		log.Fatalln("[error]解析日志配置错误 logger")
	}
	mloggers, _ := logger.(map[string]interface{})["LOGGERS"]
	if mloggers == nil {
		log.Fatalln("[error]解析日志配置错误 format")
	}
	for key, value := range mloggers.(map[string]interface{}) {
		LoggerMap[key] = generateCuttingLogger(value.(map[string]interface{}))
	}
}

func generateCuttingLogger(conf map[string]interface{}) *zap.SugaredLogger {
	fmt.Println("<<<<<<<<<", conf)
	writers := make([]io.Writer,0)
	var filenames []interface{}
	maxage := 2
	maxbackups := 1
	maxsize := 10
	compress := false
	if f,ok:=conf["logPath"];ok {
		filenames = f.([]interface{})
	}
	for _,f := range filenames{
		filename := f.(string)
		switch filename {
		case "stdout":
			writers = append(writers,os.Stdout)
		case "stderr":
			writers = append(writers,os.Stderr)
		default:
			cuttingConfig := conf["cuttingConfig"].(map[string]interface{})
			if m,ok:=cuttingConfig["maxAge"];ok {
				maxage = m.(int)
			}
			if m,ok:=cuttingConfig["maxBackUps"];ok {
				maxbackups = m.(int)
			}
			if m,ok:=cuttingConfig["maxSize"];ok {
				maxsize = m.(int)
			}
			if c,ok:=cuttingConfig["compress"];ok {
				compress = c.(bool)
			}
			jacklogger := &lumberjack.Logger{
				Filename:  filename,
				MaxAge: maxage,
				MaxBackups: maxbackups,
				MaxSize:   maxsize,
				LocalTime: true,
				Compress:  compress,
			}
			writers=append(writers,jacklogger)
		}
	}
	syncWriter := zapcore.AddSync(io.MultiWriter(writers...))
	level := getLoggerLevel("debug") //TODO 改成解析
	formatter := conf["formatter"].(string)
	encoder := FormatterMap[formatter].ToEncoderConfig()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder),syncWriter,zap.NewAtomicLevelAt(level))
	logger := zap.New(core,zap.AddCaller())
	return logger.Sugar()
}

//var GlobalLogger *zap.SugaredLogger

//func init(){
//	filepath := "./jdft.log"
//	if fileName,ok := jdft.GlobalSettings["GLOBAL_LOGGER_FILEPATH"];ok{
//		filepath=fileName.(string)
//	}
//	syncWriter := zapcore.AddSync(&lumberjack.Logger{
//		Filename: filepath,
//		MaxSize: 1<<30,
//		LocalTime: true,
//		Compress: true,
//	})
//	level := getLoggerLevel("debug")
//	encoder := zapcore.EncoderConfig{
//		TimeKey:        "",
//		NameKey:        "logger",
//		CallerKey:      "caller",
//		FunctionKey:    zapcore.OmitKey,
//		StacktraceKey:  "stacktrace",
//		LineEnding:     zapcore.DefaultLineEnding,
//		EncodeLevel:    zapcore.LowercaseLevelEncoder,
//		EncodeTime:     zapcore.EpochTimeEncoder,
//		EncodeDuration: zapcore.SecondsDurationEncoder,
//		EncodeCaller:   zapcore.ShortCallerEncoder,
//	}
//	encoder.EncodeTime=zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
//	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder),syncWriter,zap.NewAtomicLevelAt(level))
//	logger := zap.New(core,zap.AddCaller())
//	GlobalLogger = logger.Sugar()
//}
//
//func Debug(args ...interface{}) {
//	GlobalLogger.Debug(args...)
//}
//
//func Debugf(template string, args ...interface{}) {
//	GlobalLogger.Debugf(template, args...)
//}
//
//func Info(args ...interface{}) {
//	GlobalLogger.Info(args...)
//}
//
//func Infof(template string, args ...interface{}) {
//	GlobalLogger.Infof(template, args...)
//}
//
//func Warn(args ...interface{}) {
//	GlobalLogger.Warn(args...)
//}
//
//func Warnf(template string, args ...interface{}) {
//	GlobalLogger.Warnf(template, args...)
//}
//
//func Error(args ...interface{}) {
//	GlobalLogger.Error(args...)
//}
//
//func Errorf(template string, args ...interface{}) {
//	GlobalLogger.Errorf(template, args...)
//}
//
//func DPanic(args ...interface{}) {
//	GlobalLogger.DPanic(args...)
//}
//
//func DPanicf(template string, args ...interface{}) {
//	GlobalLogger.DPanicf(template, args...)
//}
//
//func Panic(args ...interface{}) {
//	GlobalLogger.Panic(args...)
//}
//
//func Panicf(template string, args ...interface{}) {
//	GlobalLogger.Panicf(template, args...)
//}
//
//func Fatal(args ...interface{}) {
//	GlobalLogger.Fatal(args...)
//}
//
//func Fatalf(template string, args ...interface{}) {
//	GlobalLogger.Fatalf(template, args...)
//}
