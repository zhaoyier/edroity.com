/*
	@copyright edroity.com
	@author zhaojianwei at 2017.07.14
 */
package log

import (
	"edroity.com/libs/seelog"
	"github.com/funny/goid"
	"os"
	"strconv"
)

var bizLogger seelog.LoggerInterface
var dataLogger seelog.LoggerInterface

func loadLogConfig() error {
	logger, err := seelog.LoggerFromConfigAsFile("./config/basic/bizlog.xml")
	if err != nil {
		return seelog.Critical("err parsing config log file", err)
	}
	err = seelog.UseLogger(logger)
	if err != nil {
		return seelog.Critical("err parsing config log file", err)
	}
	bizLogger = logger
	bizLogger.SetAdditionalStackDepth(1)
	dataLogger, err = seelog.LoggerFromConfigAsFile("./config/basic/datalog.xml")
	if err != nil {
		return seelog.Critical("err parsing config log file", err)
	}
	err = seelog.UseLogger(dataLogger)
	if err != nil {
		return seelog.Critical("err parsing config log file", err)
	}
	dataLogger.SetAdditionalStackDepth(1)
	return nil
}

func newGoroutineIDFormatter(p string) seelog.FormatterFunc {
	return func(message string, level seelog.LogLevel, context seelog.LogContextInterface) interface{} {
		return strconv.FormatInt(goid.Goid(), 10)
	}
}

func init() {
	err := seelog.RegisterCustomFormatter("GoID", newGoroutineIDFormatter)
	if err != nil {
		seelog.Critical("seelog.RegisterCustomFormatter failed.")
		os.Exit(-1)
	}

	loadLogConfig()
}

//Shutdown close all loggers, it should be used in main func as deferred
func Shutdown() {
	bizLogger.Close()
	dataLogger.Close()
}

//Datalog prints logs for the data platform according to config in datalog.xml
func DataLog(args ...interface{}) {
	//flume server碰到\n才会停止接受数据并生成一个event
	//args = append(args, "\n") //format后会带上\n
	dataLogger.Info(args...)
	dataLogger.Flush()
}

//Trace prints trace log according to config in bizlog.xml
func Trace(args ...interface{}) {
	bizLogger.Trace(args...)
	bizLogger.Flush()
}

//Debug prints debug log according to config in bizlog.xml
func Debug(args ...interface{}) {
	bizLogger.Debug(args...)
	bizLogger.Flush()
}

//Info prints info log according to config in bizlog.xml
func Info(args ...interface{}) {
	bizLogger.Info(args...)
	bizLogger.Flush()
}

//Warn prints warn log according to config in bizlog.xml
func Warn(args ...interface{}) error {
	defer bizLogger.Flush()
	return bizLogger.Warn(args...)
}

//Error prints error log according to config in bizlog.xml
func Error(args ...interface{}) error {
	defer bizLogger.Flush()
	return bizLogger.Error(args...)
}

//Critical prints critical log according to config in bizlog.xml
func Critical(args ...interface{}) error {
	defer bizLogger.Flush()
	return bizLogger.Critical(args...)
}

//Tracef prints trace log according to config in bizlog.xml
func Tracef(format string, args ...interface{}) {
	bizLogger.Tracef(format, args...)
	bizLogger.Flush()
}

//Debugf prints debug log according to config in bizlog.xml
func Debugf(format string, args ...interface{}) {
	bizLogger.Debugf(format, args...)
	bizLogger.Flush()
}

//Infof prints info log according to config in bizlog.xml
func Infof(format string, args ...interface{}) {
	bizLogger.Infof(format, args...)
	bizLogger.Flush()
}

//Warnf prints warn log according to config in bizlog.xml
func Warnf(format string, args ...interface{}) error {
	defer bizLogger.Flush()
	return bizLogger.Warnf(format, args...)
}

//Errorf prints error log according to config in bizlog.xml
func Errorf(format string, args ...interface{}) error {
	defer bizLogger.Flush()
	return bizLogger.Errorf(format, args...)
}

//Criticalf prints critical log according to config in bizlog.xml
func Criticalf(format string, args ...interface{}) error {
	defer bizLogger.Flush()
	return bizLogger.Criticalf(format, args...)
}
