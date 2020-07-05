package curlib

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
			Date: row[0],
			Discharged:    row[1],
			Expired:    row[2],
			Region:  row[3],
			Still_admitted: row[4],
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
		if cur.Region == filter ||
			cur.Expired == filter ||
			strings.Contains(strings.ToUpper(cur.Region), filter) ||
			strings.Contains(strings.ToUpper(cur.Expired), filter) {
			result = append(result, cur)
		}
	}
	return result
}
