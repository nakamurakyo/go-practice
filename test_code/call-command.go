package main

import (
  "log"
  "os/exec"
)

////////////////////////////////////////////////////
// コマンドのサンプル
//

func main() {
  log.Print("hello log!!")


  out, error := exec.Command("date").Output()
  if error != nil {
      log.Printf("error: %s", error)
      return
  }

  log.Printf("now time: %s", out)
  log.Print("======下記はerrorのテストです========")

  out2, error2 := exec.Command("false").Output()
  if error2 != nil {
      log.Printf("error2: %s", error2)
      return
  }

  log.Printf("now time: %s", out2)
}
