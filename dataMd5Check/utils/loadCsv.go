package utils

import (
	"os"
	"encoding/csv"
	"io"
)

func LoadCsv(file string) ([]string, error) {
	dirs := make([]string, 0)
	fi, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	reader := csv.NewReader(fi)
	for {
		dir, err := reader.Read()
		if err == io.EOF{
			break
		} else if err != nil {
			return nil, err
		}
		dirs = append(dirs, dir[0])
	}
	return dirs, nil
}
