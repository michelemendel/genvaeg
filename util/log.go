package util

import (
	"fmt"
	"os"
	"strings"

	"log/slog"
)

func init() {
	InitEnv()
}

// https://betterstack.com/community/guides/logging/logging-in-go/
const (
	LevelTrace  = slog.Level(-8)
	LevelNotice = slog.Level(2)
	LevelFatal  = slog.Level(12)
)

var LevelNames = map[slog.Leveler]string{
	LevelTrace:  "TRACE",
	LevelNotice: "NOTICE",
	LevelFatal:  "FATAL",
}

func StdOutLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, options))
}

var options = &slog.HandlerOptions{
	AddSource:   true,
	Level:       slog.LevelDebug,
	ReplaceAttr: MakeReplaceAttr(),
}

func MakeReplaceAttr() func(groups []string, a slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.LevelKey {
			level := a.Value.Any().(slog.Level)
			levelLabel, exists := LevelNames[level]
			if !exists {
				levelLabel = level.String()
			}
			a.Value = slog.StringValue(levelLabel)
		}

		if a.Key == slog.SourceKey {
			filename := a.Value.Any().(*slog.Source).File
			lineNumber := a.Value.Any().(*slog.Source).Line
			ps := strings.Split(filename, "/")
			a.Value = slog.StringValue(fmt.Sprintf("%s/%s:%v", ps[len(ps)-2], ps[len(ps)-1], lineNumber))
		}

		return a
	}
}
