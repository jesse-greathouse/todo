package todo

import (
	"io/ioutil"
	"log"
)

type Helpers struct{}

// reads a file into a string designated by the resource identifier
func (h Helpers) ReadFile(uri string) string {
	content, err := ioutil.ReadFile(uri)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	return string(content)
}
