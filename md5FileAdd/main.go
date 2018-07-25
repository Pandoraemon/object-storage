package main

import (
	"log"
	"io/ioutil"
	"fmt"
	"math/rand"
	"crypto/md5"
	"time"
)

var rootDir = "/Users/pandoraemon/nfs/"
var rootFile = "rootFile"
func main() {
	file, err := ioutil.ReadFile(rootDir + rootFile)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(len(file))
	for i := 0; i < 10; i++ {
		rand.Seed(int64(time.Now().UnixNano()))
		randIndex := rand.Intn(len(file))
		fmt.Println(file[randIndex])
		fmt.Printf("%x\n", md5.Sum(file))
		file[randIndex] = ^file[randIndex]
		fmt.Println(file[randIndex])
		fmt.Printf("%x\n", md5.Sum(file))
		md5.Sum(file)
		fileName := fmt.Sprintf("%x", md5.Sum(file))
		newFile := rootDir + fileName + ".md5"
		ioutil.WriteFile(newFile, file, 0755)
	}
}
