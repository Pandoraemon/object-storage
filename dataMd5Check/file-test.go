package main

import (
	"flag"
	"fmt"
	"sync"
	"log"
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"io/ioutil"
	"math/rand"
	"crypto/md5"
	"./utils"
	"os"
)

var cmd string

func init(){
	flag.StringVar(&cmd, "cmd", "create", "check md5 or rename jpg or create files")
}

func main() {
	var baseDir = "/mnt/nfs/"
	var rootFile = "rootFile"
	var wg sync.WaitGroup
	dirs, err:= utils.LoadCsv(baseDir + "dirs")
	flag.Parse()
	if err != nil {
		log.Println(err)
	}
	client := redis.NewClient(
		&redis.Options{
			Addr: "10.128.137.42:6379",
			Password: "",
			DB: 0,
		})
	for i := 0; i <10; i++ {
		go func() {
			wg.Add(1)
			index, err := client.Get("index").Result()
			if err != nil {
				log.Println(err)
			}
			client.Incr("index")
			dirIndex, _ := strconv.Atoi(index)

			dir := dirs[dirIndex]
			if cmd == "rename" {
				files, err:= utils.FindAll(baseDir + dir, "/*.jpg")
				if err != nil {
					log.Println(err)
				}
				fileCount := 0
				for _, file := range files {
					fmt.Println(fileCount)
					fileCount++

					md5Value, err := utils.CalculateMd5(file)
					if err != nil {
						log.Println(err)
					} else {
						utils.Rename(file, md5Value+".md5")
						fmt.Println(fileCount, "rename " + file + " to " + md5Value)
					}
				}
			} else if cmd == "check" {
				files, err:= utils.FindAll(baseDir + dir, "/*.md5")
				if err != nil {
					log.Println(err)
				}
				fileCount := 0
				for _, file := range files {
					fileCount++
					md5Value, err := utils.CalculateMd5(file)
					if err != nil {
						log.Println(err)
					} else {
						if md5Value == utils.GetFileName(file) {
							continue
						} else {
							log.Printf("%s md5 value should be %s but get %s", file, utils.GetFileName(file), md5Value)
						}
					}
				}
			} else if cmd == "create" {
				log.Printf("start create in %s!", dir)
				file, err := ioutil.ReadFile(baseDir + rootFile)
				if err != nil {
					log.Println(err)
				}
				os.Mkdir(baseDir + dir, 0751)
				//fmt.Println(len(file))
				for i := 0; i < 10000; i++ {
					rand.Seed(int64(time.Now().UnixNano()))
					randIndex := rand.Intn(len(file))
					file[randIndex] = ^file[randIndex]
					md5.Sum(file)
					fileName := fmt.Sprintf("%x", md5.Sum(file))
					newFile := baseDir + dir + "/" + fileName + ".md5"
					ioutil.WriteFile(newFile, file, 0755)
				}
			}
			log.Printf("finish %s 10000 files in %s", cmd, dir)
			wg.Done()
		}()
		time.Sleep(1 * time.Second)
	}

	wg.Wait()

}
