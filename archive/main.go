package main

import (
	"flag"
	"fmt"
)

func main() {
	// Parse関数はコマンドライン引数を処理して、コードの中で処理できるようにする
	flag.Parse()
	// Arg関数でコマンドライン引数を受け取る
	arg := flag.Arg(0)
	fmt.Printf("Hello %s\n", arg)
}
