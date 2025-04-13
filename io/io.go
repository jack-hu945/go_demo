package myio

import (
	"bufio"
	"fmt"
	"io"
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

func ReadFile() {
	if fin, err := os.Open("./data/verse.txt"); err != nil {
		fmt.Println("open file error:", err)
	} else {
		defer fin.Close()
		buf := make([]byte, 128)
		n, err := fin.Read(buf)
		if err != nil {
			fmt.Println("read file error:", err)
		} else {
			fmt.Printf("read %d bytes:\ncontext:\n%s", n, buf[:n])
		}

		fin.Seek(3, 0)
		fin.Read(buf)
		fmt.Printf("重定向：%s ", string(buf))

		const BATCH = 10

		fmt.Println("for loop read:")
		buffer := make([]byte, BATCH)
		fin.Seek(0, 0)
		for {
			n, err := fin.Read(buffer)
			if n > 0 {
				fmt.Println(string(buffer[:n]))
			}
			if err == io.EOF {
				break
			}
		}
	}
}

func ReadFileWithBuffer() {
	if fin, err := os.Open("./data/verse.txt"); err != nil {
		fmt.Printf("open file failed: %v\n,", err)
	} else {
		defer fin.Close()
		reader := bufio.NewReader(fin)
		for {
			line, err := reader.ReadString('\n') //read until '\n'
			if len(line) > 0 {
				fmt.Print(line)
			}

			if err == io.EOF {
				break
			}
		}
	}
}

func WriteFileWithBuffer() {
	if fout, err := os.OpenFile("./data/verse.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666); err != nil {
		fmt.Println("open file error:", err)
	} else {
		defer fout.Close()
		writer := bufio.NewWriter(fout)
		for i := 0; i < 3; i++ {
			writer.WriteString("nice to meet you\n")
			writer.Write([]byte("it is good day\n"))
		}
		writer.Flush()
	}
}
