package myio

import (
	"testing"
)

func TestWriteFile(t *testing.T) {
	WriteFile()
}

func TestReadFile(t *testing.T) {
	ReadFile()
}

func TestReadFileWithBuffer(t *testing.T) {
	ReadFileWithBuffer()
}

func TestWriteFileWithBufferLine(t *testing.T) {
	WriteFileWithBuffer()
}

func TestCreateFile(t *testing.T) {
	CreateFile("./data/myword.txt")
}

func TestTraverseFile(t *testing.T) {
	TraverseFile("./data")
}
