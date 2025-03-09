package logging

import (
	"context"
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	globalLogger *logger
	initOnce     sync.Once
)

type Option struct {
	WithFunc   bool
	LogLevel   string
	Formatter  string // "text" 或 "json"
	OutputPath string // 文件路径，空表示 stdout
}

// 初始化全局日志实例（线程安全）
func Setup(name string, option *Option) {
	initOnce.Do(func() {
		opt := Option{
			WithFunc:   true,
			LogLevel:   option.LogLevel,
			Formatter:  option.Formatter,
			OutputPath: "",
		}
		if option != nil {
			if option.WithFunc {
				opt.WithFunc = option.WithFunc
			}
			if option.LogLevel != "" {
				opt.LogLevel = option.LogLevel
			}
			if option.Formatter != "" {
				opt.Formatter = option.Formatter
			}
			if option.OutputPath != "" {
				opt.OutputPath = option.OutputPath
			}
		}

		// 创建基础 logger
		logrusLogger := logrus.New()

		// 配置日志级别
		level, err := logrus.ParseLevel(opt.LogLevel)
		if err != nil {
			level = logrus.InfoLevel
		}
		logrusLogger.SetLevel(level)

		// 配置输出格式
		switch opt.Formatter {
		case "json":
			logrusLogger.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
			})
		default:
			logrusLogger.SetFormatter(&logrus.TextFormatter{
				FullTimestamp:   true,
				TimestampFormat: "2006-01-02 15:04:05",
			})
		}

		// 配置输出目标
		if opt.OutputPath != "" {
			file, err := os.OpenFile(opt.OutputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err == nil {
				logrusLogger.SetOutput(file)
			} else {
				logrusLogger.Warn("Failed to log to file, using default stderr")
			}
		}

		globalLogger = &logger{
			name:   name,
			option: opt,
			entry_: logrusLogger.WithField("service", name), // 基础字段
		}
	})
}

// 获取全局日志实例（确保先调用 Setup）
func L() *logger {
	if globalLogger == nil {
		panic("logger not initialized, call Setup first")
	}
	return globalLogger
}

type logger struct {
	name   string
	option Option
	entry_ *logrus.Entry
}

func Infof(format string, args ...any) {
	L().getEntry().Infof(format, args...)
}

func Warnf(format string, args ...any) {
	L().getEntry().Warnf(format, args...)
}

func Errorf(format string, args ...any) {
	L().getEntry().Errorf(format, args...)
}

func Printf(format string, args ...any) {
	L().getEntry().Printf(format, args...)
}

func Println(args ...any) {
	L().getEntry().Println(args...)
}

func Fatalf(format string, args ...any) {
	L().getEntry().Fatalf(format, args...)
}

func Fatalln(args ...any) {
	L().getEntry().Fatalln(args...)
}

// 所有日志方法改为包级函数
func Trace(args ...any) {
	L().getEntry().Log(logrus.TraceLevel, args...)
}

func Debug(args ...any) {
	L().getEntry().Debug(args...)
}

func Info(args ...any) {
	L().getEntry().Info(args...)
}

func Warn(args ...any) {
	L().getEntry().Warn(args...)
}

func Error(args ...any) {
	L().getEntry().Error(args...)
}

// 带上下文的日志方法
func WithContext(ctx context.Context) *logrus.Entry {
	return L().getEntry().WithContext(ctx)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return L().getEntry().WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return L().getEntry().WithFields(fields)
}

func WithError(err error) *logrus.Entry {
	return L().getEntry().WithError(err)
}

// 其他方法同理...
// [保留原有 entry() 逻辑，但改为使用全局 entry]

// 私有方法
func (l *logger) getEntry() *logrus.Entry {
	entry := l.entry_
	if l.option.WithFunc {
		entry = entry.WithField("caller", getCaller(4)) // 修正跳转层级
	}
	return entry
}

func getCaller(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	f := runtime.FuncForPC(pc)
	_, line := f.FileLine(pc)
	return f.Name() + ":" + strconv.Itoa(line)
}
