package clog_test

import (
	"testing"

	"github.com/kumose-go/clog"
)

func TestFromContext(t *testing.T) {
	ctx := t.Context()

	logger := clog.FromContext(ctx)
	if logger != clog.Log {
		t.Fatalf("expected %v, got %v", clog.Log, logger)
	}

	logs := clog.WithField("foo", "bar")
	ctx = clog.NewContext(ctx, logs)

	logger = clog.FromContext(ctx)
	if logger != logs {
		t.Fatalf("expected %v, got %v", logs, logger)
	}
}
