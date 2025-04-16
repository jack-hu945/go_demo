package mylog

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

func NewLogger(logFile string) *log.Logger {
	//append log
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o666)
	if err != nil {
		fmt.Println("open log file error:", err)
	}

	logger := log.New(fout, "[MY_BIZ]", log.Ldate|log.Lmicroseconds)
	return logger
}

func Log(logger *log.Logger) {
	logger.Printf("%d + %d = %d\n", 1, 2, 1+2)
	logger.Printf("hello sean\n")
	logger.Println("hello sean1")

	log.Printf("log ++ 123\n")
	log.Println("hello sean3")
	//logger.Fatal("fatal")// witll os.Exit(1)
}

func NewSlogger(logFile string) *slog.Logger {
	//append log
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o666)
	if err != nil {
		fmt.Println("open log file error:", err)
	}

	slogger := slog.New(
		slog.NewJSONHandler(fout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		}),

	// slog.NewTextHandler(fout, &slog.HandlerOptions{
	// 	AddSource: : true,
	// 	Level:     slog.LevelInfo,
	// }),
	)
	return slogger
}

func Slog(slogger *slog.Logger) {
	slogger.Debug("加法运算")
	slogger.Info("hello sean\n")
	slogger.Error("hello sean1")

	slog.Debug("log ++ 123\n")
	slog.Info("hello sean3")
	//logger.Fatal("fatal")// witll os.Exit(1)
}
