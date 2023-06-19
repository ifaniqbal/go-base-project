package catcher

import (
	"github.com/ifaniqbal/go-base-project/pkg/httpserver"
)

type Catcher interface {
	Init() error
	WithContext(ctx httpserver.Context) Catcher
	Catch(err error)
	CaptureMessage(msg string)
}
