package builtin

import "io"

type Writer interface {
	io.Writer
}

type WriteCloser interface {
	io.WriteCloser
}
