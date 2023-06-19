package utils

type Catcher interface {
	Init() error
	WithContext(ctx Context) Catcher
	Catch(err error)
	CaptureMessage(msg string)
}
