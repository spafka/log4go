package main

import (
	"fmt"
	log "github.com/spafka/log4go"
	"io/ioutil"
	"os"
	"time"
)

func init() {

	jsonconfig := "{\"console\":{\"enable\":true,\"level\":\"DEBUG\",\"pattern\":\"[%D%T]%M\"},\"files\":[{\"enable\":true,\"level\":\"DEBUG\",\"filename\":\"console.log\",\"pattern\":\"[%D%T]%M\",\"category\":\"DEFAULT\",\"rotate\":true,\"maxsize\":\"10G\",\"maxlines\":\"10G\",\"daily\":true,\"maxdate\":\"90\"}]}"
	log.LoadConfigurationFromJson(jsonconfig)
}

func main() {

	dir, _ := os.Getwd()
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {

		}
	}

	for {

		log.Info("sssssssssss")
		time.Sleep(1 * time.Second)
	}

	t1 := time.Now()

	t2 := t1.AddDate(0, -1, 0)
	fmt.Println(t1)

	fmt.Println(t2)
}
