package covidlib

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

type covid struct {
	Date    string
	Discharged    string
	Expired  string
	Region string
	Still_admitted string
	Test string
	Positive string
}

func Load(path string) []covid {
	table := make([]covid, 0)
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err.Error())
		}
		c := covid{
			Date: row[4],
			Test: row[3],
			Positive:row[2],
			Discharged:    row[5],
			Expired:    row[6],
			Region:  row[7],
			Still_admitted: row[8],
		}
		table = append(table, c)
	}
	return table
}

func Find(table []covid, filter string) []covid {
	if filter == "" || filter == "*" {
		return table
	}
	result := make([]covid, 0)
	filter = strings.ToUpper(filter)
	for _, cur := range table {
		if cur.Discharged== filter ||
			cur.Date == filter ||
			cur.Still_admitted == filter ||
			cur.Test == filter ||
			cur.Positive == filter ||
			strings.Contains(strings.ToUpper(cur.Region), filter) ||
			strings.Contains(strings.ToUpper(cur.Expired), filter) {
			result = append(result, cur)
		}
	}
	return result
}
