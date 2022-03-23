package utils

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func SaveFile(fileName string, data interface{}) (err error) {
	mdFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer mdFile.Close()
	writer := bufio.NewWriter(mdFile)
	var bytes []byte
	slice,ok := data.([]string)
	if ok {
		data = strings.Join(slice, "\n")
	}
	bytes, err = json.Marshal(data)
	if err != nil {
		return
	}
	_, err = writer.Write(bytes)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return
}