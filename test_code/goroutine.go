package main

import (
	"log"
	"time"
)
// goroutineについてのメモ

func f() {
	// 1 second sleep
	time.Sleep(time.Second)
	log.Println("goroutine")
}

func main() {
	go f()

	// gorutineより先に実行される = goroutineの結果を待たない
	log.Println("main")

	// 3 second sleep
	time.Sleep(3 * time.Second)
	log.Println("finish")
}