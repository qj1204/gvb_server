package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

const ConfigFile = "settings.yaml"

func main() {
	_, err := ioutil.ReadFile(ConfigFile)
	fmt.Println(getCurrentAbPathByCaller())
	fmt.Println(err)
}
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
