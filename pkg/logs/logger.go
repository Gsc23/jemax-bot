package logs

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
	charmLog "github.com/charmbracelet/log"
	gormLogger "gorm.io/gorm/logger"
)

type Logger interface {
	Debug(_ context.Context, format string, args ...any)
	Info(_ context.Context, format string, args ...any)
	Warn(_ context.Context, format string, args ...any)
	Error(_ context.Context, format string, args ...any)
}

var defaultLoggers map[string]Logger = make(map[string]Logger)
var lock sync.Mutex = sync.Mutex{}
var globalConfig Config = Config{
	level:       charmLog.InfoLevel,
	gormLevel:   gormLogger.Info,
	isColorful:  true,
	enableTrace: true,
}

type Config struct {
	level       charmLog.Level
	gormLevel   gormLogger.LogLevel
	isColorful  bool
	enableTrace bool
}

func Setup(level string, color bool, trace bool) {
	switch level {
	case "0", "debug":
		globalConfig.level = charmLog.DebugLevel
		globalConfig.gormLevel = gormLogger.Info
	case "1", "info":
		globalConfig.level = charmLog.InfoLevel
		globalConfig.gormLevel = gormLogger.Info
	case "2", "warn":
		globalConfig.level = charmLog.WarnLevel
		globalConfig.gormLevel = gormLogger.Warn
	case "3", "error":
		globalConfig.level = charmLog.ErrorLevel
		globalConfig.gormLevel = gormLogger.Error
	}

	globalConfig.isColorful = color
	globalConfig.enableTrace = trace
}

func createStyle() *charmLog.Styles {
	styles := charmLog.DefaultStyles()

	// error
	styles.Levels[charmLog.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Background(lipgloss.Color("#fd0000")).
		Foreground(lipgloss.Color("#200000"))
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#fd0000")).Bold(true)

	// warn
	styles.Levels[charmLog.WarnLevel] = lipgloss.NewStyle().
		SetString("WARN").
		Background(lipgloss.Color("#fdfd00")).
		Foreground(lipgloss.Color("#202000"))

	// info
	styles.Levels[charmLog.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Background(lipgloss.Color("#0000fd")).
		Foreground(lipgloss.Color("#000020"))

	// debug
	styles.Levels[charmLog.DebugLevel] = lipgloss.NewStyle().
		SetString("DEBUG").
		Background(lipgloss.Color("#00fd00")).
		Foreground(lipgloss.Color("#002000"))

	return styles
}

type DefaultLogger struct {
	charmLogger *charmLog.Logger
	gormLogger  gormLogger.Interface
}

func DefaultLoggerWithName(name string) *DefaultLogger {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := defaultLoggers[name]; !ok {
		charmLogger := charmLog.NewWithOptions(os.Stdout, charmLog.Options{
			Level:           globalConfig.level,
			Prefix:          fmt.Sprintf("[%s]", name),
			TimeFormat:      time.DateTime,
			ReportTimestamp: true,
			Formatter:       charmLog.TextFormatter,
		})
		if globalConfig.isColorful {
			charmLogger.SetStyles(createStyle())
		}

		gormLogger := gormLogger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), gormLogger.Config{
			SlowThreshold:             500 * time.Millisecond,
			LogLevel:                  gormLogger.Warn,
			Colorful:                  globalConfig.isColorful,
			IgnoreRecordNotFoundError: true,
		})

		defaultLoggers[name] = &DefaultLogger{charmLogger: charmLogger, gormLogger: gormLogger}
	}
	return defaultLoggers[name].(*DefaultLogger)
}

func (c *DefaultLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	switch level {
	case gormLogger.Silent:
		c.charmLogger.SetLevel(charmLog.FatalLevel)
	case gormLogger.Info:
		c.charmLogger.SetLevel(charmLog.InfoLevel)
	case gormLogger.Warn:
		c.charmLogger.SetLevel(charmLog.WarnLevel)
	case gormLogger.Error:
		c.charmLogger.SetLevel(charmLog.ErrorLevel)
	}
	return c
}

func (c *DefaultLogger) Debug(_ context.Context, format string, args ...interface{}) {
	go c.charmLogger.Debugf(format, args...)
}

func (c *DefaultLogger) Info(_ context.Context, format string, args ...interface{}) {
	go c.charmLogger.Infof(format, args...)
}

func (c *DefaultLogger) Warn(_ context.Context, format string, args ...interface{}) {
	go c.charmLogger.Warnf(format, args...)
}
func (c *DefaultLogger) Error(_ context.Context, format string, args ...interface{}) {
	go c.charmLogger.Errorf(format, args...)
}
func (c *DefaultLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if globalConfig.enableTrace {
		c.gormLogger.Trace(ctx, begin, fc, err)
	}
}
