package bulkdo

import (
	"encoding/csv"
	"io"
)

func readItems(in io.Reader) ([]map[string]string, error) {
	reader := csv.NewReader(in)
	rows := make([]map[string]string, 0)
	var header []string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}

	}

	return rows, nil
}
