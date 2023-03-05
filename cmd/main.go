package main

import (
	"flag"
	izbot "github.com/iabzal/iskanderzhuma"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	token = flag.String("token", "6175074286:AAFiBnNELNjCSp5N7K2JfPEpeSqCnVZ30Wk", "-token=qwerty")
	debug = flag.Bool("debug", false, "-debug=true")
)

func main() {
	logger := mustLogger(*debug)
	bot, err := izbot.New(logger, *token)
	if err != nil {
		log.Fatal("failed to create bot", err)
	}

	go bot.Run()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	<-stopCh

	bot.Stop()

	logger.Info("Bot gracefully stopped")
}

func mustLogger(debug bool) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableStacktrace = true

	logLevel := zapcore.InfoLevel
	if debug {
		logLevel = zapcore.DebugLevel
	}
	cfg.Level = zap.NewAtomicLevelAt(logLevel)

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	return logger
}
