package mtl

import (
	"io"
	"log/slog"

	slogecho "github.com/samber/slog-echo"
)

var Logger *slog.Logger

/*
func InitLog(ioWriter io.Writer) {
	var opts []kitexzap.Option
	var output zapcore.WriteSyncer
	if os.Getenv("GO_ENV") != "online" {
		opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())))
		output = zapcore.AddSync(ioWriter)
	} else {
		opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())))
		// async log
		output = &zapcore.BufferedWriteSyncer{
			WS:            zapcore.AddSync(ioWriter),
			FlushInterval: time.Minute,
		}
	}
	server.RegisterShutdownHook(func() {
		output.Sync() //nolint:errcheck
	})
	log := kitexzap.NewLogger(opts...)
	klog.SetLogger(log)
	klog.SetLevel(klog.LevelTrace)
	klog.SetOutput(output)
}
*/

func InitLogger(ioWriter io.Writer, serviceName string) (*slog.Logger, slogecho.Config) {
	var opts = &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	// 创建用于echo middleware的Logger
	logger := slog.New(slog.NewTextHandler(ioWriter, opts)).With("service_name", serviceName)
	config := slogecho.Config{
		DefaultLevel:       slog.LevelInfo,
		ClientErrorLevel:   slog.LevelWarn,
		ServerErrorLevel:   slog.LevelError,
		WithSpanID:         true,
		WithTraceID:        true,
		WithRequestHeader:  false,
		WithResponseHeader: false,
	}
	// 创建全局Logger
	opts.AddSource = true
	Logger = slog.New(slog.NewJSONHandler(ioWriter, opts)).With("service_name", serviceName)
	return logger, config
}
