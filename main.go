package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"strings"
	"github.com/petar/GoLLRB/llrb"
	"os"
	"runtime"
)

const file = "/tmp/pangolin.db"

func main() {
	fmt.Println("pangolin is starting up")

 	http.HandleFunc("/", getSpec)

	tree := llrb.New() 
	fmt.Println(tree)
	fmt.Println(tree.Len())

	writeErr := Save(file, tree)
	Check(writeErr)

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  } 
}

func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}

func getSpec (w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}

func Save (path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}
