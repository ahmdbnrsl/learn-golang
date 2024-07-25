package test

import "embed"
import _"embed"
import "fmt"
import "testing"
import "io/fs"
import "io/ioutil"

//go:embed version.txt
var version string

//go:embed foto_smk.jpg
var foto_smk []byte

func TestEmbed(t *testing.T) {
    fmt.Println(version)
}

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.jpg", foto_smk, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}