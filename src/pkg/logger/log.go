package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"

	kitlog "github.com/go-kit/log"
)

var logger kitlog.Logger

func Init() {
	var (
		logout io.Writer
	)
	logout = os.Stdout

	w := kitlog.NewSyncWriter(logout)
	logger = kitlog.NewLogfmtLogger(w)
}

func Log(keyvals ...interface{}) {
	_, file, no, ok := runtime.Caller(1)
	ts := time.Now()
	caller := "%s:%d"
	if ok {
		caller = fmt.Sprintf(caller, path.Base(file), no)
	} else {
		caller = "unknow.go"
	}

	params := append([]interface{}{"ts", ts, "caller", caller}, keyvals...)
	logger.Log(params...)
}
