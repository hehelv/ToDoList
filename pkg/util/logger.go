package util

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
	"time"
)

var LogrusObj *logrus.Logger

func init() {
	if LogrusObj != nil {
		src, _ := setOutputFile()
		LogrusObj.Out = src
		return
	}

	logger := logrus.New()
	src, _ := setOutputFile()

	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	LogrusObj = logger
}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	logFilePath := dir + "/logs/"

	_, err = os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	// 日志文件
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := filepath.Join(logFilePath, logFileName)

	if _, err = os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			if _, err := os.Create(fileName); err != nil {
				log.Println(err.Error())
				return nil, err
			}
		}
	}

	// 创建写入文件，获得io.writer
	out, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	return out, nil
}
