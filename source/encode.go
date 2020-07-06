package main

import (
	"fmt"
	"github.com/devplayg/eggcrate"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs("./bootstrap4")
	if err != nil {
		panic(err)
	}
	config := eggcrate.Config{
		Dir:        dir,
		OutFile:    "output.go",
		UriPrefix:  "/assets",
		Extensions: "css,map,svg,js,eot,ttf,woff,woff2",
	}
	_, err = eggcrate.Encode(&config)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
