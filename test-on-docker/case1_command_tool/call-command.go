package main

import (
	"log"
	"os/exec"
)

func exec_command(command string, option string) string {
	out, error := exec.Command(command, option).Output()

	if error != nil {
		log.Printf("エラー: %s", error)
		return "error"
	}

	return string(out)
}
////////////////////////////////////////////////////
// docker-compose up
//
func main() {
	log.Print("これはコマンドを実行するプログラムです!\n")

	out := exec_command("ls", "-la")
	log.Printf("%s", out)
}
