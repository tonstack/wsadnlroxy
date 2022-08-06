package log

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func getGoId() int {
	var buf [64]byte

	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)

	if err != nil {
		return -1
	}
	return id
}

func Configure() error {
	logrus.SetReportCaller(true)

	logrusLvl, err := logrus.ParseLevel(os.Getenv("LOGL"))
	if err != nil {
		return fmt.Errorf("can't load \"LOGL\": %s", err)
	}

	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		ForceColors:            true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(
				"[go: %d]%s:%d",
				getGoId(), formatFilePath(f.File), f.Line,
			)
		},
	}

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrusLvl)
	logrus.SetFormatter(formatter)

	return nil
}
