package codic_tool

import (
	"io/ioutil"
	"log"
	"os"
)

func setToTmp(cache_string string, filename string){
	content := []byte(cache_string)
	ioutil.WriteFile(filename, content, os.ModePerm)
}

func getFromTmp(filename string, noExistFileMessage string) string{
	// ファイルの読み込み
	contents,err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(noExistFileMessage)
		os.Exit(0)
	}
	return string(contents)
}