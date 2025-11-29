package clog_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"

	"github.com/kumose-go/clog"
)

func TestLoggerOrdering(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var l sync.Mutex
	var outs [][]byte
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var out bytes.Buffer
			clog := clog.New(&out)
			clog.WithError(fmt.Errorf("here")).Info("a")
			clog.Debug("debug")
			clog.Debugf("warn %d", 1)
			clog.Info("info")
			clog.Infof("warn %d", 1)
			clog.Warn("warn")
			clog.Warnf("warn %d", 1)
			clog.Error("error")
			clog.Errorf("warn %d", 1)
			clog.WithField("foo", "bar").Info("foo")
			clog.IncreasePadding()
			clog.Info("increased")
			clog.WithoutPadding().WithField("foo", "bar").Info("without padding")
			clog.Info("increased")
			clog.ResetPadding()
			l.Lock()
			outs = append(outs, out.Bytes())
			l.Unlock()
		}()
	}
	wg.Wait()
	for i := range len(outs) - 1 {
		s1 := string(outs[i])
		s2 := string(outs[i+1])
		if s1 != s2 {
			t.Errorf("at least one of the outputs is different:\n%q\nvs\n%q\n", s1, s2)
		}
	}
	requireEqualOutput(t, outs[0])
}

func TestLogger_printf(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var out bytes.Buffer
	l := clog.New(&out)
	l.Infof("logged in %s", "Tobi")
	requireEqualOutput(t, out.Bytes())
}

func TestLogger_levels(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var out bytes.Buffer
	l := clog.New(&out)

	l.Debug("uploading")
	l.Info("upload complete")
	requireEqualOutput(t, out.Bytes())
}

func TestLogger_WithField(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var out bytes.Buffer
	l := clog.New(&out)

	ctx := l.WithField("file", "sloth.png").WithField("user", "Tobi")
	ctx.Debug("uploading")
	ctx.Info("upload complete")
	requireEqualOutput(t, out.Bytes())
}

func TestLogger_HandlerFunc(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var out bytes.Buffer
	l := clog.New(&out)

	l.Infof("logged in %s", "Tobi")
	requireEqualOutput(t, out.Bytes())
}

func BenchmarkLogger_small(b *testing.B) {
	var out bytes.Buffer
	l := clog.New(&out)

	for i := 0; i < b.N; i++ {
		l.Info("login")
	}
}

func BenchmarkLogger_medium(b *testing.B) {
	var out bytes.Buffer
	l := clog.New(&out)

	for i := 0; i < b.N; i++ {
		l.WithField("file", "sloth.png").
			WithField("type", "image/png").
			WithField("size", 1<<20).
			Info("upload")
	}
}

func BenchmarkLogger_large(b *testing.B) {
	var out bytes.Buffer
	l := clog.New(&out)

	err := fmt.Errorf("boom")

	for i := 0; i < b.N; i++ {
		l.WithField("file", "sloth.png").
			WithField("type", "image/png").
			WithField("size", 1<<20).
			WithField("some", "more").
			WithField("data", "here").
			WithField("whatever", "blah blah").
			WithField("more", "stuff").
			WithField("context", "such useful").
			WithField("much", "fun").
			WithError(err).Error("upload failed")
	}
}
