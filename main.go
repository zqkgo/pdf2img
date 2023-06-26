package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
	stdpath "path"
	"strings"

	"log"

	"github.com/karmdip-mi/go-fitz"
)

var path string

func init() {
	flag.StringVar(&path, "path", "", "-path=/path/to/pdf/file")
}

func main() {
	flag.Parse()
	if path == "" {
		panic("path must be specified")
	}
	dir := stdpath.Dir(path)
	name := stdpath.Base(path)
	imgDir := dir + "/" + strings.Split(name, ".")[0]
	log.Printf("use path: %s, dir: %s, name: %s', imgDir: %s", path, dir, name, imgDir)
	err := os.MkdirAll(imgDir, 0755)
	if err != nil {
		panic(err)
	}
	doc, err := fitz.New(path)
	if err != nil {
		panic(err)
	}

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}
		f, err := os.Create(stdpath.Join(imgDir, fmt.Sprintf("image-%05d.jpg", n)))
		if err != nil {
			panic(err)
		}
		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}
