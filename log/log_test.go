package mylog

import "testing"

func TestLog(t *testing.T) {
	logger := NewLogger("./data/mylog.txt")
	Log(logger)
}

func TestSlog(t *testing.T) {
	slogger := NewSlogger("./data/myslog.txt")
	Slog(slogger)
}
