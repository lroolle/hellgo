package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// walkFiles starts a goroutine to walk the directory tree at root and send the
// path of each regular file on the string channel.  It sends the result of the
// walk on the error channel.  If done is closed, walkFiles abandons its work.
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() { // HL
		// Close the paths channel after Walk returns.
		defer close(paths) // HL
		// No select needed for this send, since errc is buffered.
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error { // HL
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path: // HL
			case <-done: // HL
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

func main() {
	done := make(chan struct{})
	defer close(done)
	paths, _ := walkFiles(done, os.Args[1])
	// Check whether the Walk failed.
	// if err := <-errc; err != nil {
	// 	fmt.Println(err)
	// }
	for path := range paths {
		fmt.Println(path)
	}
}
