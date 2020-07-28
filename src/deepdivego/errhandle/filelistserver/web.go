package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// files/fib.txt
	http.HandleFunc("/files/",
		func(writer http.ResponseWriter, request *http.Request) {
			path := request.URL.Path[len("/files/"):]
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			all, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			writer.Write(all)
		})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
