package utls

import (
	"io/ioutil"
)

func FetchFileContent(path string) []byte {
	fileData, err := ioutil.ReadFile(path)
	check(err)
	return fileData
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
