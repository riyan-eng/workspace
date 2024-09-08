package config

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() {
	logFilePath := "./logs/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Fatal(err)
		return
	}

	// set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Fatal(err)
			return
		}
	}

	logger := hertzZerolog.New()
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time
		MaxAge:     10,   // A file can be saved for up to 10 days.
		Compress:   true, // Compress with gzip.
	}

	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(hlog.LevelDebug)

	hlog.SetLogger(logger)
}
