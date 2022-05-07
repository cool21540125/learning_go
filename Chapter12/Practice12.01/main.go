package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	create()
	remove()
}

func create() {
	f, err := os.Create("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Println("demo.txt 已建立")
}

func remove() {
	err := os.Remove("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("demo.txt 已移除")
}
