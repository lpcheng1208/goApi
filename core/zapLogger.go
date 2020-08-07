package core

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjackv2 "gopkg.in/natefinch/lumberjack.v2"
	"goApi/global"
	"io"
	"os"
	"strings"
	"time"
)

const (
	// WriterStdOut 标准输出
	WriterStdOut = "stdout"
	// WriterFile 文件输出
	WriterFile = "file"

	// RotateTimeDaily 按天切割
	RotateTimeDaily = "daily"
	// RotateTimeHourly 按小时切割
	RotateTimeHourly = "hourly"
)

// 在 main 函数之前被调用，根据调用关系决定执行的顺序
// newZapLogger new zap logger
func init() {
	encoder := getJSONEncoder()

	cfg := global.SERVER_CONFIG.ZapConfig

	var cores []zapcore.Core
	var options []zap.Option
	// 设置初始化字段
	option := zap.Fields(zap.String("serviceName", "goApi Service"))
	options = append(options, option)

	allLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.FatalLevel
	})

	writers := strings.Split(cfg.Writers, ",")
	for _, w := range writers {
		if w == WriterStdOut {
			core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
			cores = append(cores, core)
		}
		if w == WriterFile {
			infoFilename := cfg.LoggerFile
			infoWrite := getLogWriterWithRoll(infoFilename)
			warnFilename := cfg.LoggerWarnFile
			warnWrite := getLogWriterWithRoll(warnFilename)
			errorFilename := cfg.LoggerErrorFile
			errorWrite := getLogWriterWithRoll(errorFilename)
			infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl <= zapcore.InfoLevel
			})
			warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				stacktrace := zap.AddStacktrace(zapcore.WarnLevel)
				options = append(options, stacktrace)
				return lvl == zapcore.WarnLevel
			})
			errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				stacktrace := zap.AddStacktrace(zapcore.ErrorLevel)
				options = append(options, stacktrace)
				return lvl >= zapcore.ErrorLevel
			})

			core := zapcore.NewCore(encoder, zapcore.AddSync(infoWrite), infoLevel)
			cores = append(cores, core)
			core = zapcore.NewCore(encoder, zapcore.AddSync(warnWrite), warnLevel)
			cores = append(cores, core)
			core = zapcore.NewCore(encoder, zapcore.AddSync(errorWrite), errorLevel)
			cores = append(cores, core)
		}
		if w != WriterFile && w != WriterStdOut {
			core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
			cores = append(cores, core)
			allWriter := getLogWriterWithRoll(cfg.LoggerFile)
			core = zapcore.NewCore(encoder, zapcore.AddSync(allWriter), allLevel)
			cores = append(cores, core)
		}
	}

	combinedCore := zapcore.NewTee(cores...)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	options = append(options, caller)
	// 开启文件及行号
	development := zap.Development()
	options = append(options, development)
	// 跳过文件调用层数
	//addCallerSkip := zap.AddCallerSkip(2)
	//options = append(options, addCallerSkip)

	// 构造日志
	logger := zap.New(combinedCore, options...)
	global.LOGGER = logger

}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000Z0700"))
}

// getJSONEncoder
func getJSONEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     TimeEncoder,                    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器:FullCallerEncoder
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriterWithRoll
func getLogWriterWithRoll(filename string) io.Writer {
	logFullPath := filename
	rotationPolicy := global.VIPER_CONFIG.Get("logger.log_rolling_policy")
	backupCount := global.VIPER_CONFIG.GetInt("logger.log_backup_count")
	logRollingType := global.VIPER_CONFIG.GetString("logger.log_rolling_type")
	logRotateSize := global.VIPER_CONFIG.GetInt("logger.log_rotate_size")
	// 默认
	rotateDuration := time.Hour * 24
	if rotationPolicy == RotateTimeHourly {
		rotateDuration = time.Hour
	}
	length := len(filename) - 4
	logFile := filename[:length] + ".%Y%m%d%H" + ".log"
	var hook io.Writer
	switch logRollingType {
	case "time":
		hookFile, err := rotatelogs.New(
			logFile,                              //logFullPath+".%Y%m%d%H", // 时间格式使用shell的date时间格式
			rotatelogs.WithLinkName(logFullPath), // 生成软链，指向最新日志文件
			rotatelogs.WithRotationCount(backupCount),   // 文件最大保存份数
			rotatelogs.WithRotationTime(rotateDuration), // 日志切割时间间隔
			rotatelogs.WithLocation(time.UTC),           // 日志轮转的时间
		)
		if err != nil {
			panic(err)
		}
		hook = hookFile
	case "size":
		hookLum := lumberjackv2.Logger{
			Filename:   logFullPath,   // 日志文件路径
			MaxSize:    logRotateSize, // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: backupCount,   // 日志文件最多保存多少个备份
			MaxAge:     7,             // 文件最多保存多少天
			Compress:   true,          // 是否压缩
		}
		hook = &hookLum
	default:
		hookLum := lumberjackv2.Logger{
			Filename:   logFullPath,   // 日志文件路径
			MaxSize:    logRotateSize, // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: backupCount,   // 日志文件最多保存多少个备份
			MaxAge:     7,             // 文件最多保存多少天
			Compress:   true,          // 是否压缩
		}
		hook = &hookLum
	}

	return hook
}
