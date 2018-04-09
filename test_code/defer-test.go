package main

import (
	"log"
)

func test() {
	defer log.Print("log 3")
}

func test2() {
	defer log.Print("log 4")
}

////////////////////////////////////////////////////
// defer のサンプル
//  defer 文を使うことで、それを呼び出した関数が終了する際に実行すべき処理を記述することができる
//

func main() {
	defer log.Print("log 1")
	defer log.Print("log 2")

	test()
	test2()

	log.Print("end test")
}
