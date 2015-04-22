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

func Write(filename string, records [][]string) (e error) {
	file, e := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if e != nil {
		return e
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	for _, v := range records {
		e = writer.Write(addStrings(v))
		if e != nil {
			return e
		}
	}

	return nil
}

func addStrings(s []string) string {
	result := ""
	for _, v := range s {
		result += v + ","
	}
	return result
}
