package zap

import (
	"go.uber.org/zap"
	"log"
	"time"
	"path"
	"encoding/json"
	"fmt"
)

var (
	outpath = path.Join("zap.log")
	errpath = path.Join("err.log")
)

func SimpleZapLogger() {
	logger, err:= zap.NewProduction()
	if err != nil {
		log.Println("zap\t-", "new logger error", err.Error())
	}

	defer logger.Sync()
	sugar := logger.Sugar()

	url := "www.baidu.com"
	sugar.Infow("failed to fetch URL",
	// Structured context as loosely typed key-value pairs.
	"url", url,
	"attempts", 3,
	"backoff",time.Second)
	sugar.Infof("Failed to fetch URL: %s", url)
}

func SimpleSavingZapLogger() {
	zapConf := zap.NewDevelopmentConfig()
	zapConf.OutputPaths = []string{outpath}
	zapConf.ErrorOutputPaths = []string{errpath}

	logger, err := zapConf.Build()
	if err != nil {
		panic(err.Error())
	}

	logger.Debug("Starting zap! Have your fun!")

	logger.Info("failed to fetch URL",
		zap.String("url", "www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	log.Println("test\t-", "url", "www.baidu.com", "time", time.Now(), "attempt", 3)

	defer func() {
		logger.Sync()
	}()
}

// 这里需要注意：
// 1.使用前面两种方式，在循环里写log的时候会丢失数据，因为zap的log是有采样的
// 2.下面的json配置中development并没有什么卵用
// 3.outputpaths&erroroutputpaths可以配置多个文件，但是如果不做其他设定的话，文件内容一致
func MultipleSavingZapLogger() {
	rawJSON := []byte(`{
	  "level": "debug",
	  "development": false,
	  "encoding": "json",
	  "outputPaths": ["zap.log"],
	  "errorOutputPaths": ["err.log"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")

	url := "loopring.org"
	for i := 1; i < 100000; i++ {
		logger.Info("saving number", zap.String("url", url), zap.Int("attempt", i))
	}
}

// 第二种情况无法实现既能打印出来又能写入到log文件
func SimpleLoggerAndPrint() {
	zapConf := zap.NewDevelopmentConfig()
	zapConf.OutputPaths = []string{outpath}
	zapConf.ErrorOutputPaths = []string{errpath}

	logger, err := zapConf.Build()
	if err != nil {
		panic(err.Error())
	}

	logger.Info("it is a log test")
	logger.Info(fmt.Sprintln("it is a log println test"))
}