package main

import (
	"embed"
	_ "embed" // harus diimport
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed GS_Fgn7bIAIp998.jpeg
var image []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("new_image.jpeg", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content: ", string(content))
		}
	}
}
