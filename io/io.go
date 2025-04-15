package myio

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func CreateFile(fileName string) {
	os.Remove(fileName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file error:", err)
	} else {
		defer file.Close()
		file.Chmod(0o666)
		fmt.Printf("fd=%d\n", file.Fd())

		file.WriteString("never ever give up\n")
		info, _ := file.Stat()
		fmt.Printf("file is dir: %t\n", info.IsDir())
		fmt.Printf("file mod time: %s\n", info.ModTime())
		fmt.Printf("mode %v\n", info.ModTime())
		fmt.Printf("size %vB\n", info.Size())
	}

	os.Mkdir("./data/sys", os.ModePerm)
	os.MkdirAll("./data/sys1/dir", 0o777)

	os.Rename("./data/sys", "./data/sys2")
	os.Rename("./data/sys1/dir", "./data/sys/dir") //move

}

// traverse file
func TraverseFile(path string) error {

	filepath.Walk(path, func(Subpath string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Printf("dir: %s\n", Subpath)
		} else if info.Mode().IsRegular() {
			fmt.Printf("file: %s, fileName: %s\n", Subpath, info.Name())
		}
		return nil
	})
	return nil
}

// io.Copy(fout, fin)
func CopyFile(src, dst string) {
	fin, err := os.Open(src)
	if err != nil {
		fmt.Println("open file error1:", err)
		return
	}

	fout, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		fmt.Println("open file error2:", err)
		return
	}

	// bs := make([]byte, 1024)
	// for {
	// 	n, err := fin.Read(bs)
	// 	if n > 0 {
	// 		fout.Write(bs[:n])
	// 	}
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }

	Writer := gzip.NewWriter(fout) //压缩
	//Writer := zlip.NewWriter(fout)   //压缩算法不一致
	io.Copy(Writer, fin)

	reader, _ := gzip.NewReader(fin) //解压
	fin.Close()
	fout.Close()
	return
}
