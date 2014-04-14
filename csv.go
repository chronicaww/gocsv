package csv

import (
	"encoding/csv"
	"os"
)

// 读取.csv文件，舍弃空行，舍弃id为空的行(包含舍弃由",,,,"构成的空行，)。
func Read(filename string) (records [][]string, e error) {
	file, e := os.Open(filename)
	if e != nil {
		return records, e
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comment = '#'
	reader.TrimLeadingSpace = true
	records, e = reader.ReadAll()
	if e != nil {
		return records, e
	}

	// 舍弃id为空的行(包含舍弃由",,,,"构成的空行，)。
	var recordsCleared [][]string
	for _, v := range records {
		if v[0] == "" {
			continue
		}
		recordsCleared = append(recordsCleared, v)
	}
	return recordsCleared, nil
}
