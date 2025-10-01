package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/natefinch/lumberjack"
)

// Level 仅用于日后扩展，目前可直接用全局 Log
type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var (
	Log  *Logger
	once sync.Once
)

// Logger 封装
type Logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
}

func init() {
	once.Do(func() {
		// 确保日志目录存在
		logDir := "logs"
		if dir := os.Getenv("LOG_DIR"); dir != "" {
			logDir = dir
		}
		_ = os.MkdirAll(logDir, 0755)

		// 同时输出到控制台和按大小切割的文件
		console := os.Stdout
		file := &lumberjack.Logger{
			Filename:   filepath.Join(logDir, "app.log"),
			MaxSize:    100, // MB
			MaxBackups: 5,
			MaxAge:     30, // days
			Compress:   true,
			LocalTime:  true,
		}

		mw := io.MultiWriter(console, file)

		Log = &Logger{
			debug: log.New(mw, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile),
			info:  log.New(mw, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile),
			warn:  log.New(mw, "[WARN]  ", log.Ldate|log.Ltime|log.Lshortfile),
			err:   log.New(mw, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	})
}

// 对外方法
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Output(2, fmt.Sprintf(format, v...))
}
