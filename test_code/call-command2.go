package main

import (
  "log"
  "os/exec"
)

func main() {
  log.Print("これはコマンドを実行するプログラムです!\n")
  out, error := exec.Command("ls", "-la").Output()

  if error != nil {
      log.Printf("エラー: %s", error)
      return
  }

  log.Printf("%s", out)
}