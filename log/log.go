package log

import (
	"fmt"
	"io"
	"time"

	"github.com/rs/zerolog"
)

func buildContext(c zerolog.Context, context []interface{}) zerolog.Context {
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
			c = c.Str(k, v)
		case []byte:
			c = c.Bytes(k, v)
		case bool:
			c = c.Bool(k, v)
		case error:
			c = c.Str(k, v.Error())
		case int:
			c = c.Int(k, v)
		case int8:
			c = c.Int8(k, v)
		case int16:
			c = c.Int16(k, v)
		case int32:
			c = c.Int32(k, v)
		case int64:
			c = c.Int64(k, v)
		case uint:
			c = c.Uint(k, v)
		case uint8:
			c = c.Uint8(k, v)
		case uint16:
			c = c.Uint16(k, v)
		case uint32:
			c = c.Uint32(k, v)
		case uint64:
			c = c.Uint64(k, v)
		case float32:
			c = c.Float32(k, v)
		case float64:
			c = c.Float64(k, v)
		case time.Time:
			c = c.Time(k, v)
		case time.Duration:
			c = c.Dur(k, v)
		default:
			c = c.Interface(k, v)
		}
	}

	return c
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

func (l *logger) log(lvl zerolog.Level, event string, context ...interface{}) {
	zl := l.l

	if len(context) > 0 {
		c := zl.With()
		c = buildContext(c, context)
		zl = c.Logger()
	}

	zl.WithLevel(lvl).
		Str("event", l.prefix+event).
		Int64("time", time.Now().UnixNano()).
		Msg("")
}

func (l *logger) Info(event string, context ...interface{}) {
	l.log(zerolog.InfoLevel, event, context...)
}

func (l *logger) Debug(event string, context ...interface{}) {
	l.log(zerolog.DebugLevel, event, context...)
}

func (l *logger) With(context ...interface{}) Logger {
	if len(context) == 0 {
		return l
	}

	c := l.l.With()
	c = buildContext(c, context)
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
