package myio

import (
	"fmt"
	"os"
)

func WriteFile() {
	fout, err := os.OpenFile("./data/verse.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		fmt.Println("open file error:", err)
	} else {
		defer fout.Close()
		fout.WriteString("hello world\n")
		fout.WriteString("This is sean\n")
	}
}
