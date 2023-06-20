package exporter

import (
	"github.com/xuri/excelize/v2"
	"io"
)

func NewExcel(
	wcc WriteCloserCreator,
	ewc ExcelWriterCreator,
) *Excel {
	return &Excel{
		writeCloserCreator: wcc,
		excelWriterCreator: ewc,
	}
}

func DefaultExcel() *Excel {
	return NewExcel(
		&fileCreator{},
		&excelCreator{},
	)
}

type ExcelOptions = excelize.Options

type ExcelWriter interface {
	SetCellValue(sheet, cell string, value any) error
	Write(w io.Writer, opts ...ExcelOptions) error
}

type Excel struct {
	writeCloserCreator WriteCloserCreator
	excelWriterCreator ExcelWriterCreator
}

func (e Excel) Export(path string, table Table) error {
	excelFile := e.excelWriterCreator.Create()
	header := table.GetHeader()
	body := table.GetBody()
	content := append([][]string{header}, body...)
	for i := range content {
		for j := range content[i] {
			cellName, err := excelize.CoordinatesToCellName(j+1, i+1)
			if err != nil {
				return err
			}
			err = excelFile.SetCellValue("Sheet1", cellName, content[i][j])
			if err != nil {
				return err
			}
		}
	}

	file, err := e.writeCloserCreator.Create(path)
	if err != nil {
		return err
	}

	err = excelFile.Write(file)
	if err != nil {
		return err
	}

	return nil
}
