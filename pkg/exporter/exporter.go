package exporter

type Exporter interface {
	Export(path string, table Table) error
}

type Table interface {
	GetHeader() []string
	GetBody() [][]string
}
