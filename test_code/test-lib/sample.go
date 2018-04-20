package sample

/*
	packageの挙動
	http://cuto.unirita.co.jp/gostudy/post/go-package/

  package名の命名規則など
  https://www.ymotongpoo.com/works/goblog-ja/post/package-names/
*/

import "log"

/*
  Package内の命名規則について
	別のpackageから参照させたい変数名や関数名、フィールド名(構造体内の変数名)  
	などの識別子名は、最初の文字を大文字で始める必要がある。
	
  小文字で始まる識別子名は、定義されたpackage内でのみ使用可能である。
*/

func Foo() {
	log.Print("Hello world from foo!")
}

type Mydata struct {
	Num1 int
	Str1 string
}