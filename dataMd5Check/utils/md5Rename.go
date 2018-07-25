package utils

import (
	"crypto/md5"
	"path/filepath"
	"os"
	"strings"
	"encoding/hex"
	"io"
	"fmt"
)


func CalculateMd5(file string) (value string, err error) {
	fi, err := os.Open(file)
	if err != nil {
		return
	}
	defer fi.Close()
	h := md5.New()
	_, err = io.Copy(h, fi)
	if err != nil {
		return
	}

	value = hex.EncodeToString(h.Sum(nil))
	return
}

func FindAll(path string, regex string) (files []string, err error) {
	files, err = filepath.Glob(path + regex)
	if err != nil {
		return
	}
	return
}

func GetFileName(file string) (name string) {
	paths :=  strings.Split(file, "/")
	fileName := paths[len(paths) - 1]
	name  = strings.Split(fileName, ".")[0]
	return
}

func Rename(oriFile string, newFileName string) (err error) {
	paths := strings.Split(oriFile, "/")
	dir := strings.Join(paths[0:len(paths)-1], "/")
	fmt.Println(dir + "/" + newFileName)
	return os.Rename(oriFile, dir + "/" + newFileName)
}

