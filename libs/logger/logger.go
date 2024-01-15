package logger

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi"

	"github.com/amonaco/goapi/libs/config"
)

var enabled = false

func Init() {
	conf := config.Get()

	if enabled || conf.Environment == "local" {
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Environment: conf.Environment,
		Dsn:         "https://74fef517a6b84c35ac744436389300a2@sentry.io/5177921",
	})

	enabled = err == nil
	if !enabled {
		log.Println("Could not initialize sentry")
	}
}

// Flush waits for sentry to send all error reports
func Flush() {
	if enabled {
		sentry.Flush(time.Second * 2)
	}
}

// Middleware captures errors and sends them to sentry.io
func AddMiddleware(r chi.Router) {
	if enabled {
		r.Use(sentryhttp.New(sentryhttp.Options{
			Repanic: true,
		}).Handle)
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(&wrapRes{w, r.Context(), 0}, r)
			})
		})
	}
}

// Info logs debug informations
func Info(v ...interface{}) {
	log.Println(v...)
}

// Error logs an error and reports it on sentry
func Error(err error) {
	log.Println(err)
	if enabled {
		sentry.CaptureException(err)
	}
}

func ErrorWithContext(err error, ctx context.Context) {
	log.Println(err)

	if !enabled {
		return
	}

	hub := sentry.GetHubFromContext(ctx)

	if hub != nil {
		hub.CaptureException(err)
	} else {
		sentry.CaptureException(err)
	}
}

type wrapRes struct {
	http.ResponseWriter
	context    context.Context
	statusCode int
}

func (w *wrapRes) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *wrapRes) Write(buf []byte) (int, error) {
	if w.statusCode == 500 {
		ErrorWithContext(errors.New(string(buf)), w.context)
	}
	return w.ResponseWriter.Write(buf)
}
