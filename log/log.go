package log

import (
	"fmt"
	"io"
	"time"

	"github.com/rs/zerolog"
)

func buildEvent(context []interface{}) *zerolog.Event {
	e := zerolog.Dict()

	var (
		k  string
		ok bool
	)

	for i, x := range context {
		if i%2 == 0 {
			k, ok = x.(string)
			if !ok {
				k = fmt.Sprint(x)
			}
			continue
		}

		// Skip pairs with empty keys.
		if k == "" {
			continue
		}

		switch v := x.(type) {
		case nil:
			continue
		case string:
			e = e.Str(k, v)
		case []byte:
			e = e.Bytes(k, v)
		case bool:
			e = e.Bool(k, v)
		case error:
			e = e.Str(k, v.Error())
		case int:
			e = e.Int(k, v)
		case int8:
			e = e.Int8(k, v)
		case int16:
			e = e.Int16(k, v)
		case int32:
			e = e.Int32(k, v)
		case int64:
			e = e.Int64(k, v)
		case uint:
			e = e.Uint(k, v)
		case uint8:
			e = e.Uint8(k, v)
		case uint16:
			e = e.Uint16(k, v)
		case uint32:
			e = e.Uint32(k, v)
		case uint64:
			e = e.Uint64(k, v)
		case float32:
			e = e.Float32(k, v)
		case float64:
			e = e.Float64(k, v)
		case time.Time:
			e = e.Time(k, v)
		case time.Duration:
			e = e.Dur(k, v)
		default:
			e = e.Interface(k, v)
		}
	}

	return e
}

type Logger interface {
	EnableDebug(bool)
	Info(event string, context ...interface{})
	Debug(event string, context ...interface{})
	With(context ...interface{}) Logger
	Prefix(string) Logger
}

type logger struct {
	l      zerolog.Logger
	prefix string
}

func (l *logger) EnableDebug(debug bool) {
	if debug {
		l.l = l.l.Level(zerolog.DebugLevel)
	} else {
		l.l = l.l.Level(zerolog.InfoLevel)
	}
}

func (l *logger) Info(event string, context ...interface{}) {
	e := l.l.WithLevel(zerolog.InfoLevel).
		Str("event", l.prefix+event).
		Int64("time", time.Now().UnixNano())

	if len(context) > 0 {
		e = e.Dict("context", buildEvent(context))
	}

	e.Msg("")
}

func (l *logger) Debug(event string, context ...interface{}) {
	e := l.l.WithLevel(zerolog.DebugLevel).
		Str("event", l.prefix+event).
		Int64("time", time.Now().UnixNano())

	if len(context) > 0 {
		e = e.Dict("context", buildEvent(context))
	}

	e.Msg("")
}

func (l *logger) With(context ...interface{}) Logger {
	c := l.l.With()

	if len(context) > 0 {
		c = c.Dict("context", buildEvent(context))
	}

	x := *l
	x.l = c.Logger()
	return &x
}

func (l *logger) Prefix(p string) Logger {
	p += "."

	if l.prefix != "" {
		p = l.prefix + p
	}

	x := *l
	x.prefix = p
	return &x
}

func New(w io.Writer) Logger {
	return &logger{
		l: zerolog.New(w).Level(zerolog.InfoLevel),
	}
}
