package main

import (
	"fmt"
	"log"
	"net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

////////////////////////////////////////////////////
// docker-compose up
//
func main() {
	log.Print("これはコマンドを実行するプログラムです!\n")

	out := exec_command("ls", "-la")
	log.Printf("%s", out)

	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}
