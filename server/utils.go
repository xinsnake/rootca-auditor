package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const foldername = "wwwroot"

func WriteFile(filename string, content []byte) (err error) {
	err = ioutil.WriteFile(fmt.Sprintf("./%s/%s", foldername, filename), content, 0644)

	if err != nil {
		log.Println(fmt.Sprintf("Error processing file %s", filename))
		return err
	}

	log.Println(fmt.Sprintf("Processed file %s", filename))
	return nil
}
