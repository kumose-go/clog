package clog_test

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/kumose-go/clog"
)

func TestRootLogOptions(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var out bytes.Buffer
	clog.Log = clog.New(&out)
	clog.SetLevel(clog.DebugLevel)
	clog.SetLevelFromString("info")
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
	requireEqualOutput(t, out.Bytes())
}

func TestRace(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var wg sync.WaitGroup
	for range 9999 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clog.Infof("a")
		}()
	}
	wg.Wait()
}

// Unstructured logging is supported, but not recommended since it is hard to query.
func Example_unstructured() {
	clog.Infof("%s logged in", "Tobi")
}

// Structured logging is supported with fields, and is recommended over the formatted message variants.
func Example_structured() {
	clog.WithField("user", "Tobo").Info("logged in")
}

// Errors are passed to WithError(), populating the "error" field.
func Example_errors() {
	err := errors.New("boom")
	clog.WithError(err).Error("upload failed")
}

var update = flag.Bool("update", false, "update .golden files")

func requireEqualOutput(tb testing.TB, in []byte) {
	tb.Helper()

	bts := useLinuxEOL(in)
	golden := "testdata/" + tb.Name() + ".golden"
	if *update {
		if err := os.MkdirAll(filepath.Dir(golden), 0o755); err != nil {
			tb.Fatal(err)
		}
		if err := os.WriteFile(golden, bts, 0o600); err != nil {
			tb.Fatal(err)
		}
	}

	gbts, err := os.ReadFile(golden)
	if err != nil {
		tb.Fatal(err)
	}
	gbts = useLinuxEOL(gbts)

	if !bytes.Equal(bts, gbts) {
		sg := format(string(gbts))
		so := format(string(bts))
		tb.Fatalf("output do not match:\ngot:\n%s\n\nexpected:\n%s\n\n", so, sg)
	}
}

func useLinuxEOL(bts []byte) []byte {
	return bytes.ReplaceAll(bts, []byte("\r\n"), []byte("\n"))
}

func format(str string) string {
	return strings.NewReplacer(
		"\x1b", "\\x1b",
		"\n", "\\n",
		"\r", "\\r",
		"\t", "\\t",
	).Replace(str)
}
