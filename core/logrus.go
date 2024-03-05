package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"os"
	"path"
	"time"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 34
	gray   = 37
)

type LogFormatter struct{}

func (this *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	logSetting := global.Config.LoggerSetting
	// 自定义日期格式
	timeStamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "[%s] [%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", logSetting.Prefix, timeStamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] [%s] \x1b[%dm[%s]\x1b[0m %s\n", logSetting.Prefix, timeStamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// InitLogger 初始化日志
func InitLogger() *logrus.Logger {
	myLog := logrus.New()                                       // 新建一个日志对象
	myLog.SetOutput(os.Stdout)                                  // 设置日志输出到标准输出
	myLog.SetReportCaller(global.Config.LoggerSetting.ShowLine) // 开启返回函数名和行号
	myLog.SetFormatter(&LogFormatter{})                         // 设置自己定义的Formatter
	level, err := logrus.ParseLevel(global.Config.LoggerSetting.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	myLog.SetLevel(level) // 设置最低的日志级别
	InitFileLogger(myLog) // 初始化日志文件
	return myLog
}

// InitDefaultLogger 初始化全局日志
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(global.Config.LoggerSetting.ShowLine)
	logrus.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.LoggerSetting.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) // 设置最低的日志级别
}

type FileDateHook struct {
	file     *os.File
	logPath  string
	fileDate string //判断日期，切换目录
	appName  string
}

func (this *FileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (this *FileDateHook) Fire(entry *logrus.Entry) error {
	line, _ := entry.String()
	timer := entry.Time.Format("2006-01-02_15-04")
	// 判断日期，切换目录
	if this.fileDate != timer {
		this.file.Close()
		os.MkdirAll(fmt.Sprintf("%s/%s", this.logPath, timer), os.ModePerm)
		filename := fmt.Sprintf("%s/%s/%s.log", this.logPath, timer, this.appName)
		this.file, _ = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		this.fileDate = timer
	}
	this.file.Write([]byte(line))
	return nil
}

// InitFileLogger 初始化日志文件
func InitFileLogger(myLog *logrus.Logger) {
	fileDate := time.Now().Format("2006-01-02_15-04")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", global.Config.LoggerSetting.Director, fileDate), os.ModePerm)
	if err != nil {
		myLog.Error(err)
		return
	}
	//创建文件
	filename := fmt.Sprintf("%s/%s/%s.log", global.Config.LoggerSetting.Director, fileDate, global.Config.LoggerSetting.Prefix)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		myLog.Error(err)
		return
	}

	fileHook := FileDateHook{file, global.Config.LoggerSetting.Director, fileDate, global.Config.LoggerSetting.Prefix}
	myLog.AddHook(&fileHook)
	myLog.SetFormatter(&LogFormatter{})
}
