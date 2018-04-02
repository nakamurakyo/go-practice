package main

import (
  "log"
)

func test(){
  defer log.Print("log3")
}

func test2(){
  defer log.Print("log4")
}

func main() {
  defer log.Print("log 1")
  defer log.Print("log 2")

  test()
  test2()

  log.Print("end test")
}