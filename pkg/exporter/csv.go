package exporter

func NewCSV(
	wcc WriteCloserCreator,
	cwv CSVWriterCreator,
) *CSV {
	return &CSV{
		writeCloserCreator: wcc,
		csvWriterCreator:   cwv,
	}
}

func DefaultCSV() *CSV {
	return NewCSV(&fileCreator{}, &csvCreator{})
}

type CSVWriteFlusher interface {
	Write(record []string) error
	Flush()
}

type CSV struct {
	writeCloserCreator WriteCloserCreator
	csvWriterCreator   CSVWriterCreator
}

func (c CSV) Export(path string, table Table) error {
	file, err := c.writeCloserCreator.Create(path)
	if err != nil {
		return err
	}

	writer := c.csvWriterCreator.Create(file)
	err = writer.Write(table.GetHeader())
	if err != nil {
		return err
	}

	rows := table.GetBody()
	for i := range rows {
		err = writer.Write(rows[i])
		if err != nil {
			return err
		}
	}

	writer.Flush()
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
