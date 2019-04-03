package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	dir, _ := os.Getwd()
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {

		}
	}

	t1 := time.Now()

	t2 := t1.AddDate(0, -1, 0)
	fmt.Println(t1)

	fmt.Println(t2)
}
