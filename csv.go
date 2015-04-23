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
	file, e := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if e != nil {
		return e
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	e = writer.WriteAll(records)
	if e != nil {
		return e
	}

	return nil
}

func reWrite(filename string, records [][]string) (e error) {
	file, e := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if e != nil {
		return e
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	e = writer.WriteAll(records)
	if e != nil {
		return e
	}

	return nil
}

func DelRecord(filename string, records [][]string) error { //注释某行
	file, e := os.OpenFile(filename, os.O_RDWR, 0777)
	if e != nil {
		return e
	}
	defer file.Close()

	recOld, e := Read(filename)
	if e != nil {
		return e
	}
	// k := 0
	j := len(recOld) - 1
	for i, v := range recOld {
		for _, u := range records {
			if u[0] == v[0] {
				tmprec := recOld[j]
				recOld[j] = recOld[i]
				recOld[i] = tmprec
				j--
				break
			}
		}
		if !(i < j) {
			break
		}
	}
	result := [][]string{}
	if j >= 0 {
		result = recOld[:j+1]
	}
	e = reWrite(filename, result)
	if e != nil {
		return e
	}
	return nil
}
