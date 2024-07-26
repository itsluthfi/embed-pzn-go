package test

import (
	"embed"
	_ "embed" // harus diimport
	"fmt"
	"io/fs"
	"os"
	"testing"
)

// harus pake komen di bawah diikut nama filenya

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
	// outputnya sama kayak isi version.txt tanpa harus baca file secara manual
}

//go:embed GS_Fgn7bIAIp998.jpeg
var image []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("new_image.jpeg", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFIles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content: ", string(content))
		}
	}
}
