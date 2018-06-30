package firewood

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/negroni"
)

type Logger struct{}

// NewLogger returns a new Logger instance
func NewLogger() *Logger {
	return &Logger{}
}

func (*Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	if origin := r.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, HEAD, OPTIONS, PUT, PATCH, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type") //有使用自定义头 需要这个,Action, Module是例子
	}

	if r.Method != "OPTIONS" {
		next(rw, r)
	}

	res := rw.(negroni.ResponseWriter)
	logrus.Infof("[%s] %s %s %d %v", start.String(), r.Method, r.URL.Path, res.Status(), time.Since(start))
}