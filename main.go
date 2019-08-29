package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	rootDir     = kingpin.Arg("path", "Root directory for walking").Default(".").String()
	rootAbs     string
	outputDirs  = kingpin.Flag("dirs", "Output dirs?").Bool()
	outputFiles = kingpin.Flag("files", "Output files?").Bool()
	depthMax    = kingpin.Flag("depth", "Walk depth").Default("-1").Int()
)

func main() {
	kingpin.Parse()
	var err error
	rootAbs, err = filepath.Abs(*rootDir)
	rootDepth := len(strings.Split(rootAbs, string(os.PathSeparator)))
	if err != nil {
		panic(err)
	}
	err = filepath.Walk(rootAbs, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
			// return err // maybe we want to complain but not explode?
		}
		if depth := len(strings.Split(path, string(os.PathSeparator))); *depthMax >= 0 && depth-rootDepth > *depthMax {
			return filepath.SkipDir
		}
		if !*outputDirs && info.IsDir() {
			return nil
		}
		if !*outputFiles && !info.IsDir() {
			return nil
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
