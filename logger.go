package helpers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/moisespsena-go/aorm"
	"github.com/moisespsena-go/path-helpers"
	"github.com/vjeantet/jodaTime"
)

var now = time.Now()

type SimpleFileLogger struct {
	Pth string
	Err func(err error, args ...interface{})
	f   io.WriteCloser
}

func NewSimpleFileLogger(dir string, Err func(err error, args ...interface{})) *SimpleFileLogger {
	return &SimpleFileLogger{
		filepath.Join(dir, jodaTime.Format("YYYY/MM/dd/YYYY-MM-dd_HHmmss.S.log", now)),
		Err,
		nil,
	}
}

func (l *SimpleFileLogger) Log(args ...interface{}) {
	var msg = jodaTime.Format("YYYY-MM-dd HH:mm:ss.S", time.Now()) + " "
	if _, err := l.Write([]byte(msg + fmt.Sprint(args...))); err != nil {
		l.Err(err, args...)
	}
}

func (l *SimpleFileLogger) Write(p []byte) (n int, err error) {
	if l.f == nil {
		var s os.FileInfo
		if s, err = os.Stat(l.Pth); err != nil {
			if os.IsNotExist(err) {
				if err = path_helpers.MkdirAllIfNotExists(filepath.Dir(l.Pth)); err == nil {
					l.f, err = os.Create(l.Pth)
				}
			}
		} else {
			l.f, err = os.OpenFile(l.Pth, os.O_APPEND|os.O_WRONLY, s.Mode())
		}
		if err == nil {
			l.f.Write(p)
		}
	}
	return
}

func (l *SimpleFileLogger) Close() error {
	if l.f != nil {
		return l.f.Close()
	}
	return nil
}

func SimpleFileScopeQueryLogger(dir string, err func(err error, args ...interface{})) (logger *SimpleFileLogger, cb func(sction string, scope *aorm.Scope)) {
	logger = NewSimpleFileLogger(dir, err)
	return logger, func(action string, scope *aorm.Scope) {
		msg := "[" + scope.RealTableName() + "] " + strings.ToUpper(action) + " "
		switch action {
		case "update", "delete":
			if f, ok := scope.FieldByName("ID"); ok {
				msg += "#" + f.Field.String() + ""
			}
		}
		msg += ": " + scope.QueryString()
		logger.Log(msg)
	}
}

func SimpleFileScopeErrorLogger(dir string, err func(err error, args ...interface{})) (logger *SimpleFileLogger, cb func(scope *aorm.Scope, err error)) {
	logger = NewSimpleFileLogger(dir, err)
	return logger, func(scope *aorm.Scope, err error) {
		logger.Log(fmt.Sprint("["+scope.RealTableName()+"]", err))
	}
}
