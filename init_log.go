package common

import (
    "github.com/guyannanfei25/go-logger"
)

func InitLogger(logDir, logName string, logLevel int) error {
    if logDir == "" {
        logDir = "/tmp/"
    }

    if logName == "" {
        logName = "info.log"
    }

	if logLevel < 0 {
		logLevel = 0
	}
	if logLevel > 6 {
		logLevel = 6
	}
	logger.SetRollingDaily(logDir, logName)
	logger.SetLevel(logger.LEVEL(logLevel))

	// 关闭Console输出
	logger.SetConsole(false)
	logger.Debugf("InitLog success, logDir: %s, logName: %s, logLevel: %v", logDir, logName, logLevel)

    return nil
}
