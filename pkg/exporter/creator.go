package exporter

import (
	"encoding/csv"
	"github.com/xuri/excelize/v2"
	"io"
	"os"
)

type fileCreator struct{}

func (f fileCreator) Create(path string) (io.WriteCloser, error) {
	return os.Create(path)
}

type csvCreator struct{}

func (c csvCreator) Create(writer io.Writer) CSVWriteFlusher {
	return csv.NewWriter(writer)
}

type WriteCloserCreator interface {
	Create(path string) (io.WriteCloser, error)
}

type CSVWriterCreator interface {
	Create(writer io.Writer) CSVWriteFlusher
}

type excelCreator struct{}

func (e excelCreator) Create() ExcelWriter {
	return excelize.NewFile()
}

type ExcelWriterCreator interface {
	Create() ExcelWriter
}
