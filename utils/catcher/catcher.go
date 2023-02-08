package catcher

import (
	"bitbucket.org/ifan-moladin/base-project/utils/httpserver"
)

type Catcher interface {
	Init() error
	WithContext(ctx httpserver.Context) Catcher
	Catch(err error)
	CaptureMessage(msg string)
}
