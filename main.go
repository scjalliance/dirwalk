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

func depth(path string) int {
	return len(strings.Split(path, string(os.PathSeparator)))
}

func main() {
	kingpin.Parse()
	var err error
	rootAbs, err = filepath.Abs(*rootDir)
	rootDepth := depth(strings.TrimRight(rootAbs, string(os.PathSeparator)))
	if err != nil {
		panic(err)
	}
	err = filepath.Walk(rootAbs, func(path string, info os.FileInfo, err error) error {
		depth := depth(path)
		if err != nil {
			return nil
			// return err // maybe we want to complain but not explode?
		}
		if *depthMax >= 0 && depth-rootDepth > *depthMax {
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
