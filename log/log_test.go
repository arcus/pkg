package log

import (
	"bytes"
	"io/ioutil"
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

func TestContext(t *testing.T) {
	var w bytes.Buffer
	l := New(&w).With(
		"foo", 1,
		"bar", 2,
	)

	l.Info("test")
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
