package main

import (
	"go.uber.org/zap"
	"encoding/json"
	"go.uber.org/zap/zapcore"
)

func main(){
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "./logtest.log"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase",
		"timeKey": "time"
	  }
	}`)
	var cfg zap.Config
	cfg.EncoderConfig.EncodeTime=zapcore.TimeEncoderOfLayout("01-02 : 2006 15")
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")
}
