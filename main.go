package main

import (
	"log"
	"os"
	"syscall/js"

	"github.com/greymd/ojichat/generator"
)


var usage = `Usage:
  ojichat [options] [<name>]

Options:
  -h, --help      ヘルプを表示.
  -V, --version   バージョンを表示.
  -e <number>     絵文字/顔文字の最大連続数 [default: 4].
  -p <level>      句読点挿入頻度レベル [min:0, max:3] [default: 0].`


func run(this js.Value, i []js.Value) interface{} {
	config := generator.Config{}
	name := i[0].String()
	config.TargetName = name

	result, err := generator.Start(config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return result
}

func main() {
	js.Global().Set("__github_com_stdll00_ojichatjs", js.FuncOf(run))
	select {
	}
}