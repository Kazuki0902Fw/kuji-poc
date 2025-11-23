package csvs

import (
	"bytes"
	"encoding/csv"

	"github.com/pkg/errors"
)

func readCSV(src []byte) ([][]string, error) {
	bytesReader := bytes.NewReader(src)
	csvReader := csv.NewReader(bytesReader)
	csvReader.Comma = ','
	csvReader.FieldsPerRecord = -1

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return records, nil
}
