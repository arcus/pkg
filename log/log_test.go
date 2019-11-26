package log

import (
	"bytes"
	"errors"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	var w bytes.Buffer
	l := New(&w)

	l.Info("test")
	if len(w.Bytes()) == 0 {
		t.Errorf("no output")
	}
	t.Log(w.String())
}

func TestInfoContext(t *testing.T) {
	var (
		w   bytes.Buffer
		exp = regexp.MustCompile(`{"level":"info","foo":1,"bar":2,"event":"test","time":[0-9]+}`)
	)

	l := New(&w)
	l.Info("test",
		"foo", 1,
		"bar", 2,
	)

	if !exp.Match(w.Bytes()) {
		t.Error("expected output with context not matched")
	}
	t.Log(w.String())

	w.Reset()

	// Should not have context from first info call
	exp = regexp.MustCompile(`{"level":"info","event":"test","time":[0-9]+}`)

	l.Info("test")
	if !exp.Match(w.Bytes()) {
		t.Error("expected output without context not matched")
	}
	t.Log(w.String())
}

func TestContext(t *testing.T) {
	var (
		w   bytes.Buffer
		exp = regexp.MustCompile(`{"level":"info","foo":1,"bar":2,"event":"test","time":[0-9]+}`)
	)
	l := New(&w).With(
		"foo", 1,
		"bar", 2,
	)

	l.Info("test")
	if !exp.Match(w.Bytes()) {
		t.Error("expected output not matched")
	}
	t.Log(w.String())
}

func TestError(t *testing.T) {
	var (
		w   bytes.Buffer
		exp = regexp.MustCompile(`{"level":"info","event":"test","time":[0-9]+}`)
	)
	l := New(&w)

	l.Info(errors.New("test"))
	if !exp.Match(w.Bytes()) {
		t.Error("expected output not matched")
	}
	t.Log(w.String())
}

type Foo struct{}

func (f Foo) String() string {
	return "foo"
}

func TestStringer(t *testing.T) {
	var (
		w   bytes.Buffer
		exp = regexp.MustCompile(`{"level":"info","event":"foo","time":[0-9]+}`)
	)
	l := New(&w)

	l.Info(Foo{})
	if !exp.Match(w.Bytes()) {
		t.Error("expected output not matched")
	}
	t.Log(w.String())
}

func TestInt(t *testing.T) {
	var (
		w   bytes.Buffer
		exp = regexp.MustCompile(`{"level":"info","event":"13","time":[0-9]+}`)
	)
	l := New(&w)

	l.Info(13)
	if !exp.Match(w.Bytes()) {
		t.Error("expected output not matched")
	}
	t.Log(w.String())
}

func TestDebug(t *testing.T) {
	var w bytes.Buffer
	l := New(&w)

	// Disabled by default.
	l.Debug("test")
	if len(w.Bytes()) != 0 {
		t.Errorf("got: %s", w.String())
	}

	w.Reset()

	// Enable it.
	l.EnableDebug(true)

	l.Debug("test")
	if len(w.Bytes()) == 0 {
		t.Error("no output")
	}
	t.Log(w.String())
}

func TestPrefix(t *testing.T) {
	var w bytes.Buffer
	l := New(&w).Prefix("foo")
	l.Info("test")
	if !strings.Contains(w.String(), "foo.test") {
		t.Error("wrong event")
	}
	t.Log(w.String())

	w.Reset()

	l = l.Prefix("bar.baz")
	l.Info("test")
	if !strings.Contains(w.String(), "foo.bar.baz.test") {
		t.Error("wrong event")
	}
	t.Log(w.String())
}

func BenchmarkInfo(b *testing.B) {
	l := New(ioutil.Discard)

	for i := 0; i < b.N; i++ {
		l.Info("test")
	}
}

func BenchmarkInfo_Context(b *testing.B) {
	l := New(ioutil.Discard)

	for i := 0; i < b.N; i++ {
		l.Info("test",
			"str", "foo",
			"int", 1,
			"float32", 4.5132,
			"bool", true,
			"dur", time.Second,
		)
	}
}

func BenchmarkInfo_Prefix(b *testing.B) {
	l := New(ioutil.Discard).Prefix("foo.bar")

	for i := 0; i < b.N; i++ {
		l.Info("test")
	}
}
