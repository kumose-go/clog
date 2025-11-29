package main

import (
	"errors"

	"github.com/kumose-go/clog"
)

func main() {

	clog.SetLevel(clog.DebugLevel)
	clog.
		WithField("field5", "value5").
		WithField("field2", "value2").
		WithField("field1", "value1").
		WithField("field4", "value4").
		WithField("FOO", "bar").
		WithField("field3", "value3").
		Info("AQUI")
	clog.WithField("foo", "bar").Debug("debug")
	clog.WithField("foo", "bar").Info("info")
	clog.WithField("foo", "bar").Warn("warn")
	clog.WithField("multiple", "fields").
		WithField("yes", true).
		Info("a longer line in this particular log")
	clog.IncreasePadding()
	clog.WithField("foo", "bar").Info("info with increased padding")
	clog.IncreasePadding()
	clog.WithField("foo", "bar").
		WithField("text", "a multi\nline text going\non for multiple lines\nhello\nworld!").
		Info("info with a more increased padding")
	clog.WithoutPadding().WithField("foo", "bar").Info("info without padding")
	clog.WithField("foo", "bar").Info("info with a more increased padding")
	clog.ResetPadding()
	clog.WithField("foo", "bar").
		WithField("text", "a multi\nline text going\non for multiple lines\nhello\nworld!").
		WithField("another", "bar").
		WithField("lalalal", "bar").
		Info("info with a more increased padding")
	clog.WithError(errors.New("some error")).Error("error")
	clog.WithError(errors.New("some fatal error")).Fatal("fatal")
}
